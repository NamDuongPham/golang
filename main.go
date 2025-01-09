package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TodoItem struct {
	Id    int    `json:"id" gorm:"column:id"`
	Title string `json:"title" gorm:"column:title"`
	// Image string `json:"image"`
	Description string     `json:"description" gorm:"column:description"`
	Status      string     `json:"status" gorm:"column:status"`
	CreatedAt   *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}
type TodoItemCreate struct {
	Id          int    `json:"-" gorm:"column:id"`
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	Status      string `json:"status" gorm:"column:status"`
}

func (TodoItem) TableName() string {
	return "todo_items"
}
func (TodoItemCreate) TableName() string {
	return TodoItem{}.TableName()
}
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
			items.POST("", CreateItem(db))
			items.GET("")
			items.GET("/:id", GetDetailItem(db))
			items.PATCH("/:id")
			items.DELETE("/:id")
		}
	}
	r.Run()
}
func CreateItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data TodoItemCreate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&data)
		c.JSON(http.StatusOK, gin.H{"data": data.Id})
	}
}
func GetDetailItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data TodoItem
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("id = ?", id).First(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
