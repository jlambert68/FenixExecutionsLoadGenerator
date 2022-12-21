package main

import (
	"FenixExecutionsLoadGenerator/common_config"
	"FenixExecutionsLoadGenerator/gcp"
	"FenixExecutionsLoadGenerator/resources"
	"flag"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	//"flag"
	"fmt"
	"log"
	"os"
)

// mustGetEnv is a helper function for getting environment variables.
// Displays a warning if the environment variable is not set.
func mustGetenv(environmentVariableName string) string {

	var environmentVariable string

	if useInjectedEnvironmentVariables == "true" {
		// Extract environment variables from parameters feed into program at compilation time

		switch environmentVariableName {
		case "applicationShouldRunInTray":
			environmentVariable = applicationShouldRunInTray

		case "loggingLevel":
			environmentVariable = loggingLevel

		case "executionsLoadGeneratorPort":
			environmentVariable = executionsLoadGeneratorPort

		case "executionLocationForExecutionsLoadGenerator":
			environmentVariable = executionLocationForExecutionsLoadGenerator

		case "executionLocationForGuiExecutionServer":
			environmentVariable = executionLocationForGuiExecutionServer

		case "guiExecutionServerAddress":
			environmentVariable = guiExecutionServerAddress

		case "guiExecutionServerPort":
			environmentVariable = guiExecutionServerPort

		case "gCPAuthentication":
			environmentVariable = gCPAuthentication

		case "useServiceAccount":
			environmentVariable = useServiceAccount

		default:
			log.Fatalf("Warning: %s environment variable not among injected variables.\n", environmentVariableName)

		}

		if environmentVariable == "" {
			log.Fatalf("Warning: %s environment variable not set.\n", environmentVariableName)
		}

	} else {
		//
		environmentVariable = os.Getenv(environmentVariableName)
		if environmentVariable == "" {
			log.Fatalf("Warning: %s environment variable not set.\n", environmentVariableName)
		}

	}
	return environmentVariable
}

// Variables injected at compilation time
var (
	useInjectedEnvironmentVariables             string
	applicationShouldRunInTray                  string
	loggingLevel                                string
	executionsLoadGeneratorPort                 string
	executionLocationForExecutionsLoadGenerator string
	executionLocationForGuiExecutionServer      string
	guiExecutionServerAddress                   string
	guiExecutionServerPort                      string
	gCPAuthentication                           string
	useServiceAccount                           string
)

func dumpMap(space string, m map[string]interface{}) {
	for k, v := range m {
		if mv, ok := v.(map[string]interface{}); ok {
			fmt.Printf("{ \"%v\": \n", k)
			dumpMap(space+"\t", mv)
			fmt.Printf("}\n")
		} else {
			fmt.Printf("%v %v : %v\n", space, k, v)
		}
	}
}

func main() {

	// Parse flags if there are any. Used to override hard set values from build process
	flagLoggingLevel := flag.String("flagLoggingLevel", "", "flagLoggingLevel=InfoLevel [expects: 'DebugLevel', 'InfoLevel']")
	flagRunInTray := flag.String("flagRunInTray", "", "flagRunInTray=xxxxx [expects: 'true', 'false']")

	flag_ldflags := flag.String("ldflags", "", "ldflags should not be used")

	/*
	   RunInTray:
	   	true
	   UseInternalWebServerForTest:
	   	true
	   useServiceAccount:
	   	true


	*/
	// Parse flags a secure that only expected value are used
	flag.Parse()

	fmt.Println(*flag_ldflags)

	// Verify flag for 'loggingLevel'

	switch *flagLoggingLevel {

	case "":

	case "DebugLevel":
		common_config.LoggingLevel = logrus.DebugLevel

	case "InfoLevel":
		common_config.LoggingLevel = logrus.InfoLevel

	default:
		fmt.Println("Unknown loggingLevel: '" + loggingLevel + "'. Expected one of the following: 'DebugLevel', 'InfoLevel'")
		os.Exit(0)
	}

	// Verify flag for 'RunInTray '
	switch *flagRunInTray {

	case "", "true", "false":

	default:
		fmt.Println("Unknown RunInTray-parameter '" + *flagRunInTray + "'. Expected one of the following: '', 'true', 'false'")
		os.Exit(0)
	}

	var logFileName string

	// Extract from environment variables if it should run as a tray application or not
	var shouldRunInTray string
	if *flagRunInTray == "" {
		shouldRunInTray = mustGetenv("ApplicationShouldRunInTray")
	} else {
		shouldRunInTray = *flagRunInTray
	}

	// When run as Tray application then add log-name
	if shouldRunInTray == "true" {
		common_config.ApplicationShouldRunInTray = true
		logFileName = "fenixConnectorLog.log"
	} else {
		logFileName = ""
	}

	// Initiate logger in common_config
	InitLogger(logFileName)

	// When Execution Worker runs on GCP, then set up access
	if common_config.ExecutionLocationForGuiExecutionServer == common_config.GCP &&
		common_config.GCPAuthentication == true {
		gcp.Gcp = gcp.GcpObjectStruct{}

		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

		// Generate first time Access token
		_, returnMessageAckNack, returnMessageString := gcp.Gcp.GenerateGCPAccessToken(ctx)
		if returnMessageAckNack == false {

			// If there was any problem then exit program
			common_config.Logger.WithFields(logrus.Fields{
				"id": "20c90d94-eef7-4819-ba8c-b7a56a39f995",
			}).Fatalf("Couldn't generate access token for GCP, return message: '%s'", returnMessageString)

		}
	}

	// Start Connector Engine
	go fenixExecutionsLoadGeneratorMain()

	// Start up Tray Application if it should that as that
	if shouldRunInTray == "true" {
		// Start application as TrayApplication

		a := app.NewWithID("FenixCAConnector")

		// Store reference to application, used for turning icon in tray RED/GREEN depending on when there is a connection to Worker
		common_config.FenixCAConnectorApplicationReference = &a

		a.SetIcon(resources.ResourceFenix83red32x32Png)
		mainFyneWindow := a.NewWindow("SysTray")

		if desk, ok := a.(desktop.App); ok {
			m := fyne.NewMenu("Fenix Execution Connector",
				fyne.NewMenuItem("Hide", func() {
					mainFyneWindow.Hide()
					newNotification := fyne.NewNotification("Fenix Execution Connector", "Fenix will rule the 'Test World'")

					a.SendNotification(newNotification)
				}))
			desk.SetSystemTrayMenu(m)
		}

		// Create Fenix Splash screen
		var splashWindow fyne.Window
		if drv, ok := fyne.CurrentApp().Driver().(desktop.Driver); ok {
			splashWindow = drv.CreateSplashWindow()

			// Fenix Header
			fenixHeaderText := canvas.Text{
				Alignment: fyne.TextAlignCenter,
				Color:     nil,
				Text:      "Fenix Inception - SaaS",
				TextSize:  20,
				TextStyle: fyne.TextStyle{Bold: true},
			}

			// Text Footer
			halFinney := widget.NewLabel("\"If you want to change the world, don't protest. Write code!\" - Hal Finney (1994)")

			// Fenix picture
			image := canvas.NewImageFromResource(resources.ResourceFenix12Png)
			image.FillMode = canvas.ImageFillOriginal

			// Container holding Header, picture and Footer
			spashContainer := container.New(layout.NewVBoxLayout(), &fenixHeaderText, image, halFinney)

			splashWindow.SetContent(spashContainer)
			splashWindow.CenterOnScreen()
			splashWindow.Show()

			go func() {
				time.Sleep(time.Millisecond * 1000)

				mainFyneWindow.Hide()

				time.Sleep(time.Second * 7)
				splashWindow.Close()

			}()

			mainFyneWindow.SetContent(widget.NewLabel("Fyne System Tray"))
			mainFyneWindow.SetCloseIntercept(func() {
				mainFyneWindow.Hide()
			})

			mainFyneWindow.Hide()
			go func() {
				count := 10
				for {
					time.Sleep(time.Millisecond * 100)
					//mainFyneWindow.Hide()
					count = count - 1
					if count == 0 {
						break
					}
					return
				}

				//mainFyneWindow.Hide()
			}()

			mainFyneWindow.ShowAndRun()
		}

	} else {
		// Run as console program and exit as on standard exiting signals
		sig := make(chan os.Signal, 1)
		done := make(chan bool, 1)

		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			sig := <-sig
			fmt.Println()
			fmt.Println(sig)
			done <- true

			fmt.Println("ctrl+c")
		}()

		fmt.Println("awaiting signal")
		<-done
		fmt.Println("exiting")
	}
}

/*


executionsLoadGeneratorPort=6673
executionLocationForExecutionsLoadGenerator=LOCALHOST_NODOCKER
executionLocationForGuiExecutionServer=LOCALHOST_NODOCKER
*gCPAuthentication=false
*guiExecutionServerAddress=127.0.0.1
*guiExecutionServerPort=6671
*loggingLevel=DebugLevel
applicationShouldRunInTray=truex
*useServiceAccount=false

*/

func init() {

	var err error

	// Get Environment variable to tell were ExecutionsLoadGenerator is running
	var tempExecutionLocationForExecutionsLoadGenerator = mustGetenv("ExecutionLocationForExecutionsLoadGenerator")

	switch tempExecutionLocationForExecutionsLoadGenerator {
	case "LOCALHOST_NODOCKER":
		common_config.ExecutionLocationForExecutionsLoadGenerator = common_config.LocalhostNoDocker

	case "LOCALHOST_DOCKER":
		common_config.ExecutionLocationForExecutionsLoadGenerator = common_config.LocalhostDocker

	case "GCP":
		common_config.ExecutionLocationForExecutionsLoadGenerator = common_config.GCP

	default:
		fmt.Println("Unknown Execution location for ExecutionsLoadGenerator: " + tempExecutionLocationForExecutionsLoadGenerator + ". Expected one of the following: 'LOCALHOST_NODOCKER', 'LOCALHOST_DOCKER', 'GCP'")
		os.Exit(0)

	}
	// Get Environment variable to tell were GuiExecutionServer is running
	var executionLocationForExecutionLocationForGuiExecutionServer = mustGetenv("ExecutionLocationForGuiExecutionServer")

	switch executionLocationForExecutionLocationForGuiExecutionServer {
	case "LOCALHOST_NODOCKER":
		common_config.ExecutionLocationForGuiExecutionServer = common_config.LocalhostNoDocker

	case "LOCALHOST_DOCKER":
		common_config.ExecutionLocationForGuiExecutionServer = common_config.LocalhostDocker

	case "GCP":
		common_config.ExecutionLocationForGuiExecutionServer = common_config.GCP

	default:
		fmt.Println("Unknown Execution location for Fenix Execution Worker Server: " + executionLocationForExecutionLocationForGuiExecutionServer + ". Expected one of the following: 'LOCALHOST_NODOCKER', 'LOCALHOST_DOCKER', 'GCP'")
		os.Exit(0)

	}

	// Address to Fenix GuiExecution Server
	common_config.GuiExecutionServerAddress = mustGetenv("GuiExecutionServerAddress")

	// Port for Fenix GuiExecution Server
	common_config.GuiExecutionServerPort, err = strconv.Atoi(mustGetenv("GuiExecutionServerPort"))
	if err != nil {
		fmt.Println("Couldn't convert environment variable 'GuiExecutionServerPort' to an integer, error: ", err)
		os.Exit(0)

	}

	// Port for Fenix Execution Connector Server
	common_config.ExecutionsLoadGeneratorPort, err = strconv.Atoi(mustGetenv("ExecutionsLoadGeneratorPort"))
	if err != nil {
		fmt.Println("Couldn't convert environment variable 'ExecutionsLoadGeneratorPort' to an integer, error: ", err)
		os.Exit(0)

	}

	// Build the Dial-address for gPRC-call
	common_config.FenixGuiExecutionServerAddressToDial = common_config.GuiExecutionServerAddress + ":" + strconv.Itoa(common_config.GuiExecutionServerPort)

	// Extract Debug level
	var tempLoggingLevel = mustGetenv("LoggingLevel")

	switch tempLoggingLevel {

	case "DebugLevel":
		common_config.LoggingLevel = logrus.DebugLevel

	case "InfoLevel":
		common_config.LoggingLevel = logrus.InfoLevel

	default:
		fmt.Println("Unknown loggingLevel '" + tempLoggingLevel + "'. Expected one of the following: 'DebugLevel', 'InfoLevel'")
		os.Exit(0)

	}

	// Extract if there is a need for authentication when going toward GCP
	boolValue, err := strconv.ParseBool(mustGetenv("GCPAuthentication"))
	if err != nil {
		fmt.Println("Couldn't convert environment variable 'gCPAuthentication:' to an boolean, error: ", err)
		os.Exit(0)
	}
	common_config.GCPAuthentication = boolValue

	// Extract if a service account should be used towards GCP
	boolValue, err = strconv.ParseBool(mustGetenv("UseServiceAccount"))
	if err != nil {
		fmt.Println("Couldn't convert environment variable 'UseServiceAccount:' to an boolean, error: ", err)
		os.Exit(0)
	}
	common_config.UseServiceAccount = boolValue

}
