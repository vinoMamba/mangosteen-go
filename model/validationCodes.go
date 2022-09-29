package model

type ValidationCode struct {
	ID        uint   `json:"id"`
	Code      string `json:"code"`
	ExpiredAt int64  `json:"expired_at"`
}
