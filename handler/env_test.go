package handler

import (
	"os"
	"testing"
)

func TestItCanLoadDefaultEnvFiles(t *testing.T) {
	t.Parallel()

	env, err := Make(PackageDir)

	if err != nil {
		t.Errorf("%v", err)
		t.FailNow()
	}

	if env.GetRootDir() != PackageDir {
		t.Errorf("The given env root directoy is invalid")
	}

	if env.Get("FILE_NAME") != ".env" {
		t.Errorf("The given env file is invalid")
	}
}

func TestItCanLoadGivenEnvFiles(t *testing.T) {
	t.Parallel()

	env, err := MakeWith(PackageDir, ".env.example")

	if err != nil {
		t.Errorf("%v", err)
		t.FailNow()
	}

	if env.GetRootDir() != PackageDir {
		t.Errorf("The given env root directoy is invalid")
	}

	if env.Get("FILE_NAME") != ".env.example" {
		t.Errorf("The given env file is invalid")
	}
}

func TestItFallbackToOsEnvIfKeysAreNotFound(t *testing.T) {
	t.Parallel()

	key := "GITHUB_HANDLER"
	value := "gocanto"

	err := os.Setenv(key, value)

	if err != nil {
		t.Errorf("%v", err)
		t.FailNow()
	}

	env, _ := Make(PackageDir)

	if env.GetRootDir() != PackageDir {
		t.Errorf("The given env root directoy is invalid")
	}

	if env.Get(key) != value {
		t.Errorf("The given key [%s] is invalid", key)
	}
}
