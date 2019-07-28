package db

import (
    "database/sql"

    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"

    "github.com/paxapy/goods/internal/model"
)

const boat = "boat"
const page = "page"

type Config struct {
    ConnectString string
}

func InitDb(cfg Config) (*pgDb, error) {
    if dbConn, err := sqlx.Connect("postgres", cfg.ConnectString); err != nil {
        return nil, err
    } else {
        pg := &pgDb{dbConn: dbConn}
        if err := pg.dbConn.Ping(); err != nil {
            return nil, err
        }
        if err := pg.createTablesIfNotExist(); err != nil {
            return nil, err
        }
        if err := pg.prepareSqlStatements(); err != nil {
            return nil, err
        }
        return pg, nil
    }
}

type pgDb struct {
    dbConn *sqlx.DB
    sqlList map[string]*sqlx.Stmt
    sqlSelectItem *sql.Stmt
}

var createSql = map[string]string {
    boat: `CREATE TABLE IF NOT EXISTS boats (
     id SERIAL NOT NULL PRIMARY KEY,
     title VARCHAR (42) NOT NULL,
     slug VARCHAR (42) NOT NULL UNIQUE,
     type VARCHAR (13) NOT NULL UNIQUE,
     length REAL,
     width REAL,
     height REAL,
     thickness REAL,
     weight REAL,
     capacity INTEGER
    );`,
    page: `CREATE TABLE IF NOT EXISTS pages (
     id SERIAL NOT NULL PRIMARY KEY,
     title VARCHAR (42) NOT NULL,
     slug VARCHAR (42) NOT NULL UNIQUE,
     body TEXT NOT NULL
    );`,
}

var selectSql = map[string]string {
    boat: "SELECT * FROM boats",
    page: "SELECT id, title, slug, body FROM pages",
}

func (pg *pgDb) createTablesIfNotExist() error {
    for _, q := range createSql {
        if rows, err := pg.dbConn.Query(q); err != nil {
            return err
        } else {
            rows.Close()
        }
    }
    return nil
}

func (pg *pgDb) prepareSqlStatements() (err error) {
    pg.sqlList = make(map[string]*sqlx.Stmt)
    for k, q := range selectSql {
        if pg.sqlList[k], err = pg.dbConn.Preparex(q); err != nil {
            return err
        }
    }
    return nil
}

func (pg *pgDb) SelectBoats() ([]*model.Boat, error) {
    boats := make([]*model.Boat, 0)
    if err := pg.sqlList[boat].Select(&boats); err != nil {
        return nil, err
    }
    return boats, nil
}

func (pg *pgDb) SelectPages() ([]*model.Page, error) {
    pages := make([]*model.Page, 0)
    if err := pg.sqlList[page].Select(&pages); err != nil {
        return nil, err
    }
    return pages, nil
}
