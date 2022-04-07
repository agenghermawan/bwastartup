package main

import (
	"bwastartup/user"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	user := user.Users{
		Name: "Tes Simpan",
	}
	userRepository.Save(user)
	//router := gin.Default()
	//router.GET("/users", handler)
	//router.Run()
}

func handler(c *gin.Context) {
	//dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//
	//var users []user.Users
	//db.Find(&users)
	//
	//c.JSONP(http.StatusOK, users)
}
