package apiserver

import (
	"database/sql"
	"github.com/PesTospertnyj/http-rest-api/internal/app/store/sqlstore"
	"net/http"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer func() {
		_ = db.Close()
	}()

	store := sqlstore.New(db)
	srv := newServer(store)

	return http.ListenAndServe(config.BindAddr, srv.router)
}

func newDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
