package db

import (
	"errors"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

const (
	BoltLastHashKey      = "1"
	BoltFileFormat       = "dotchain_%s.db"
	BoltBlocksBucket     = "blocks"
	BoltUTXOBucket       = "chainstate"
	BoltTxMemPool        = "txmempool"
	BoltBlockIndexBucket = "blockindex"
)

var ErrorBlockNotFound = errors.New("Block is not found")

// getDBFileName get dbfile's name with NodeID
// like dotchain_XXXXXX.db
func GetDBFileName(nodeID string) string {
	return fmt.Sprintf(BoltFileFormat, nodeID)
}

// 移除block
func RemoveBlock(db *bolt.DB, blockHash []byte) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BoltBlocksBucket))
		err := b.Delete(blockHash)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
func CreateBlockIndexBucket(db *bolt.DB) error {
	err := db.Update(func(tx *bolt.Tx) error {
		_, errc := tx.CreateBucket([]byte(BoltBlockIndexBucket))
		return errc
	})
	return err
}

// CreateBlockBucket
func CreateBlockBucket(db *bolt.DB) error {
	err := db.Update(func(tx *bolt.Tx) error {
		_, errc := tx.CreateBucket([]byte(BoltBlocksBucket))
		return errc
	})
	return err
}

// 将区块索引存入数据库
func SaveBlockIndex(db *bolt.DB, key, blockIndex []byte) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BoltBlockIndexBucket))
		err := b.Put(key, blockIndex)
		return err
	})
	return err
}
func SaveBlock(db *bolt.DB, blockHash, blockData []byte) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BoltBlocksBucket))
		err := b.Put(blockHash, blockData)
		if err != nil {
			return err
		}

		err = b.Put([]byte(BoltLastHashKey), blockHash)
		if err != nil {
			return err
		}
		return nil
	})
	//TODO:log db operate
	return err
}

// 根据hash获取区块
// 如果不存在，返回ErrorBlockNotFount
func GetBlock(db *bolt.DB, blockHash []byte) (blockData []byte, err error) {
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BoltBlocksBucket))

		blockData = b.Get(blockHash)

		if blockData == nil {
			return ErrorBlockNotFound
		}
		return nil
	})
	return
}
func GetLastBlock(db *bolt.DB) (lastHash, lastBlockData []byte, err error) {
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BoltBlocksBucket))
		lastHash = b.Get([]byte(BoltLastHashKey))
		lastBlockData = b.Get(lastHash)
		return nil
	})
	return
}
func GetLashBlockHash(db *bolt.DB) (lastHash []byte, err error) {
	lastHash, _, err = GetLastBlock(db)
	return
}
func GetTXMemPool(db *bolt.DB) ([]byte, error) {
	var txPool []byte
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BoltBlocksBucket))

		txPool = b.Get([]byte(BoltTxMemPool))
		return nil
	})
	return txPool, err
}

// SaveTXMemPool save mempool into db
func SaveTXMemPool(db *bolt.DB, txPool []byte) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BoltBlocksBucket))
		err := b.Put([]byte(BoltTxMemPool), txPool)
		if err != nil {
			return err
		}
		return nil
	})
	//TODO:log db operate
	return err
}
func CountTransactions(db *bolt.DB) int {
	counter := 0
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BoltUTXOBucket))
		c := b.Cursor()

		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			counter++
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return counter
}
