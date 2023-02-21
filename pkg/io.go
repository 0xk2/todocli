package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func getDataFolder() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	dataFolder := filepath.Join(homedir, "tddata")
	err = os.MkdirAll(dataFolder, 0755)
	if err != nil {
		fmt.Println("err in ", dataFolder)
		fmt.Println(err)
	}
	return dataFolder, err
}

func saveTask(task *Task) bool {
	dataFolder, _ := getDataFolder()
	content, err := json.Marshal(task)
	if err != nil {
		fmt.Println("err in Marshal")
		fmt.Println(err)
		return false
	}
	err = ioutil.WriteFile(dataFolder+"/"+task.Id+".json", content, 0755)
	if err != nil {
		fmt.Println("err in Write file")
		fmt.Println(err)
		return false
	}
	return true
}

func LoadTask(id string) *Task {
	dataFolder, _ := getDataFolder()
	content, err := ioutil.ReadFile(dataFolder + "/" + id + ".json")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var task Task
	err = json.Unmarshal(content, &task)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &task
}
