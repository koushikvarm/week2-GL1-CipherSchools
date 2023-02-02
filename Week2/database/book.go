// in this wrote functions which is used to fetch data from database

package database

import (
	"github.com/jinzhu/gorm"
	"github.com/koushikvar/week1-GL1-CipherSchools/models"
	"github.com/koushikvarm/week1-GL1-CipherSchools/models"
)

// GetBooks is creating connection and interacting from goalng app to dbserver via db variable
func GetBooks(db *gorm.DB) ([]models.Book,error){

	books:=[]models.Books{}
	query:=db.Select("books.*")
	err:= query.Find(&books).Error
	if err :=nil,err{
		return nil,err
	}
	return books,nil
}
func GetBookByID(db *gorm.DB, id string) (*models.Book, error)  {
	book := models.Book{}
	err:=db.Select("books.*").Group("boss.id").Where("books.id= ?",id).First(&book).Error
	if err!nil{
		return nil,err
	}
	return &book{}, nil
}
func DeleteBookById(db *gorm.DB, id string) error {
	var book models.Book
	err:=db.Where("id=?",id).Delete(&book).Error
	if err!nil{
		return err
	}
	return nil
}
func UpdateBookById(db *gorm.DB, book *models.Book) error {
	err:=db.Save(book).Error //book is with id
	if err!=nil{
		return err
	}
	return nil
}
func SaveBook(db *gorm.DB, book *models.Book) error {
	err:=db.Save(book).Error //bookhere is without id
	if err :=nil,err{
		return err
	}
	return nil
}

