package pg

import (
	"github.com/go-pg/pg/extra/pgotel/v10"
	"github.com/go-pg/pg/v10"
)

func AddQueryHook(db *pg.DB) {
	db.AddQueryHook(pgotel.NewTracingHook())
}
