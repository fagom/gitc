package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path"

	. "github.com/fagom/gitc/models"
)

var Logger log.Logger

func BuildConfiguration() {
	dir := GetHomeDir()

	configFilePath := path.Join(dir, RootFolder, ConfigFileName)

	fileInfo, _ := os.Stat(configFilePath)

	if fileInfo == nil {
		err := os.WriteFile(configFilePath, nil, 0644)

		if err != nil {
			fmt.Println("Error creating Config file. Exiting...")
			os.Exit(1)
		}
	}

}

func GetHomeDir() string {
	dir, err := os.UserHomeDir()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = os.Mkdir(path.Join(dir, RootFolder), 0755)
	err = os.Mkdir(path.Join(dir, RootFolder, LogFolder), 0755)

	return dir
}

func OpenLogger() *os.File {
	dir := GetHomeDir()

	logFilePath := path.Join(dir, RootFolder, LogFolder, LogFileName)

	logFile, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		fmt.Println("Unable to reach log file. Exiting...")
		os.Exit(1)
	}
	Logger.SetOutput(logFile)
	Logger.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	return logFile
}

func GetConfiguration() []Config {
	dir := GetHomeDir()

	configFilePath := path.Join(dir, RootFolder, ConfigFileName)

	file, err := os.Open(configFilePath)
	defer file.Close()

	if err != nil {
		Logger.Fatal("Unable to open config file. File doesn't exist or have necessary permissions.")
		fmt.Println("Error in configuration. Please check logs for more details.")
		os.Exit(1)
	}

	var config []Config

	byteValue, err := io.ReadAll(file)

	if err != nil {
		Logger.Fatal("Unable to read config file data.")
		fmt.Println("Error in configuration. Please check logs for more details.")
		os.Exit(1)
	}

	json.Unmarshal(byteValue, &config)

	return config
}
