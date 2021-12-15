package handler

import (
	"os"
)

const (
	DefaultFileName = ".env"
	PackageDir      = "env"
)

type Env struct {
	items    *map[string]string
	filePath string
	rootDir  string
}

func Make(rootDir string) (Env, error) {
	path := resolveFilePath(rootDir, DefaultFileName)

	return buildFrom(rootDir, path)
}

func MakeWith(rootDir string, file string) (Env, error) {
	path := resolveFilePath(rootDir, file)

	return buildFrom(rootDir, path)
}

func (current Env) Get(key string) string {
	items := *current.items
	value := items[key]

	if value == "" {
		value = os.Getenv(key)
	}

	return value
}

func (current Env) GetRootDir() string {
	return current.rootDir
}
