package db

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/syndtr/goleveldb/leveldb/util"
	"github.com/tkido/gostock/xbrl"
)

func TestGetString(t *testing.T) {
	key := "test_key"
	value := "test_value"
	err := PutString(key, value)
	if err != nil {
		t.Error(err)
	}
	if !Has(key) {
		t.Errorf("not saved %s", value)
	}
	saved, err := GetString(key)
	if err != nil {
		t.Error(err)
	}
	if value != saved {
		t.Errorf("not equal value %s and saved %s", value, saved)
	}
}

func TestHas(t *testing.T) {
	key := `1418\edinet\jpcrp030000-asr-001_E24512-000_2019-02-28_01_2019-05-24.xbrl`
	fmt.Println(Has(key))
	bs, err := db.Get([]byte(key), nil)
	if err != nil {
		t.Error(err)
	}

	x := &xbrl.XBRL{}
	err = xbrl.UnmarshalXBRL(x, bytes.NewBuffer(bs))
	if err != nil {
		t.Error(err)
	}

	iter := db.NewIterator(util.BytesPrefix([]byte(`1418\edinet\`)), nil)
	for iter.Next() {
		fmt.Println(string(iter.Key()))
		// iter.Value()
	}
	iter.Release()
	err = iter.Error()

	// path := `./jpcrp030000-asr-001_E24512-000_2019-02-28_01_2019-05-24.xbrl`
	// file, err := os.Create(path)
	// if err != nil {
	// 	t.Error(err)
	// }
	// defer file.Close()
	// _, err = io.Copy(file, bytes.NewBuffer(bs))
	// if err != nil {
	// 	t.Error(err)
	// }
}
