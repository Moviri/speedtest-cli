package report

import (
	"time"
)

// CSVReport represents the output data fields in a CSV file
type CSVReport struct {
	Timestamp time.Time `csv:"Timestamp"`
	Name      string    `csv:"Server Name"`
	Address   string    `csv:"Address"`
	Ping      float64   `csv:"Ping"`
	Jitter    float64   `csv:"Jitter"`
	// Download bandwidth in bytes per second
	Download int64 `csv:"Download"`
	// Upload bandwidth in bytes per second
	Upload int64  `csv:"Upload"`
	Share  string `csv:"Share"`
	IP     string `csv:"IP"`
}
