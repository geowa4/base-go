package foos

import (
	"github.com/graphql-go/graphql/language/ast"
	"github.com/jmoiron/sqlx"
)

type fooDataLoader interface {
	loadOne(fooID int) ([]*foo, error)
	loadAll() ([]*foo, error)
	loadOneWithBars(fooID int) ([]*foo, error)
	loadAllWithBars() ([]*foo, error)
}

type fooLoadResolver struct {
	accessor fooDataLoader
	fooID    int
	loadOne  bool
}

func newLoadResolver(db *sqlx.DB) *fooLoadResolver {
	return &fooLoadResolver{
		accessor: &fooDataAccessor{
			db: db,
		},
	}
}

func (f *fooLoadResolver) shouldLoadBars(fieldASTs []*ast.Field) bool {
	for _, fieldAST := range fieldASTs {
		for _, selection := range fieldAST.SelectionSet.Selections {
			field, ok := selection.(*ast.Field)
			if !ok {
				continue
			}
			if field.Name.Value == "bars" {
				return true
			}
		}
	}
	return false
}

func (f *fooLoadResolver) loadFoosWithBars() ([]*foo, error) {
	if f.loadOne {
		return f.accessor.loadOneWithBars(f.fooID)
	}
	return f.accessor.loadAllWithBars()
}

func (f *fooLoadResolver) loadFoos() ([]*foo, error) {
	if f.loadOne {
		return f.accessor.loadOne(f.fooID)
	}
	return f.accessor.loadAll()
}

func (f *fooLoadResolver) query(args map[string]interface{}, fieldASTs []*ast.Field) (interface{}, error) {
	fooIDParam, fooIDParamOK := args["id"].(int)
	if fooIDParamOK {
		f.fooID = fooIDParam
		f.loadOne = true
	}

	shouldLoadBars := f.shouldLoadBars(fieldASTs)
	var (
		foos []*foo
		err  error
	)
	if shouldLoadBars {
		foos, err = f.loadFoosWithBars()
	} else {
		foos, err = f.loadFoos()
	}
	if err != nil {
		return nil, err
	}

	return foos, nil
}
