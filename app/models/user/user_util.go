package user

import (
	"fmt"

	"github.com/vinoMamba/mangosteen-go/pkg/database"
)

func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(&User{}).Where("email = ?", email).Count(&count)
	fmt.Println(count)
	return count > 0
}
