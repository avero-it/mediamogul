package config

import (
	"io/ioutil"
	"os"
	"testing"

	"flag"
	"github.com/Flaque/filet"
)

type testConfig struct {
	Key1 string `env:"SERVICE_A1" validate:"required"`
	Key2 string `env:"SERVICE_A2"`
}

func TestDefaultLoader(t *testing.T) {
	t.Run("Testing non * config", func(*testing.T) {
		config := testConfig{}
		err := NewLoader().Load(config)

		if err == nil || err.Error() != "configuration has to be a * to a struct but got config.testConfig" {
			t.Error("non-pointer structure should return an error", err)
		}
	})

	t.Run("Testing with non existing backup file", func(*testing.T) {
		os.Setenv("SERVICE_A1", "from-env-A1")

		config := &testConfig{}
		err := NewLoader().
			WithDefaultFileName("/tmp/non-existing-backup-file-env").
			Load(config)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if config.Key1 != "from-env-A1" {
			t.Errorf("invalid value loading config")
		}
	})

	t.Run("Testing with existing backup env file", func(*testing.T) {
		os.Clearenv()

		file, err := fileWithConfig(`
SERVICE_A1=from-file-A1
SERVICE_A2=from-file-A2
			`)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		defer os.Remove(file)

		config := &testConfig{}
		err = NewLoader().
			WithDefaultFileName(file).
			Load(config)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if config.Key1 != "from-file-A1" || config.Key2 != "from-file-A2" {
			t.Errorf("invalid value loading config")
		}
	})

	t.Run("Testing validations", func(*testing.T) {
		os.Clearenv()
		os.Setenv("SERVICE_A1", "")
		os.Setenv("SERVICE_A2", "whatever2")

		config := &testConfig{}
		err := NewLoader().Load(config)

		if err == nil {
			t.Error("invalid configuration should be validated returning an error")
		}
	})
}

func TestLoaderWithSpecificFilename(t *testing.T) {
	os.Setenv("SERVICE_A2", "it-should-be-overwritten-from-env-A2")

	file, err := fileWithConfig(`
SERVICE_A1=from-file-A1
SERVICE_A2=from-file-A2
	`)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	defer os.Remove(file)

	config := &testConfig{}
	err = NewLoader().
		WithFileName(file).
		Load(config)

	if err != nil {
		t.Error("unexpected error:", err)
	}

	if config.Key1 != "from-file-A1" || config.Key2 != "it-should-be-overwritten-from-env-A2" {
		t.Errorf("invalid value loading config")
	}
}

func TestLoaderWithSpecificFileFlag(t *testing.T) {
	defer filet.CleanUp(t)

	tmpdir := filet.TmpDir(t, "")

	filet.File(t, tmpdir+"/.env", `
SERVICE_A1=from-file-A1
SERVICE_A2=from-file-A2
	`)

	var _ = flag.String("file", tmpdir+"/.env", "config file name")

	os.Setenv("SERVICE_A2", "it-should-be-overwritten-from-env-A2")

	config := &testConfig{}
	err := NewLoader().
		WithFileFlag("file").
		Load(config)

	if err != nil {
		t.Error("unexpected error:", err)
	}

	if config.Key1 != "from-file-A1" || config.Key2 != "it-should-be-overwritten-from-env-A2" {
		t.Errorf("invalid value loading config")
	}
}

func TestLoaderOnlyEnvironmental(t *testing.T) {
	t.Run("Only environment variables", func(*testing.T) {
		os.Setenv("SERVICE_A1", "whatever")
		os.Setenv("SERVICE_A2", "whatever2")

		file, err := fileWithConfig("SERVICE_A1=it-should-be-not-loaded")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		defer os.Remove(file)

		config := &testConfig{}
		err = NewLoader().
			WithDefaultFileName(file).
			OnlyEnvironment().
			Load(config)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if config.Key1 != "whatever" || config.Key2 != "whatever2" {
			t.Errorf("invalid value loading config")
		}
	})
}

func fileWithConfig(content string) (string, error) {
	tmpfile, err := ioutil.TempFile("/tmp", "*.env")

	if err != nil {
		return "", err
	}

	if _, err := tmpfile.Write([]byte(content)); err != nil {
		return "", err
	}

	if err := tmpfile.Close(); err != nil {
		return "", err
	}

	return tmpfile.Name(), nil
}
