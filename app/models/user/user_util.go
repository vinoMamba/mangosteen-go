package user

import "github.com/vinoMamba/mangosteen-go/pkg/database"

func IsEmailExist(email string) bool {
	var count int64
	database.DB.Where("email = ?", email).Count(&count)
	return count > 0
}
