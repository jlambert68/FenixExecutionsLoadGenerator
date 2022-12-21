package common_config

import (
	"fyne.io/fyne/v2"
	"github.com/sirupsen/logrus"
)

// Used for keeping track of the proto file versions for this ExecutionsLoadGenerator
var highestExecutionsLoadGeneratorProtoFileVersion int32 = -1

// Logger that all part of the system can use
var Logger *logrus.Logger

const LocalWebServerAddressAndPort = "127.0.0.1:8080"

var FenixCAConnectorApplicationReference *fyne.App
