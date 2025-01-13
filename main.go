// Go System Monitoring Tool
// This program monitors CPU, memory, and disk usage and serves the results via an HTTP server

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

// SystemStats holds the data for system usage
type SystemStats struct {
	CPUUsage    float64 `json:"cpu_usage"`
	MemoryUsage string  `json:"memory_usage"`
	DiskUsage   string  `json:"disk_usage"`
}

// formatBytes converts bytes to a human-readable format
func formatBytes(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	exp := int64(0)
	for n := bytes / unit; n >= unit; n /= unit {
		exp++
	}
	pre := "KMGTPE"[exp : exp+1]
	return fmt.Sprintf("%.2f %sB", float64(bytes)/float64(uint64(unit)<<exp*10), pre)
}
