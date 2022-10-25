package bootstrap

import (
	"fmt"

	"github.com/vinoMamba/mangosteen-go/app/models/user"
	"github.com/vinoMamba/mangosteen-go/pkg/config"
	"github.com/vinoMamba/mangosteen-go/pkg/database"
	"gorm.io/driver/postgres"
)

func SetupDB() {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai",
		config.Get("db.host"),
		config.Get("db.user"),
		config.Get("db.password"),
		config.Get("db.dbname"),
		config.Get("db.port"),
	)
	database.Connect(postgres.Open(dsn))
	database.DB.AutoMigrate(&user.User{})
	fmt.Println("Database connected")
}
