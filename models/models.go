package models

type Response struct {
	Timestamp int    `json:"timestamp"`
	TimeLeft  int	 `json:"timeLeft"`
	Error     string `json:"error"`
}

type Timestamp struct {
	Timestamp int	 `json:"timestamp"`
	TimeLeft  int	 `json:"timeLeft"`
}