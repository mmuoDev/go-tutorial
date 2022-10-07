package internal

//AddBookRequest represents data needed to add a book
type AddBookRequest struct {
	Name        string `bson:"name"`
	Description string `bson:"description"`
}
