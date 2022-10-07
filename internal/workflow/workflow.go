package workflow

import (
	"github.com/mmuoDev/go-tutorial/internal/db"
	"github.com/mmuoDev/go-tutorial/internal/mapping"
	"github.com/mmuoDev/go-tutorial/pkg"
	"github.com/pkg/errors"
)

//AddBook adds a book 
func AddBook(addBook db.AddBookFunc) AddBookFunc {
	return func(req pkg.AddBookRequest) error {
		r := mapping.ToDBAddBook(req)
		if err := addBook(r); err != nil {
			return errors.Wrapf(err, "workflow - unable to add book")
		}
		return nil 
	}
}