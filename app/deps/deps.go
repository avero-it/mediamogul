package deps

import (
	"github.com/avero-it/mediamogul/app/aws/s3"
	"github.com/avero-it/mediamogul/app/config"
	"github.com/avero-it/mediamogul/app/httpserver"
	"github.com/avero-it/mediamogul/app/signals"
	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
	"net/http"
)

type AppInfo struct {
	Name        string
	Description string
	AppFunction string
	Version     string
	Build       string
}

type Deps struct {
	Verbose     bool // config
	Router      *mux.Router
	S3Client    s3.Client
	NewRelicApp *newrelic.Application
	Log         *logrus.Entry
}

func NewDeps(config *config.Config, i AppInfo) (*Deps, error) {

	var deps Deps

	// =====================================================================================================================
	// s.logRUS

	// use JSONFormatter
	logrus.SetFormatter(&logrus.JSONFormatter{})

	rus := logrus.New()

	if config.Log.Verbose {
		rus.SetLevel(logrus.TraceLevel)
	}

	contextualizedLog := rus.WithFields(logrus.Fields{
		"appname": i.Name,
		"build":   i.Build,
		"version": i.Version,
	})

	deps.Log = contextualizedLog

	deps.Log.Info("log: initialized")

	// =====================================================================================================================
	// SIGNALS

	signals.InitSignals()
	deps.Log.Infof("signals: initialized")

	// =====================================================================================================================
	// S3 CLIENT

	deps.S3Client = s3.CreateS3Client(&config.AWS.S3.BucketName, config.AWS.Config)

	// =====================================================================================================================
	// HTTP SERVER

	// Create a traced mux router
	deps.Router = mux.NewRouter()

	helper := httpserver.NewHttpHelper()

	server := httpserver.NewHttpServer(deps.Router, deps.S3Client, deps.Log, deps.NewRelicApp, config.STT.URI, helper)

	server.Routes()

	go func() {
		var err error
		if config.HTTPServer.Secure {
			deps.Log.Infof("https: listening on: %s:%s", config.HTTPServer.Host,
				config.HTTPServer.Port)
			err = http.ListenAndServeTLS(config.HTTPServer.Host+":"+config.HTTPServer.Port,
				config.HTTPServer.CertFile,
				config.HTTPServer.KeyFile,
				deps.Router)
		} else {
			deps.Log.Infof("http: listening on: %s:%s", config.HTTPServer.Host,
				config.HTTPServer.Port)
			err = http.ListenAndServe(config.HTTPServer.Host+":"+config.HTTPServer.Port, deps.Router)
		}
		deps.Log.Errorf("http: listener closed with error: %v", err)
	}()

	var add_s string
	if config.HTTPServer.Secure == true {
		add_s = "s"
	}

	//////////////////////////////////////////////////////////////////////
	// OpenAPI docs (swagger)

	doc := redoc.Redoc{
		Title:       "Example API",
		Description: "Example API Description",
		SpecFile:    "./swagger.yml", // "./openapi.yaml"
		SpecPath:    "/swagger.yml",  // "/openapi.yaml"
		DocsPath:    "/docs",
	}

	http.ListenAndServe(config.DOCServer.Host+":"+config.DOCServer.Port, doc.Handler())

	// TODO: move handler to the same mux, no need for a separate server
	// TODO: embed swagger.yml in bin

	//////////////////////////////////////////////////////////////////////

	deps.Log.Infof("http" + add_s + " Deps: initialized")

	return &deps, nil
}
