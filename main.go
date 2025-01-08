package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TodoItem struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	// Image string `json:"image"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

func main() {

	now := time.Now().UTC()
	item := TodoItem{
		Id:          1,
		Title:       "Learn Go",
		Description: "Learn Go programming language",
		Status:      "Doing",
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}
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
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": item,
		})
	})
	r.Run()
}
