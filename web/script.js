// Function to fetch system stats from the API
function fetchStats() {
    fetch('/stats')  // URL of the Go API (same server)
        .then(response => response.json())
        .then(data => {
            // Update the DOM with the fetched data
            document.getElementById('cpu-usage').textContent = data.cpu_usage + '%';
            document.getElementById('memory-usage').textContent = data.memory_usage + '%';
            document.getElementById('disk-usage').textContent = data.disk_usage + '%';
  
            // Ensure the data is parsed as floats before using in charts
            updateCharts({
                cpu_usage: parseFloat(data.cpu_usage), // Convert to float
                memory_usage: parseFloat(data.memory_usage), // Convert to float
                disk_usage: parseFloat(data.disk_usage) // Convert to float
            });
        })
        .catch(error => {
            console.error('Error fetching data:', error);
            // Handle error if the API request fails
            document.getElementById('cpu-usage').textContent = 'Error';
            document.getElementById('memory-usage').textContent = 'Error';
            document.getElementById('disk-usage').textContent = 'Error';
        });
  }
  
  // Chart.js Setup for CPU, Memory, and Disk
  const cpuCtx = document.getElementById('cpuChart').getContext('2d');
  const memoryCtx = document.getElementById('memoryChart').getContext('2d');
  const diskCtx = document.getElementById('diskChart').getContext('2d');
  
  // Initializing the charts
  const cpuChart = new Chart(cpuCtx, {
    type: 'line',
    data: {
        labels: [], // Time labels will go here
        datasets: [{
            label: 'CPU Usage',
            borderColor: '#ff6f61',
            data: [],
            fill: false,
        }]
    },
    options: {
        responsive: true,
        maintainAspectRatio: false,  // Prevent automatic resizing of the chart
        scales: {
            x: {
                type: 'linear',
                position: 'bottom'
            },
            y: {
                beginAtZero: true,
            }
        }
    }
  });
  
  const memoryChart = new Chart(memoryCtx, {
    type: 'line',
    data: {
        labels: [],
        datasets: [{
            label: 'Memory Usage',
            borderColor: '#4caf50',
            data: [],
            fill: false,
        }]
    },
    options: {
        responsive: true,
        maintainAspectRatio: false,  // Prevent automatic resizing of the chart
        scales: {
            x: {
                type: 'linear',
                position: 'bottom'
            },
            y: {
                beginAtZero: true,
            }
        }
    }
  });
  
  const diskChart = new Chart(diskCtx, {
    type: 'line',
    data: {
        labels: [],
        datasets: [{
            label: 'Disk Usage',
            borderColor: '#2196f3',
            data: [],
            fill: false,
        }]
    },
    options: {
        responsive: true,
        maintainAspectRatio: false,  // Prevent automatic resizing of the chart
        scales: {
            x: {
                type: 'linear',
                position: 'bottom'
            },
            y: {
                beginAtZero: true,
            }
        }
    }
  });
  
  // Function to update the charts with new data
  function updateCharts(data) {
    const currentTime = new Date().getTime();
    
    // Add new data to the charts (CPU, Memory, Disk)
    cpuChart.data.labels.push(currentTime);
    cpuChart.data.datasets[0].data.push(data.cpu_usage);
  
    memoryChart.data.labels.push(currentTime);
    memoryChart.data.datasets[0].data.push(data.memory_usage);
  
    diskChart.data.labels.push(currentTime);
    diskChart.data.datasets[0].data.push(data.disk_usage);
  
    // Keep the data up to a certain length (e.g., only show last 50 data points)
    if (cpuChart.data.labels.length > 50) {
        cpuChart.data.labels.shift();
        cpuChart.data.datasets[0].data.shift();
    }
    if (memoryChart.data.labels.length > 50) {
        memoryChart.data.labels.shift();
        memoryChart.data.datasets[0].data.shift();
    }
    if (diskChart.data.labels.length > 50) {
        diskChart.data.labels.shift();
        diskChart.data.datasets[0].data.shift();
    }
  
    // Update the charts
    cpuChart.update();
    memoryChart.update();
    diskChart.update();
  }
  
  // Fetch stats every 5 seconds to keep the data updated
  setInterval(fetchStats, 5000);
  
  // Initial fetch when the page loads
  fetchStats();
  