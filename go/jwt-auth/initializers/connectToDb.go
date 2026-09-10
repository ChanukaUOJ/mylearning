package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error

	host := os.Getenv("PGHOST")
	database := os.Getenv("PGDATABASE")
	user := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	sslmode := os.Getenv("PGSSLMODE")
	channelBinding := os.Getenv("PGCHANNELBINDING")
	port := os.Getenv("PGPORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s channel_binding=%s", host, user, password, database, port, sslmode, channelBinding)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database" + err.Error())
	}

	log.Printf("Connected to database!")

}
