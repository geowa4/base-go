package foos

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type fooDataCreator interface {
	saveFoo(name string) (id int64, err error)
	saveBar(fooID, value int) (id int64, err error)
}

type createReturnValue struct {
	ID int64 `json:"id"`
}

type fooCreationResolver struct {
	creator fooDataCreator
}

func newCreateResolver(db *sqlx.DB) *fooCreationResolver {
	return &fooCreationResolver{
		creator: &fooDataAccessor{
			db: db,
		},
	}
}

func (f *fooCreationResolver) createFoo(name string) (interface{}, error) {
	id, err := f.creator.saveFoo(name)
	if err != nil {
		log.Error().Err(err).Str("name", name).Msg("Error creating Foo.")
		return nil, err
	}
	return createReturnValue{ID: id}, nil
}

func (f *fooCreationResolver) createBar(fooID, value int) (interface{}, error) {
	id, err := f.creator.saveBar(fooID, value)
	if err != nil {
		log.Error().Err(err).Int("foo_id", fooID).Int("value", value).Msg("Error creating Bar.")
		return nil, err
	}
	return createReturnValue{ID: id}, nil
}
