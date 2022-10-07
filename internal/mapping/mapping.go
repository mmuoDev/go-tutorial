package mapping

import (
	"github.com/mmuoDev/go-tutorial/internal"
	"github.com/mmuoDev/go-tutorial/pkg"
)

//ToDBAddBook maps req to DB ....
func ToDBAddBook(req pkg.AddBookRequest) internal.AddBookRequest {
	return internal.AddBookRequest{
		Name:        req.Name,
		Description: req.Description,
	}
}
