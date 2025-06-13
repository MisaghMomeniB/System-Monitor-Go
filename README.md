# 📊 System Monitor (Go)

A lightweight and interactive **terminal-based system monitoring tool** written in Go. Provides real-time metrics for CPU, memory, disk, network, and processes—ideal for quick CLI insights or embedding into your DevOps tooling.

---

## 📋 Table of Contents

1. [Overview](#overview)  
2. [Features](#features)  
3. [Tech & Requirements](#tech--requirements)  
4. [Installation & Build](#installation--build)  
5. [Usage Examples](#usage-examples)  
6. [Code Structure](#code-structure)  
7. [Under the Hood](#under-the-hood)  
8. [Future Enhancements](#future-enhancements)  
9. [Contributing](#contributing)  
10. [License](#license)

---

## 💡 Overview

This CLI tool functions like a lightweight `top` or `htop`, showing essential system stats—CPU load, memory usage, disk I/O, network throughput, and running processes—all in a clean terminal UI. It uses Go libraries to achieve performance and portability. :contentReference[oaicite:1]{index=1}

---

## ✅ Features

- 🧠 Real-time monitoring of:
  - CPU usage (per core & total)  
  - RAM and swap usage  
  - Disk read/write rates  
  - Network bandwidth  
  - Live process list, sortable by CPU/memory  
- 🎨 Interactive terminal UI—refresh rate and keybindings  
- 🔁 Minimal dependencies—built using Go’s standard libraries and pprof-compatible design  
- 🌍 Cross-platform support (Linux/macOS; Windows experimental)

---

## 🛠️ Tech & Requirements

- Go **1.18+** (modules enabled)  
- Libraries:
  - `gopsutil` for system metrics (CPU, memory, disk, network)  
  - `termbox` or `tview` (or similar) for interactive TUI  
- No external vendor dependencies

---

## ⚙️ Installation & Build

```bash
git clone https://github.com/MisaghMomeniB/System-Monitor-Go.git
cd System-Monitor-Go

# Build the CLI
go build -o sysmon main.go
# Or run directly:
go run main.go
````

---

## 🚀 Usage Examples

Run the monitor:

```bash
./sysmon
```

Example key interactions:

* Press `q` or `Ctrl+C` to exit
* Use arrow keys to scroll through processes
* Press `c`, `m`, or `p` to sort by CPU, memory, or PID respectively
  *(Keybindings may vary; check UI header)*

---

## 📁 Code Structure

```
System-Monitor-Go/
├── main.go         # Entrypoint + TUI loops
├── metrics.go      # CPU, memory, disk, network retrieval
├── ui.go           # Terminal UI rendering & input handling
└── utils.go        # Helper functions (formatting, sorting)
```

---

## 🔍 Under the Hood

* Uses `gopsutil` to retrieve system statistics efficiently
* Renders interactive UI using a light terminal library (e.g., `tview`)—redrawing every second
* Shows process lists with sorting/filtering enabled
* Can expose internal metrics for integration with Prometheus or OpenTelemetry ([omgubuntu.co.uk][1], [tecmint.com][2], [en.wikipedia.org][3], [uptrace.dev][4], [github.com][5], [dev.to][6], [en.wikipedia.org][7], [medium.com][8])

---

## 💡 Future Enhancements

* ✅ Export system metrics to Prometheus/Grafana
* 🖼️ Add customizable widgets & dashboards
* ⏱️ Enable real-time alerts (e.g., high CPU or memory usage)
* 🧭 Add command-line flags to filter top processes, set refresh intervals
* 🌐 Build a web or desktop dashboard frontend

---

## 🤝 Contributing

Contributions are welcome! Consider adding:

* More metrics (e.g., GPU, temperature)
* Cross-platform support
* Persistent data/snapshots
* Modular plugin architecture

To contribute:

1. Fork the project
2. Create a `feature/...` branch
3. Submit a PR with your changes

---

## 📄 License

Licensed under the **MIT License** — see `LICENSE` for details.
