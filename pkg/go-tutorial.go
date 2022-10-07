package pkg

//AddBookRequest represents data needed to add a book
type AddBookRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
