package workflow

import "github.com/mmuoDev/go-tutorial/pkg"

//AddBookFunc returns functionality that adds a book
type AddBookFunc func(req pkg.AddBookRequest) error
