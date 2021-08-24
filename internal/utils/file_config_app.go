package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type AppConfig struct {
	Name        string `json:"name"`
	Environment string `json:"environment"`
	Debug       bool   `json:"debug"`
}

type fileFuncType func(string) (AppConfig, error)

func OpenFileContinuingCycle(path string) {
	readFileFunc := OpenReadJsonFile()

	for {
		result, err := readFileFunc(path)
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Println(result)
		time.Sleep(2 * time.Second)
	}
}

func OpenFileInCycleByCount(path string) ([]AppConfig, error) {
	list := make([]AppConfig, 5)

	readFileFunc := OpenReadJsonFile()

	for i := 0; i < 5; i++ {
		result, err := readFileFunc(path)
		if err == nil {
			list[i] = result
		} else {
			return nil, err
		}
	}

	return list, nil
}

func OpenReadJsonFile() fileFuncType {
	return func(path string) (AppConfig, error) {
		var result AppConfig
		jsonFile, err := os.Open(path)
		if err != nil {
			return result, fmt.Errorf("not found file in path: %s", path)
		}
		if jsonFile == nil {
			return result, fmt.Errorf("not found file in path: %s", path)
		}
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		data := []byte(byteValue)
		err = json.Unmarshal(data, &result)
		if err != nil {
			return result, fmt.Errorf("Not parsed file in path: %s", path)
		}

		if result.Environment == "" {
			return result, fmt.Errorf("not found struct 'AppConfig' for file on path: %s", path)
		}

		return result, nil
	}
}
