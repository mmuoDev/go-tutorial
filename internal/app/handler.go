package app

import (
	"encoding/json"
	"net/http"

	"github.com/mmuoDev/go-tutorial/internal/db"
	"github.com/mmuoDev/go-tutorial/internal/workflow"
	"github.com/mmuoDev/go-tutorial/pkg"
)

//AddBookHandler handles an http request to add a book
func AddBookHandler(addBook db.AddBookFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req pkg.AddBookRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		addBook := workflow.AddBook(addBook)
		err := addBook(req)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}
