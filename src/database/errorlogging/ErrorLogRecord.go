package errorlogging

import "time"

type ErrorLogRecord struct {
	IP          string `json:"ip"`
	Application string `json:"application"`
	Location    string `json:"location"`
	Message     string `json:"message"`
	Stack       string `json:"stack"`
	CreatedOn   string `json:"created_on"`
}

func NewErrorRecord(ip string, location string, message string, stack string) *ErrorLogRecord {
	return &ErrorLogRecord{
		IP:          ip,
		Application: "MainApplication",
		Location:    location,
		Message:     message,
		Stack:       stack,
		CreatedOn:   time.Now().Format(time.RFC3339),
	}
}
