package model

type Error struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"error_code"`
}
