package main
import(
	"gin-mssql/controllers"
	"gin-mssql/initializers"
	// middleware "gin-mssql/milddleware"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	// r.GET("/user", controllers.GetUser) 

	r.Run()
}