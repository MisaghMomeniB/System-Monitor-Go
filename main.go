package main

import (
	"encoding/json" // Import the json package for encoding/decoding JSON data
	"fmt"           // Import fmt package for formatted I/O operations
	"net/http"      // Import net/http package to create HTTP servers and handle requests
	"os"            // Import os package for handling OS-level operations, like exiting
	"time"          // Import time package for time-related functions (e.g., sleep)

	"github.com/shirou/gopsutil/v3/cpu"  // Import gopsutil to get CPU usage statistics
	"github.com/shirou/gopsutil/v3/disk" // Import gopsutil to get disk usage statistics
	"github.com/shirou/gopsutil/v3/mem"  // Import gopsutil to get memory usage statistics
)

// SystemStats holds the system usage data that will be returned in the JSON response
type SystemStats struct {
	CPUUsage    float64 `json:"cpu_usage"`    // CPU usage percentage
	MemoryUsage string  `json:"memory_usage"` // Human-readable memory usage
	DiskUsage   string  `json:"disk_usage"`   // Human-readable disk usage
}

// formatBytes converts bytes to a human-readable string (e.g., "1.23 GB")
func formatBytes(bytes uint64) string {
	const unit = 1024 // Define the base unit for conversion (1024 bytes)
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes) // If less than 1024 bytes, return as bytes
	}
	exp := int64(0) // Exponent for scaling the byte count
	// Loop to find the correct scaling factor (KB, MB, GB, etc.)
	for n := bytes / unit; n >= unit; n /= unit {
		exp++
	}
	// Define the appropriate unit (K, M, G, etc.)
	pre := "KMGTPE"[exp : exp+1]
	// Format the bytes with two decimal points and the unit suffix
	return fmt.Sprintf("%.2f %sB", float64(bytes)/float64(uint64(unit)<<exp*10), pre)
}

// getSystemStats collects the current system statistics for CPU, memory, and disk usage
func getSystemStats() (SystemStats, error) {
	// Get the CPU usage percentage for the last second
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return SystemStats{}, err // Return error if unable to fetch CPU stats
	}

	// Get memory statistics (e.g., used, free, total memory)
	memStats, err := mem.VirtualMemory()
	if err != nil {
		return SystemStats{}, err // Return error if unable to fetch memory stats
	}
	// Get disk usage statistics for the root directory
	diskStats, err := disk.Usage("/")
	if err != nil {
		return SystemStats{}, err // Return error if unable to fetch disk stats
	}

	// Return the collected stats as a SystemStats struct
	return SystemStats{
		CPUUsage:    cpuPercent[0],               // Set the CPU usage
		MemoryUsage: formatBytes(memStats.Used),  // Format and set the memory usage
		DiskUsage:   formatBytes(diskStats.Used), // Format and set the disk usage
	}, nil
}

// statsHandler handles the HTTP request to fetch system stats and returns them in JSON format
func statsHandler(w http.ResponseWriter, r *http.Request) {
	stats, err := getSystemStats() // Get the current system stats
	if err != nil {
		http.Error(w, "Could not retrieve system stats", http.StatusInternalServerError) // Handle errors in fetching stats
		return
	}

	// Set the response header to indicate JSON content
	w.Header().Set("Content-Type", "application/json")
	// Encode the system stats as JSON and write to the response
	json.NewEncoder(w).Encode(stats)
}

func main() {
	// Serve static files from the "web" directory
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	// Handle requests to the "/stats" endpoint and call statsHandler to provide system stats
	http.HandleFunc("/stats", statsHandler)

	// Start the server on port 8080 and log an error if it fails
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err) // Log the error if the server cannot start
		os.Exit(1)                                 // Exit the program with an error code
	}
}
