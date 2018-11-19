package foos

import "database/sql"

type fooDB interface {
	Select(dest interface{}, query string, args ...interface{}) error
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

type fooDataAccessor struct {
	db fooDB
}

const baseFooStatement string = "select foos.id, foos.name from foos"
const fooWithBarsStatement string = "select foos.id, foos.name, bars.id \"bars.id\", bars.value \"bars.value\" from foos left join bars on (foos.id = bars.foo_id)"
const loadFooByIDClause string = " where foos.id = $1"

func (fda *fooDataAccessor) loadOne(fooID int) ([]*foo, error) {
	statement := baseFooStatement + loadFooByIDClause
	var foos []*foo
	err := fda.db.Select(&foos, statement, fooID)
	if err != nil {
		return nil, err
	}
	return foos, nil
}

func (fda *fooDataAccessor) loadAll() ([]*foo, error) {
	statement := baseFooStatement
	var foos []*foo
	err := fda.db.Select(&foos, statement)
	if err != nil {
		return nil, err
	}
	return foos, nil
}

func (fda *fooDataAccessor) parseRows(rows *sql.Rows) (map[int]*foo, error) {
	var (
		fooID    int
		fooName  string
		barID    int
		barValue int
	)
	fooMap := make(map[int]*foo)
	for rows.Next() {
		err := rows.Scan(&fooID, &fooName, &barID, &barValue)
		if err != nil {
			return nil, err
		}
		if f, ok := fooMap[fooID]; ok {
			f.Bars = append(f.Bars, &bar{
				ID:    barID,
				Foo:   f,
				Value: barValue,
			})
		} else {
			newFoo := &foo{
				ID:   fooID,
				Name: fooName,
				Bars: make([]*bar, 0),
			}
			newFoo.Bars = append(newFoo.Bars, &bar{
				ID:    barID,
				Foo:   newFoo,
				Value: barValue,
			})
			fooMap[fooID] = newFoo
		}
	}
	return fooMap, nil
}

func (fda *fooDataAccessor) handleRowsWithBars(rows *sql.Rows, err error) ([]*foo, error) {
	fooMap, err := fda.parseRows(rows)
	if err != nil {
		return nil, err
	}
	return fda.mapToSlice(fooMap), nil
}

func (fda *fooDataAccessor) loadOneWithBars(fooID int) ([]*foo, error) {
	statement := fooWithBarsStatement + loadFooByIDClause
	rows, err := fda.db.Query(statement, fooID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return fda.handleRowsWithBars(rows, err)
}

func (fda *fooDataAccessor) loadAllWithBars() ([]*foo, error) {
	statement := fooWithBarsStatement
	rows, err := fda.db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return fda.handleRowsWithBars(rows, err)
}

func (fda *fooDataAccessor) mapToSlice(foosByID map[int]*foo) []*foo {
	fooSlice := make([]*foo, 0, len(foosByID))
	for _, foo := range foosByID {
		fooSlice = append(fooSlice, foo)
	}
	return fooSlice
}
