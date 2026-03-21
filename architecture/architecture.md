# System Architecture
```
                                                              +--------------------------+
                                                              |     Robot Simulator      |
                                                              |          (Go)            |
                                                              |--------------------------|
                                                              | - Generates telemetry    |
                                                              | - Battery, Speed         |
                                                              | - LiDAR Status           |
                                                              | - Position (X, Y)        |
                                                              +------------+-------------+
                                                                           |
                                                                           | HTTP POST (Telemetry Data)
                                                                           v
                                                      +-------------------------------------------+
                                                      |      Backend Monitoring Service (Go)      |
                                                      |-------------------------------------------|
                                                      | - Telemetry Ingestion API (/robot-data)   |
                                                      | - Health Monitoring Logic                 |
                                                      | - Alert Detection                         |
                                                      | - In-memory Data Storage                  |
                                                      +-------------------+-----------------------+
                                                                          |
                                                                          | Store Logs
                                                                          v
                                                              +--------------------------+
                                                              |     Log Storage          |
                                                              |     (JSON File)          |
                                                              |--------------------------|
                                                              | - Persistent Logs        |
                                                              | - Debugging Support      |
                                                              +------------+-------------+
                                                                           |
                                                                           | HTTP GET (/logs)
                                                                           v
                                                              +--------------------------+
                                                              |   Monitoring Dashboard   |
                                                              |     (HTML + JS)          |
                                                              |--------------------------|
                                                              | - Displays Robot Data    |
                                                              | - Battery Status Colors  |
                                                              | - LiDAR Alerts           |
                                                              | - Auto Refresh (2 sec)   |
                                                              +--------------------------+

```
