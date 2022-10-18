package user

import "github.com/vinoMamba/mangosteen-go/app/models"

type User struct {
	models.BaseModel
	Name   string `json:"name,omitempty"`
	Email  string `json:"-"`
	Passwd string `json:"-"`
	models.CommonTimestampsField
}
