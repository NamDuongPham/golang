package biz

import (
	"context"
	"my-app/modules/items/model"
	"strings"
)

type CreateStore interface {
	CreateItem(ctx context.Context, data *model.TodoItemCreate) error
}
type createItemBiz struct {
	store CreateStore
}

func NewCreateItemBiz(store CreateStore) *createItemBiz {
	return &createItemBiz{store: store}
}
func (biz *createItemBiz) CreateNewItem(ctx context.Context, data *model.TodoItemCreate) error {
	title := strings.TrimSpace((data.Title))
	if title == "" {
		return model.ErrTitleIsBlank
	}
	if err := biz.store.CreateItem(ctx, data); err != nil {
		return err
	}
	return nil
}
