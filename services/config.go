package services

import (
	"fmt"
	"os"
	"path"

	. "github.com/fagom/gitc/internal"
)

func GetConfig() string {
	dir, err := os.UserHomeDir()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = os.Mkdir(path.Join(dir, RootFolder), 0755)

	configFilePath := path.Join(dir, RootFolder, ConfigFileName)

	fileInfo, err := os.Stat(configFilePath)

	if fileInfo == nil {
		fmt.Println("config file does not exist. Creating file...")
		err = os.WriteFile(configFilePath, nil, 0644)
	}

	return configFilePath
}
