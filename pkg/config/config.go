package config

import (
	"flag"
	"fmt"
	"os"
	"reflect"

	"github.com/go-playground/validator"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/imdario/mergo"
)

type Config struct {
	HTTPServer struct {
		Secure   bool   `env:"HTTPSERVER_SECURE"`
		Host     string `env:"HTTPSERVER_HOST" validate:"required"`
		Port     string `env:"HTTPSERVER_PORT" validate:"required"`
		CertFile string `env:"HTTPSERVER_CERTFILE"`
		KeyFile  string `env:"HTTPSERVER_KEYFILE"`
	}
	DOCServer struct {
		Host string `env:"DOCSERVER_HOST" validate:"required"`
		Port string `env:"DOCSERVER_PORT" validate:"required"`
	}
	STT struct {
		URI string `env:"STT_URI" validate:"required"`
	}
	Log struct {
		Verbose bool `env:"LOG_VERBOSE"`
	}
}

// guard
var _ Loader = loader{}

// Loader interface used by the LoadConfig method
type Loader interface {
	Load(cfg any) error
}

type loader struct {
	defaultFileName string
	fileFlag        string
	fileName        string
}

// NewLoader generates a default config loader
func NewLoader() *loader {
	return &loader{
		defaultFileName: ".env",
		fileFlag:        "config",
		fileName:        "",
	}
}

// WithDefaultFileName allows you to specify a custom default file name
// Default file name will be check in case no other file name is specified
func (l *loader) WithDefaultFileName(flnm string) *loader {
	l.defaultFileName = flnm

	return l
}

// WithFileFlag allows you to specify a custom file flag from where
// file name will be read. This file name will be used in case it is spcified the flag
// before to backup into the default one
// Take into consideration that going for flag will disable filename option
func (l *loader) WithFileFlag(f string) *loader {
	l.fileFlag = f
	l.fileName = ""

	return l
}

// WithFileName allows you to specify an specific file name
// Take into consideration that going for filename will disable flag option
func (l *loader) WithFileName(flnm string) *loader {
	l.fileName = flnm
	l.fileFlag = ""

	return l
}

// OnlyEnvironment makes loader only read from environment variables
func (l *loader) OnlyEnvironment() *loader {
	l.fileFlag = ""
	l.fileName = ""

	return l
}

// Load will fill the specified configuration structure from the source
func (l loader) Load(cfg any) error {
	if reflect.ValueOf(cfg).Kind() != reflect.Ptr {
		return fmt.Errorf("configuration has to be a pointer to a struct but got %T", cfg)
	}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return fmt.Errorf("error reading environment variables: %v", err)
	}

	// Copying structure to parse file configuration without affecting
	// environment variables in case there are some defined
	fileCfg := reflect.New(reflect.ValueOf(cfg).Elem().Type()).Interface()
	fileName := &l.fileName

	// Puts priority in case file name was specified instead of going for the flag
	// If both are empty we do nothing...
	if *fileName == "" && l.fileFlag != "" {
		f := flag.Lookup(l.fileFlag)

		if f == nil {
			fileName = flag.String(l.fileFlag, "", "Specify configuration file")
			flag.Parse()
		} else {
			*fileName = f.Value.String()
		}

		flag.Parse()

		// Only will set to read default config file if loader was configured to expect file flag
		// but the application was executed without specifying it and backup config file really exists
		if _, err := os.Stat(l.defaultFileName); *fileName == "" && err == nil {
			*fileName = l.defaultFileName
		}
	}

	// If it ends up in no file name defined neither from config flag, specific file name
	// nor backup file we don't try to read anything
	if *fileName != "" {
		if err := cleanenv.ReadConfig(*fileName, fileCfg); err != nil {
			return fmt.Errorf("error reading configuration from file: %v", err)
		}
	}

	// This will always put priority into the configuration comming from env variables
	if err := mergo.Merge(cfg, fileCfg); err != nil {
		return fmt.Errorf("unexpected error merging configuration (%v)", err)
	}

	// Validating in case it was defined within the structure
	if err := validator.New().Struct(cfg); err != nil {
		return fmt.Errorf("configuration %v", err)
	}

	return nil
}
