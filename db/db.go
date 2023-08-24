package db

import (
	"bytes"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/util"
	"github.com/tkido/gostock/config"
)

// LevelDB's singleton instance
var db *leveldb.DB

func init() {
	var err error
	db, err = leveldb.OpenFile(config.DBPath, nil)
	if err != nil {
		panic(err)
	}
}

// Close is wrapper of *leveldb.DB.Close
func Close() error {
	return db.Close()
}

// Has is wrapper of *leveldb.DB.Has
func Has(key string) bool {
	ret, err := db.Has([]byte(key), nil)
	if err != nil {
		panic(err)
	}
	return ret
}

// Delete is wrapper of *leveldb.DB.Delete
func Delete(key string) error {
	return db.Delete([]byte(key), nil)
}

// Get is wrapper of *leveldb.DB.Get
func Get(key string) (value []byte, err error) {
	return db.Get([]byte(key), nil)
}

// Writer is io like interface
type Writer struct {
	key string
	buf bytes.Buffer
}

// NewWriter returns new Writer
func NewWriter(key string) *Writer {
	return &Writer{
		key,
		bytes.Buffer{},
	}
}

// Write for io.Writer interface
func (w *Writer) Write(p []byte) (n int, err error) {
	return w.buf.Write(p)
}

// Close closes Writer
func (w *Writer) Close() error {
	return db.Put([]byte(w.key), w.buf.Bytes(), nil)
}

// Put is wrapper of *leveldb.DB.Put
func Put(key string, value []byte) error {
	return db.Put([]byte(key), value, nil)
}

// NewIterator with 1 argument means that it has prefix
func NewIterator(prefix string) iterator.Iterator {
	return db.NewIterator(util.BytesPrefix([]byte(prefix)), nil)
}

// PutString put string value
func PutString(key string, value string) error {
	return db.Put([]byte(key), []byte(value), nil)
}

// GetString get string value
func GetString(key string) (value string, err error) {
	bs, err := db.Get([]byte(key), nil)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}
