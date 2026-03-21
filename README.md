# 🤖 Robot Deployment Monitoring Platform

A system that simulates monitoring of Autonomous Mobile Robots (AMRs) in a warehouse environment.

This project mimics real-world robotic deployment systems by generating robot telemetry, processing it through a backend service, and visualizing it on a live dashboard.

---

## 🚀 Features

- 🤖 Simulated robot telemetry (battery, speed, LiDAR, position)
- 📡 Backend API for telemetry ingestion
- 🚨 Alert detection for abnormal conditions
- 📊 Live monitoring dashboard (auto-refresh)
- 📝 JSON-based log storage for debugging

---

## 🧠 System Architecture

The detailed system architecture and flowchart are documented here:

📄 `architecture/system_architecture.md`

---

## 🔄 Data Flow

Robot Simulator → Backend API → Alert Detection → Log Storage → Dashboard

---

## 🛠️ Tech Stack

- **Backend:** Go  
- **Simulator:** Go  
- **Frontend:** HTML, CSS, JavaScript  
- **Storage:** JSON files  

---

## 📂 Project Structure
```
robot-deployment-monitor
│
├── backend
│ ├── main.go
│ ├── handlers.go
│ ├── alerts.go
│ └── storage.go
│
├── simulator
│ └── robot_simulator.go
│
├── frontend
│ ├── index.html
│ ├── dashboard.js
│ └── style.css
│
├── data
│ └── robot_logs.json
│
├── architecture
│ └── architecture.md
│
└── README.md
```


---

## ▶️ How to Run the Project

### 1️⃣ Clone the repository
```
git clone https://github.com/navya_1907/robot-deployment-monitor.git
```
```
cd robot-deployment-monitor
```


---

### 2️⃣ Initialize Go Module (only first time)
```
go mod init robot-deployment-monitor
go mod tidy
```


---

### 3️⃣ Run Backend Server
```
cd backend
```
```
go run .
```


Server runs at:
```
http://localhost:8080/
```

---

### 4️⃣ Run Robot Simulator

Open a new terminal:
```
cd simulator
```
```
go run robot_simulator.go
```


Use Live Server Extension in VS Code to the dashboard.

---

## 📊 Dashboard Overview

The dashboard displays:

- Robot ID  
- Battery level (color-coded)  
- Speed  
- LiDAR status (OK / NOISE)  
- Position (X, Y)  

Data updates automatically every 2 seconds.

---

## 🚨 Alert Conditions

The backend detects:

- 🔴 Battery < 30% → Low Battery Alert  
- ⚠️ LiDAR = NOISE → Sensor Interference  
- 🐢 Speed < 0.2 → Possible Robot Stuck  

Alerts are printed in the backend console.

---

## 🎯 Purpose of the Project

This project was built to understand:

- Robot telemetry systems  
- Monitoring and alerting pipelines  
- Deployment-level debugging  
- System architecture design  

It simulates how real robotic fleets are monitored in automation.

---


## 👨‍💻 Author

Navya Srivastava 