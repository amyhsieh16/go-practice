package initializers

import "gin-mssql/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}