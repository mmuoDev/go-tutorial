package db

import (
	"github.com/mmuoDev/go-tutorial/internal"
)

const (
	collectionName = "books"
)

//AddBookFunc returns functionality that adds a book to the DB
type AddBookFunc func(req internal.AddBookRequest) error
