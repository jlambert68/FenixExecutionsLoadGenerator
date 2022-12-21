RunGrpcGui:
	cd ~/egen_kod/go/go_workspace/src/jlambert/grpcui/standalone && grpcui -plaintext localhost:6673

filename :=
filenamePartFirst := FenixCAConnectorCrossBuild_
filenamePartLast := .exe
datetime := `date +'%y%m%d_%H%M%S'`

GenerateDateTime:
	$(eval fileName := $(filenamePartFirst)$(datetime)$(filenamePartLast))

	echo $(fileName)

GenerateTrayIcons:
	./bundleIcons.sh

BuildExeForWindows:
#	fyne-cross windows -arch=amd64 --ldflags="-X 'main.useInjectedEnvironmentVariables=true' -X 'main.runInTray=truex' -X 'main.loggingLevel=DebugLevel' -X 'main.ExecutionsLoadGeneratorPort=6672' -X 'main.executionLocationForConnector=LOCALHOST_NODOCKER' -X 'main.executionLocationForWorker=GCP' -X 'main.GuiExecutionAddress=fenixGuiExecution-ca-nwxrrpoxea-lz.a.run.app' -X 'main.GuiExecutionPort=443' -X 'main.gcpAuthentication=false'"
#	GOOD=windows GOARCH=amd64 go build -o FenixCAConnectorWindow.exe -ldflags="-X 'main.useInjectedEnvironmentVariables=true' -X 'main.runInTray=truex' -X 'main.loggingLevel=DebugLevel' -X 'main.ExecutionsLoadGeneratorPort=6672' -X 'main.executionLocationForConnector=LOCALHOST_NODOCKER' -X 'main.executionLocationForWorker=GCP' -X 'main.GuiExecutionAddress=fenixGuiExecution-ca-nwxrrpoxea-lz.a.run.app' -X 'main.GuiExecutionPort=443' -X  'main.gcpAuthentication=true' -X 'main.caEngineAddress=127.0.0.1' -X 'main.caEngineAddressPath=/"
	env GOOD=windows GOARCH=amd64 go build  -o FenixCAConnector.WindowsExe -ldflags="-X 'main.useInjectedEnvironmentVariables=true' -X 'main.runInTray=truex' -X 'main.loggingLevel=DebugLevel' -X 'main.ExecutionsLoadGeneratorPort=6672' -X 'main.executionLocationForConnector=LOCALHOST_NODOCKER' -X 'main.executionLocationForWorker=GCP' -X 'main.GuiExecutionAddress=fenixGuiExecution-ca-must-be-logged-in-nwxrrpoxea-lz.a.run.app' -X 'main.GuiExecutionPort=443' -X 'main.gcpAuthentication=true' -X 'main.caEngineAddress=127.0.0.1' -X 'main.caEngineAddressPath=x' -X 'main.useInternalWebServerForTest=true' -X 'main.useServiceAccount=true'" /home/jlambert/egen_kod/go/go_workspace/src/jlambert/FenixCAConnector
BuildExeForLinux:
	GOOD=linux GOARCH=amd64 go build  -o FenixCAConnector.LinuxExe -ldflags="-X 'main.useInjectedEnvironmentVariables=true' -X 'main.runInTray=truex' -X 'main.loggingLevel=DebugLevel' -X 'main.ExecutionsLoadGeneratorPort=6672' -X 'main.executionLocationForConnector=LOCALHOST_NODOCKER' -X 'main.executionLocationForWorker=GCP' -X 'main.GuiExecutionAddress=fenixGuiExecution-ca-must-be-logged-in-nwxrrpoxea-lz.a.run.app' -X 'main.GuiExecutionPort=443' -X 'main.gcpAuthentication=true' -X 'main.caEngineAddress=http://127.0.0.1:3000' -X 'main.caEngineAddressPath=/TestCaseExecution/ExecuteTestActionMethod' -X 'main.useInternalWebServerForTest=true' -X 'main.useServiceAccount=true' -X 'main.turnOffCallToWorker=true'"

CrossBuildForWindows:
	$(eval fileName := $(filenamePartFirst)$(datetime)$(filenamePartLast))
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc go build -o $(fileName) -ldflags="-X 'main.useInjectedEnvironmentVariables=true' -X 'main.runInTray=truex' -X 'main.loggingLevel=DebugLevel' -X 'main.ExecutionsLoadGeneratorPort=6672' -X 'main.executionLocationForConnector=LOCALHOST_NODOCKER' -X 'main.executionLocationForWorker=GCP' -X 'main.GuiExecutionAddress=fenixGuiExecution-ca-must-be-logged-in-nwxrrpoxea-lz.a.run.app' -X 'main.GuiExecutionPort=443' -X 'main.gcpAuthentication=true' -X 'main.caEngineAddress=http://127.0.0.1:3000' -X 'main.caEngineAddressPath=/TestCaseExecution/ExecuteTestActionMethod' -X 'main.useInternalWebServerForTest=true' -X 'main.useServiceAccount=true' -X 'main.turnOffCallToWorker=false'" .