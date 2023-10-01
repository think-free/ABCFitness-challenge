package cliparams

import (
	"github.com/kelseyhightower/envconfig"
)

type ClientParameters struct {
	DatabaseType string `envconfig:"dbtype" required:"false" default:"memory"`
	LogLevel     string `envconfig:"loglevel" required:"false" default:"debug"`
}

func New() *ClientParameters {
	cp := &ClientParameters{}
	envconfig.MustProcess("", cp)
	return cp
}
