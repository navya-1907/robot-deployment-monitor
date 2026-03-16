package main

import (
	"encoding/json"
	"net/http"
)

type RobotData struct {
	RobotID string  `json:"robot_id"`
	Battery int     `json:"battery"`
	Speed   float64 `json:"speed"`
	Lidar   string  `json:"lidar_status"`
	X       int     `json:"x"`
	Y       int     `json:"y"`
}

var robotLogs []RobotData

func robotHandler(w http.ResponseWriter, r *http.Request) {

	var data RobotData

	json.NewDecoder(r.Body).Decode(&data)

	robotLogs = append(robotLogs, data)

	checkAlerts(data)

	saveLogs()
}

func logsHandler(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(robotLogs)
}
