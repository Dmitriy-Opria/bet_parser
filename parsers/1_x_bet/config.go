package main

import (
	"errors"
	"github.com/Dmitriy-Opria/bet_parser/log"
	"github.com/pkg/errors"
	"net/url"
	"strings"
	"time"
)

// Config is application configuration.
type Config struct {
	AppName     string        `json:"appName"`
	HTTP        HTTPConfig    `json:"http"`
	LogLevel    string        `json:"logLevel"`
	MessageType string        `json:"messageType"`
	Tracing     TracingConfig `json:"tracing"`
}

// Validate validates the app configuration.
func (c Config) Validate() []error {
	errs := c.HTTP.Validate()

	if strings.TrimSpace(c.LogLevel) == "" {
		errs = append(errs, errors.New("MessageProcessorConfig requires a non-empty LogLevel config value"))
	}
	if len(c.MessageType) == 0 {
		errs = append(errs, errors.New("message-type is required"))
	}
	return append(errs, c.Tracing.Validate()...)
}

// ExportLogFields exports fields for logging.
func (c Config) ExportLogFields() log.Fields {
	return log.Fields{
		log.FieldService:    c.AppName,
		log.FieldHost:       c.HTTP.Host,
		log.FieldRemotePort: c.HTTP.Port,
		"shutdowntimeout":   c.HTTP.ShutdownTimeout,
	}
}

// HTTPConfig is configuration of an HTTP service.
type HTTPConfig struct {
	Host            string        `json:"host"`
	Port            int           `json:"port"`
	ShutdownTimeout time.Duration `json:"shutdownTimeout"`
}

// Validate validates HTTPConfig configuration values.
func (c HTTPConfig) Validate() []error {
	var errs []error

	if len(c.Host) == 0 {
		errs = append(errs, errors.Errorf("HTTPConfig requires a non-empty Host config value"))
	}

	if c.Port <= 0 {
		errs = append(errs, errors.Errorf("HTTPConfig requires a postive Port config value"))
	}

	return errs
}

// TracingConfig is configuration to connect to a tracing collector.
type TracingConfig struct {
	CollectorURLStr string
	CollectorURL    *url.URL `json:"collectorUrl"`
}

// Validate validates tracing configuration.
func (c TracingConfig) Validate() []error {
	var errs []error

	if _, err := url.Parse(c.CollectorURLStr); err != nil {
		errs = append(errs, errors.Wrap(err, "parse collector url string"))
	}

	return errs
}
