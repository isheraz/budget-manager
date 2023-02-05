package main

import (
	"expense-manager/m/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	request := gin.Default()
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println(db, err)

	request.GET("/", func(cntx *gin.Context) {

		expense := models.Expense{ID: 0, Amount: 45.552, Description: "test description", CategoryID: 1}
		fmt.Printf("expense: %v\n", expense)
		cntx.JSON(http.StatusOK, gin.H{"data": db})
	})

	request.Run()
}
