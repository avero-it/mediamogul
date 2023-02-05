package main

import (
	_ "expvar"
	"github.com/avero-it/mediamogul/pkg/config"
	"github.com/avero-it/mediamogul/pkg/deps"
	"github.com/avero-it/mediamogul/pkg/mediamogul"
	"github.com/sirupsen/logrus"
)

// =====================================================================================================================
// MAIN
func main() {
	// new comment
	info := deps.AppInfo{
		Name:        "MediaMogul",
		Description: "List of media likeds",
		AppFunction: "Allows people to list and match media content",
		Build:       "3",
		Version:     "0.0.0",
	}

	cfg := &config.Config{}
	loader := config.NewLoader()

	if err := loader.WithFileName("dist.env").Load(cfg); err != nil {
		logrus.Fatalf("failed to load config: %v", err)
	}

	deps, err := deps.NewDeps(cfg, info)
	if err != nil {
		logrus.Fatalf("failed to prepare deps: %v", err)
	}

	MediaMogul := mediamogul.NewMediaMogul(
		deps.Router,
		deps.NewRelicApp,
		deps.Log,
	)

	deps.Log.Infof("Starting: \"%s\" - Build: %s - Version \"%s\"", info.Name, info.Build, info.Version)
	if err := MediaMogul.Run(); err != nil {
		logrus.Println(info.Name+" returned error: ", err)
	}
}
