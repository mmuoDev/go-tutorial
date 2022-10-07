package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mmuoDev/go-tutorial/internal/db"
	esusuMongo "github.com/Esusu2017/rrs-commons/mongo"
)

//App represents all functionalities in this app
type App struct {
	AddBookHandler http.HandlerFunc
}

//Handler handles http requests
func (a *App) Handler() http.HandlerFunc {
	router := httprouter.New()
	router.HandlerFunc(http.MethodPost, "/books", a.AddBookHandler)
	return http.HandlerFunc(router.ServeHTTP)
}

//OptionalArgs represents all optional args
type OptionalArgs struct {
	AddBook db.AddBookFunc
}

// Option is a representation of a function that modifies optional arguments
type Option func(oa *OptionalArgs)

func New(provideDB esusuMongo.DbProviderFunc, opts ...Option) App {
	oa := OptionalArgs{
		AddBook: db.AddBook(provideDB),
	}
	for _, opt := range opts {
		opt(&oa)
	}

	addBookHandler := AddBookHandler(oa.AddBook)
	return App{
		AddBookHandler: addBookHandler,
	}
}
