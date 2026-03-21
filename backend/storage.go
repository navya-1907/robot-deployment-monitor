package main

import (
	"encoding/json"
	"os"
)

func saveLogs() {

	file, _ := os.Create("../data/robot_logs.json")

	json.NewEncoder(file).Encode(robotLogs)

	file.Close()
}
