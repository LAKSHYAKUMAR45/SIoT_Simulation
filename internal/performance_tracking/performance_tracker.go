package performance_tracking

import (
	"fmt"
	"log"
	"os"
	"time"
)

// PerformanceMetrics represents the performance metrics of the system
type PerformanceMetrics struct {
	Throughput       float64 // Transactions processed per second
	Latency          float64 // Latency in milliseconds
	PredictionAccuracy float64 // Accuracy of predictions in percentage
	FalsePositiveRate float64 // False positive rate in percentage
	FalseNegativeRate float64 // False negative rate in percentage
}

// SimulatedEvent represents an event generated during simulation
type SimulatedEvent struct {
	EventType string    // Type of event (e.g., "Transaction")
	Timestamp time.Time // Timestamp of the event
	Data      string    // Additional data related to the event
}

// PerformanceTracker represents the performance tracking system
type PerformanceTracker struct {
	EventChannel  chan SimulatedEvent // Channel to receive simulated events
	LogFile       *os.File            // Log file to record performance metrics
	StartTime     time.Time           // Start time of performance tracking
	TotalEvents   int                 // Total number of events processed
	TotalSafe     int                 // Total number of safe transactions
	TotalMalicious int                 // Total number of malicious transactions
	TotalPredictedSafe     int         // Total number of transactions predicted safe
	TotalPredictedMalicious int         // Total number of transactions predicted malicious
}

// NewPerformanceTracker creates a new PerformanceTracker instance
func NewPerformanceTracker(logFileName string) *PerformanceTracker {
	// Open log file
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	// Initialize performance tracker
	return &PerformanceTracker{
		EventChannel: make(chan SimulatedEvent),
		LogFile:      logFile,
		StartTime:    time.Now(),
	}
}

// StartMonitoring starts monitoring performance metrics
func (pt *PerformanceTracker) StartMonitoring() {
	fmt.Println("Performance monitoring started...")
	// Start monitoring performance metrics by processing events from the event channel
	for event := range pt.EventChannel {
		// Update performance metrics based on the event
		pt.UpdateMetrics(event)
	}
}

// UpdateMetrics updates performance metrics based on the received event
func (pt *PerformanceTracker) UpdateMetrics(event SimulatedEvent) {
	// Increment total events processed
	pt.TotalEvents++
	// Calculate throughput
	elapsedTime := time.Since(pt.StartTime).Seconds()
	pt.Throughput = float64(pt.TotalEvents) / elapsedTime
	// Calculate latency
	latency := time.Since(event.Timestamp).Milliseconds()
	pt.Latency = float64(latency)
	// Log performance metrics
	pt.LogMetrics(event)
}

// LogMetrics logs performance metrics to the log file
func (pt *PerformanceTracker) LogMetrics(event SimulatedEvent) {
	// Format performance metrics
	logMsg := fmt.Sprintf("[%s] Throughput: %.2f transactions/sec, Latency: %.2f ms\n", event.EventType, pt.Throughput, pt.Latency)
	// Write to log file
	if _, err := pt.LogFile.WriteString(logMsg); err != nil {
		log.Fatalf("Failed to write to log file: %v", err)
	}
}

// StopMonitoring stops monitoring performance metrics
func (pt *PerformanceTracker) StopMonitoring() {
	// Close the event channel to stop monitoring
	close(pt.EventChannel)
	// Close log file
	if err := pt.LogFile.Close(); err != nil {
		log.Fatalf("Failed to close log file: %v", err)
	}
	fmt.Println("Performance monitoring stopped.")
}
