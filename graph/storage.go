package graph

import "go-graphql/graph/model"

// ---- In-memory "database" ----
var Books []*model.Book
var IdCounter = 1
