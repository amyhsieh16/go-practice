package initializers

import(
	"os"

	"gorm.io/gorm"
	"gorm.io/driver/sqlserver"
)

var DB *gorm.DB

func ConnectToDb(){
	var err error

	dsn := os.Getenv("DB")
	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	// database, err := gorm.Open(sqlserver.Open("server=localhost;user id=sa;password=123456;database=test"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	
}