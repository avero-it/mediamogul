package main

import (
	_ "expvar"
	"github.com/avero-it/mediamogul/app/config"
	"github.com/avero-it/mediamogul/app/deps"
	"github.com/avero-it/mediamogul/app/mediamogul"
	"github.com/sirupsen/logrus"
)

// =====================================================================================================================
// MAIN
func main() {
	info := deps.AppInfo{
		Name:        "MediaModul",
		Description: "Speech To X",
		AppFunction: "NLP processing from audio",
		Build:       "3",
		Version:     "0.0.0",
	}

	cfg := &config.Config{}
	loader := config.NewLoader()

	if err := loader.WithFileName(".env").Load(cfg); err != nil {
		logrus.Fatalf("failed to load config: %v", err)
	}

	deps, err := deps.NewDeps(cfg, info)
	if err != nil {
		logrus.Fatalf("failed to prepare deps: %v", err)
	}

	Stx := mediamogul.NewStx(
		deps.Router,
		deps.NewRelicApp,
		deps.Log,
	)

	deps.Log.Infof("Starting: \"%s\" - Build: %s - Version \"%s\"", info.Name, info.Build, info.Version)
	if err := Stx.Run(); err != nil {
		logrus.Println(info.Name+" returned error: ", err)
	}
}
