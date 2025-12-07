package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"
)

var dataPath string = "/quedata"

var sentPath string = "sent.json"
var displayPath string = "display.json"

func ensureFiles() {
	fmt.Println(dataPath)

	requiredFiles := []string{sentPath, displayPath}

	for _, r := range requiredFiles {
		fullPath := path.Join(dataPath, r)

		_, stat := os.Stat(fullPath)

		if os.IsNotExist(stat) {
			os.WriteFile(fullPath, []byte("[]"), 0755)
		}
	}
}

func writeQuestion(q Question) int {
	ensureFiles()

	filePath := path.Join(dataPath, sentPath)

	var oldData []Question

	fileData, _ := os.ReadFile(filePath)

	err := json.Unmarshal(fileData, &oldData)

	if err != nil {
		fmt.Println("[ERROR] sent.json is not a valid question array")
		return http.StatusInternalServerError
	}

	oldData = append(oldData, q)

	newData, _ := json.Marshal(oldData)

	os.WriteFile(filePath, newData, 0755)

	return http.StatusOK
}

func getDisplayData() []byte {
	ensureFiles()

	filePath := path.Join(dataPath, displayPath)

	fileData, _ := os.ReadFile(filePath)

	return fileData
}
