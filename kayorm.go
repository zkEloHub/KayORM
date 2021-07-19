package KayORM

import (
	"KayORM/log"
	"KayORM/session"
	"database/sql"
)

type Engine struct {
	db *sql.DB
}


// NewEngine ...
// 1. connect to database, return *sql.DB
// 2. Ping database, keep alive
func NewEngine(driver, source string) (*Engine, error) {
	// Need to Ping after Open.
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	// Send a ping to make sure the source name is valid & init connection.
	if err = db.Ping(); err != nil {
		log.Error(err)
		return nil, err
	}
	e := &Engine{
		db: db,
	}
	log.Info("connect database success")
	return e, nil
}

// Close close the sql.DB
func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("Failed to close database")
	}
	log.Info("Close database success")
}

// NewSession Create session by Engine.
func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db)
}


