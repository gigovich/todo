package todo

import (
	"github.com/gigovich/fargo/app"
	"github.com/gigovich/fargo/core/log"
	"github.com/sirupsen/logrus"
)

var App = app.New(
	app.WithSettingsFile(),
	app.WithLogger(log.WrapLogrus(logrus.StandardLogger()))
)
