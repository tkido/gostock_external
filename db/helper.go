package db

import (
	"io"

	"github.com/tkido/gostock/my"
)

// Download ファイルをダウンロードする
func Download(path, url string) error {
	resp, err := my.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	w := NewWriter(path)
	defer w.Close()
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
