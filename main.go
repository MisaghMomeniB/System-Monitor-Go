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