package initializers

import "go/jwt-auth/models"

func SyncData() {
	DB.AutoMigrate(&models.User{})
}
