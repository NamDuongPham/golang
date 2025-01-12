package main

import (
	"fmt"
	"log"
	"my-app/common"
	"my-app/modules/items/model"
	ginitem "my-app/modules/items/transport/gin"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	fmt.Println(db)
	// now := time.Now().UTC()
	// item := TodoItem{
	// 	Id:          1,
	// 	Title:       "Learn Go",
	// 	Description: "Learn Go programming language",
	// 	Status:      "Doing",
	// 	CreatedAt:   &now,
	// 	UpdatedAt:   &now,
	// }
	// jsonData, err := json.Marshal(item)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(jsonData))

	// jsonStr := "{\"id\":1,\"title\":\"Learn Go\",\"description\":\"Learn Go programming language\",\"status\":\"Doing\",\"created_at\":\"2021-09-01T00:00:00Z\",\"updated_at\":null}"
	// var item2 TodoItem
	// err = json.Unmarshal([]byte(jsonStr), &item2)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(item2)

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", ginitem.CreateItem(db))
			items.GET("", ListItem(db))
			items.GET("/:id", GetDetailItem(db))
			items.PATCH("/:id", UpdateItem(db))
			items.DELETE("/:id", DeleteItem(db))
		}
	}
	r.Run()
}

//	func CreateItem(db *gorm.DB) func(c *gin.Context) {
//		return func(c *gin.Context) {
//			var data model.TodoItemCreate
//			// ShouldBind gọi thàm UnmarshalJson
//			if err := c.ShouldBind(&data); err != nil {
//				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//				return
//			}
//			db.Create(&data)
//			c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
//		}
//	}
func GetDetailItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoItem
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("id = ?", id).First(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
func UpdateItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoItemUpdate
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
func DeleteItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Table(model.TodoItem{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
			"status": "Delete ",
		}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
func ListItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		paging.Process()
		var result []model.TodoItem
		db := db.Where("status <> ?", "Delete")
		if err := db.Table(model.TodoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Order("id desc").
			Offset((paging.Page - 1) * paging.Limit).
			Limit(paging.Limit).
			Find(&result).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
