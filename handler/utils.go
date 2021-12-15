package handler

import (
	"github.com/joho/godotenv"
	"os"
	"regexp"
)

func resolveFilePath(directory string, fileName string) string {
	expression := regexp.MustCompile(`^(.*` + directory + `)`)

	cwd, _ := os.Getwd()
	path := expression.Find([]byte(cwd))

	return string(path) + `/` + fileName
}

func buildFrom(rootDir string, path string) (Env, error) {
	env := Env{}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return env, err
	}

	if items, err := godotenv.Read(path); err != nil {
		return env, err
	} else {
		env.items = &items
	}

	env.filePath = path
	env.rootDir = rootDir

	return env, nil
}
