package connect

import (
	"fmt"
	"log"
	"server/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var connection *gorm.DB
const username string = "postgres"
const password string = "password"
const database string = "apigo"
const port int = 5432
const host string = "localhost"
func InitDB() *gorm.DB{
	connectionStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",host,username,password,database,port,"disable")
	connection = connectORM(connectionStr)
	log.Println("Connected to database")
	return connection
}
func CloseConnection(){
	db,err := connection.DB()
	if err != nil{
		log.Fatal(err)
		return
	}
	log.Println("Closing connection")
	db.Close()
}
func connectORM(connectionStr string)*gorm.DB{
	connectionDB,err := gorm.Open(postgres.Open(connectionStr),&gorm.Config{})
	if err != nil{
		log.Fatal(err)
		return nil
	}
	return connectionDB
}
func GetUser(id string)utils.User{
	user := utils.User{}
	connection.Where("id = ?",id).First(&user)
	return user
}

func CreateUser(newUser utils.User) utils.User{
  connection.Create(&newUser)
	return newUser
}

func UpdateUser(id string,user utils.User) utils.User{
	currentUser := utils.User{}
	connection.Where("id = ?",id).First(&currentUser)
	currentUser.Username = user.Username
	currentUser.First_Name = user.First_Name
	currentUser.Last_Name = user.Last_Name
	connection.Save(&currentUser)
	return currentUser
}
func DeleteUser(id string){
	user := utils.User{}
	connection.Where("id = ?",id).First(&user)
	connection.Delete(&user)
}