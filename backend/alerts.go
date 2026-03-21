package main

import "fmt"

func checkAlerts(data RobotData) {

	if data.Battery < 30 {

		fmt.Println("ALERT:", data.RobotID, "battery low")
	}

	if data.Lidar == "NOISE" {

		fmt.Println("ALERT:", data.RobotID, "LiDAR interference")
	}

	if data.Speed < 0.2 {

		fmt.Println("ALERT:", data.RobotID, "robot possibly stuck")
	}
}
