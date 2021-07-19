package session

import (
	"KayORM/log"
	"database/sql"
	"fmt"
	"strings"
)

// Session 核心结构
type Session struct {
	db      *sql.DB
	sql     strings.Builder
	sqlVars []interface{}
}

func New(db *sql.DB) *Session {
	return &Session{db: db}
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
}

// DB can be used concurrency
func (s *Session) DB() *sql.DB {
	return s.db
}

// Raw set sql & vars to Session.
func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(fmt.Sprintf("%s ", sql))
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

// 封装
// 1. 统一打印日志
// 2. 执行完成后，清空 Session, 这样 Session 可以复用，**开启一次会话，可以执行多次 SQL**

// Exec raw with sqlVars
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

// QueryRow gets a record from db
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

// QueryRows gets a list of records from db
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}


