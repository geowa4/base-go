package foos

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

type fooDataLoader interface {
	loadOne(fooID int) ([]*foo, error)
	loadAll() ([]*foo, error)
	loadOneWithBars(fooID int) ([]*foo, error)
	loadAllWithBars() ([]*foo, error)
}

type fooResolver struct {
	accessor fooDataLoader
	fooID    int
	loadOne  bool
}

func (f *fooResolver) shouldLoadBars(fieldASTs []*ast.Field) bool {
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

func (f *fooResolver) loadFoosWithBars() ([]*foo, error) {
	if f.loadOne {
		return f.accessor.loadOneWithBars(f.fooID)
	}
	return f.accessor.loadAllWithBars()
}

func (f *fooResolver) loadFoos() ([]*foo, error) {
	if f.loadOne {
		return f.accessor.loadOne(f.fooID)
	}
	return f.accessor.loadAll()
}

func (f *fooResolver) resolve(p graphql.ResolveParams) (interface{}, error) {
	fooIDParam, fooIDParamOK := p.Args["id"].(int)
	if fooIDParamOK {
		f.fooID = fooIDParam
		f.loadOne = true
	}

	shouldLoadBars := f.shouldLoadBars(p.Info.FieldASTs)
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
