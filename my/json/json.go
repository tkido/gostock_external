package json

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Save data as json
func Save(path string, v interface{}) error {
	json, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, json, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// Load data as json
func Load(path string, v interface{}) error {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bs, v)
	if err != nil {
		return err
	}
	return nil
}
