package foos

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type fooDataAccessor struct {
	db *sqlx.DB
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

func (fda *fooDataAccessor) parseFoosWithBars(rows *sql.Rows) (map[int64]*foo, error) {
	var (
		fooID    int64
		fooName  string
		barID    sql.NullInt64
		barValue sql.NullInt64
	)
	fooMap := make(map[int64]*foo)
	for rows.Next() {
		err := rows.Scan(&fooID, &fooName, &barID, &barValue)
		if err != nil {
			return nil, err
		}
		if f, ok := fooMap[fooID]; ok {
			f.Bars = append(f.Bars, &bar{
				ID:    barID.Int64,
				Foo:   f,
				Value: barValue.Int64,
			})
		} else {
			newFoo := &foo{
				ID:   fooID,
				Name: fooName,
				Bars: make([]*bar, 0),
			}
			if barID.Valid {
				newFoo.Bars = append(newFoo.Bars, &bar{
					ID:    barID.Int64,
					Foo:   newFoo,
					Value: barValue.Int64,
				})
			}
			fooMap[fooID] = newFoo
		}
	}
	return fooMap, nil
}

func (fda *fooDataAccessor) handleRowsWithBars(rows *sql.Rows, err error) ([]*foo, error) {
	fooMap, err := fda.parseFoosWithBars(rows)
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
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Error().Err(err).Msg("Error closing rows when loading single foo with bars.")
		}
	}()
	return fda.handleRowsWithBars(rows, err)
}

func (fda *fooDataAccessor) loadAllWithBars() ([]*foo, error) {
	statement := fooWithBarsStatement
	rows, err := fda.db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Error().Err(err).Msg("Error closing rows when loading foos with bars.")
		}
	}()
	return fda.handleRowsWithBars(rows, err)
}

func (fda *fooDataAccessor) mapToSlice(foosByID map[int64]*foo) []*foo {
	fooSlice := make([]*foo, 0, len(foosByID))
	for _, foo := range foosByID {
		fooSlice = append(fooSlice, foo)
	}
	return fooSlice
}

func (fda *fooDataAccessor) saveFoo(name string) (id int64, err error) {
	err = fda.db.QueryRow("INSERT INTO foos(name) VALUES ($1) RETURNING id", name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (fda *fooDataAccessor) saveBar(fooID, value int) (id int64, err error) {
	err = fda.db.QueryRow("INSERT INTO bars(foo_id, value) VALUES ($1,$2) RETURNING id", fooID, value).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
