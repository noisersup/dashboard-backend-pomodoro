package models

type Response struct {
	Timestamp int    `json:"timestamp"`
	Error     string `json:"error"`
}

type TimestampDb struct {
	Timestamp int
}