package routers
import (
	"github.com/gin-gonic/gin"
	"github.com/koushikvarm/week1-GL1-CipherSchools/database"
	"github.com/koushikvarm/week1-GL1-CipherSchools/handler"
)
func Routers() *gin.Engine {
	router:= gin.Default() // get the default engine for further customization
	api:=handler.Handler{
		DB:database.GetDB(), // set the handler DB
	}

	router.GET("/books",api.GetBooks) //set the function for https://localhost:8000/books : Get request
	//while calling handler function, gin will pass *gin.Context in the function
	router.POST("/book", api.SaveBook)
	return router
	

}
