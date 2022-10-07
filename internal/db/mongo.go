package db

import (
	esusuMongo "github.com/Esusu2017/rrs-commons/mongo"
	"github.com/mmuoDev/go-tutorial/internal"
	"github.com/pkg/errors"
)

//AddBook adds a book to the DB
func AddBook(provideDB esusuMongo.DbProviderFunc) AddBookFunc {
	return func(b internal.AddBookRequest) error {
		col := esusuMongo.NewCollection(provideDB, collectionName)
		_, err := col.Insert(b)
		if err != nil {
			return errors.Wrapf(err, "db - unable to insert book with req%+v", b)
		}
		return nil
	}
}
