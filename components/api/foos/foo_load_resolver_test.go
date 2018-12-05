package foos

import (
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

type mockAccessor struct {
	loadedOne      bool
	loadedWithBars bool
}

func (m *mockAccessor) loadOne(fooID int) ([]*foo, error) {
	m.loadedOne = true
	return make([]*foo, 0), nil
}

func (m *mockAccessor) loadAll() ([]*foo, error) {
	return make([]*foo, 0), nil
}

func (m *mockAccessor) loadOneWithBars(fooID int) ([]*foo, error) {
	m.loadedOne = true
	m.loadedWithBars = true
	return make([]*foo, 0), nil
}

func (m *mockAccessor) loadAllWithBars() ([]*foo, error) {
	m.loadedWithBars = true
	return make([]*foo, 0), nil
}

func TestLoadOneFoo(t *testing.T) {
	params := graphql.ResolveParams{
		Args: map[string]interface{}{
			"id": 1,
		},
	}
	mock := &mockAccessor{}
	resolver := &fooLoadResolver{accessor: mock}
	value, err := resolver.query(params.Args, params.Info.FieldASTs)
	if err != nil {
		t.Fail()
	}
	_, ok := value.([]*foo)
	if !ok {
		t.Fail()
	}
	if !mock.loadedOne {
		t.Error("loaded many foos but should have loaded one")
	}
}

func TestLoadManyFoos(t *testing.T) {
	params := graphql.ResolveParams{}
	mock := &mockAccessor{}
	resolver := &fooLoadResolver{accessor: mock}
	value, err := resolver.query(params.Args, params.Info.FieldASTs)
	if err != nil {
		t.Fail()
	}
	_, ok := value.([]*foo)
	if !ok {
		t.Fail()
	}
	if mock.loadedOne {
		t.Error("loaded one foo but should have loaded many")
	}
}

func TestLoadOneFooWithBars(t *testing.T) {
	params := graphql.ResolveParams{
		Args: map[string]interface{}{
			"id": 1,
		},
		Info: graphql.ResolveInfo{
			FieldASTs: []*ast.Field{
				&ast.Field{
					SelectionSet: &ast.SelectionSet{
						Selections: []ast.Selection{
							&ast.Field{
								Name: &ast.Name{Value: "bars"},
							},
						},
					},
				},
			},
		},
	}
	mock := &mockAccessor{}
	resolver := &fooLoadResolver{accessor: mock}
	value, err := resolver.query(params.Args, params.Info.FieldASTs)
	if err != nil {
		t.Fail()
	}
	_, ok := value.([]*foo)
	if !ok {
		t.Fail()
	}
	if !mock.loadedOne {
		t.Error("loaded many foos but should have loaded one")
	}
	if !mock.loadedWithBars {
		t.Error("did not load bars")
	}
}

func TestLoadManyFoosWithBars(t *testing.T) {
	params := graphql.ResolveParams{
		Info: graphql.ResolveInfo{
			FieldASTs: []*ast.Field{
				&ast.Field{
					SelectionSet: &ast.SelectionSet{
						Selections: []ast.Selection{
							&ast.Field{
								Name: &ast.Name{Value: "bars"},
							},
						},
					},
				},
			},
		},
	}
	mock := &mockAccessor{}
	resolver := &fooLoadResolver{accessor: mock}
	value, err := resolver.query(params.Args, params.Info.FieldASTs)
	if err != nil {
		t.Fail()
	}
	_, ok := value.([]*foo)
	if !ok {
		t.Fail()
	}
	if mock.loadedOne {
		t.Error("loaded one foo but should have loaded many")
	}
	if !mock.loadedWithBars {
		t.Error("did not load bars")
	}
}
