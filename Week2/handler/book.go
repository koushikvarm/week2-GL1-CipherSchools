package handler
import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/koushikvar/week1-GL1-CipherSchools/database"
)

type Handler struct{
	DB *gorm.DB
}
func (h *Handler)GetBooks(in *gin.Context) {
	books,err:=database.GetBooks(h.DB) //h.DB is fully configures and can give the access of book table
	if err!=nil{
		log.Println(err)
		in.JSON(http.StatusInternalServerError,err)
	}
	in.JSON(http.StatusOK,books)

}
func (h *Handler) SaveBook(in *gin.Context) {
	book:=models.Book{}
	err:=in.BindJSON(&book)
	database.SaveBook(h.DB, &book)
	if err!=nil{
		log.Println(err)
		in.JSON(http.StatusInternalServerError,err)
	}
	err=database.SaveBook(h.DB, &book)
	if err!=nil{
		log.Println(err)
		in.JSON(http.StatusInternalServerError,err)
	}
	in.JSON(http.StatusOK,book)

}

