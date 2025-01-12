package ginitem

import (
	"my-app/common"
	"my-app/modules/items/biz"
	"my-app/modules/items/model"
	"my-app/modules/items/storages"
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

// xử lý đầu tiên
func CreateItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoItemCreate
		// ShouldBind gọi thàm UnmarshalJson
		//Gọi ShouldBind để parse JSON từ request body vào struct TodoItemCreate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storages.NewSQLStore(db)
		bussiness := biz.NewCreateItemBiz(store)
		if err := bussiness.CreateNewItem(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
