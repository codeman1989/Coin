package db

import (
	"testing"
)

func TestNewDB(t *testing.T) {

	db := NewDB("my.db")
	if db == nil {
		panic("failed")
	}

}
