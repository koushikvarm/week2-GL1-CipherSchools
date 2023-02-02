// db is connection between golang application and database application, and gave control to the DB variable
// everything is controlled by DB which is in line 26
// gorm is a library
package database
import (
	"log"
	"github.com/koushikvarm/week1-GL1-CipherSchools/models"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)
var DB *gorm.DB

func GetDB()*gorm.DB{
	return DB
}
func Setup() {
	host:= "localhost"
	port:= "5432"
	dbName:= "book"
	username:= "postgres"
	password:= "postgres"
	arg:="host="+host+"port="+ port+ "user=" + username +"dbname="+ dbname+"sslmode=disable password="+password
	db,err:=gorm.Open("postgres",args)
	if err!=nil{
		log.Fatal(err)
	}
	db.AutoMigrate(models.Book{})
    DB = db
}
