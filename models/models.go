package models

type Response struct {
	Timestamp int    `json:"timestamp"`
	Error     string `json:"error"`
}

type Timestamp struct {
	Timestamp int	 `json:"timestamp"`
}