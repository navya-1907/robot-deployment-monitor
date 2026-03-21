package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type RobotData struct {
	RobotID string  `json:"robot_id"`
	Battery int     `json:"battery"`
	Speed   float64 `json:"speed"`
	Lidar   string  `json:"lidar_status"`
	X       int     `json:"x"`
	Y       int     `json:"y"`
}

func main() {

	robots := []string{"AMR-01", "AMR-02", "AMR-03"}

	for {

		rand.Seed(time.Now().UnixNano())

		data := RobotData{
			RobotID: robots[rand.Intn(len(robots))],
			Battery: rand.Intn(80) + 20,
			Speed:   rand.Float64()*1.2 + 0.3,
			Lidar:   randomLidar(),
			X:       rand.Intn(50),
			Y:       rand.Intn(50),
		}

		sendTelemetry(data)

		time.Sleep(2 * time.Second)
	}
}

func randomLidar() string {

	options := []string{"OK", "NOISE"}

	return options[rand.Intn(len(options))]
}

func sendTelemetry(data RobotData) {

	jsonData, _ := json.Marshal(data)

	http.Post(
		"http://localhost:8080/robot-data",
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	fmt.Println("Telemetry sent:", data)
}
