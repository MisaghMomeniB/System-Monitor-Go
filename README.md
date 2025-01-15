# 📊 Go System Monitoring Tool

A lightweight system monitoring tool built with Go. This application monitors CPU, memory, and disk usage and serves the results as a JSON response via an HTTP server.

## 🚀 Features

- 📈 **Real-time Monitoring:** Tracks CPU, memory, and disk usage.
- 🌐 **HTTP Server:** Exposes system metrics on a RESTful API endpoint.
- ⚡ **Efficient & Fast:** Built with Go for high performance.
- 📦 **Portable:** Easily deployable on servers and workstations.

## 📦 Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/yourusername/go-system-monitor.git
   cd go-system-monitor
   ```
2. **Install dependencies:**
   ```bash
   go mod tidy
   ```
3. **Run the application:**
   ```bash
   go run main.go
   ```

## 📡 Usage

### Start the Server
The server will start on `http://localhost:8080/stats`

### API Endpoint
- **GET** `/stats` - Returns current CPU, memory, and disk usage in JSON format.

Example Response:
```json
{
  "cpu_usage": 15.3,
  "memory_usage": "2.34 GB",
  "disk_usage": "120.45 GB"
}
```

## 📚 Code Overview

- **`main.go`** - The core logic of the system monitoring tool.
- **Key Functions:**
  - `getSystemStats()` - Gathers system stats using `gopsutil`
  - `formatBytes()` - Converts bytes to a human-readable format
  - `statsHandler()` - Serves system stats as a JSON response

## 🛠️ Built With

- **Go** 🐹
- **gopsutil** - For system metrics

## 🐛 Contributing

Contributions are welcome! To contribute:
1. Fork the repository.
2. Create a feature branch.
3. Commit your changes.
4. Create a pull request.

## 📄 License

This project is licensed under the MIT License. See the `LICENSE` file for more details.

---

💡 **Happy Monitoring!** 📊
