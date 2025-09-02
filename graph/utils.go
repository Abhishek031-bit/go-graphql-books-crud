package graph

import (
	"go-graphql/graph/model"
	"go-graphql/models"
	"strconv"
)

func ConvertToGraphQLBook(b *models.Book) *model.Book {
	if b == nil {
		return nil
	}
	return &model.Book{
		ID:     strconv.Itoa(int(b.ID)),
		Title:  b.Title,
		Author: b.Author,
	}
}
