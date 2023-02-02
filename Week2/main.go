package main

import (
	"github.com/koushikvarm/week1-GL1-CipherSchools/database"
	"github.com/koushikvarm/week1-GL1-CipherSchools/routers"
)
func init() {
	database.Setup()
}
func init() {
	fmt.Print(1)
}
func init() {
	fmt.Print(2)
}
func main(){
	
	database.Setup() 
	engine:=routers.Router()
	err :=engine.Run("127.0.0.1:8000")
	if err!=nil {
		log.Fatal(err)
	}
	
}
