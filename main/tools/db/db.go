package db

import (
	"time"

	"github.com/boltdb/bolt"
)

type CDB struct {
	strFile string
	pdb     *bolt.DB
}

func NewDB(path string) *CDB {
	db, err := bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic("DB falied")
	}
	return &CDB{strFile: path, pdb: db}
}
func (db CDB) Write(f func(*bolt.Tx) bool) {
	dbtx, err := db.pdb.Begin(true)
	success := false
	if err != nil {
		panic("read failed")
	}
	defer func() {
		if success {
			dbtx.Commit()
		} else {
			dbtx.Rollback()
		}
	}()
	success = f(dbtx)

}
func (b CDB) Read()   {}
func (b CDB) Erase()  {}
func (b CDB) Exists() {}
