package cache

import "github.com/sirupsen/logrus"

var logger = logrus.WithFields(logrus.Fields{
	"component":   "cache",
	"api_version": "v1",
})
