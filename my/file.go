package my

import (
	"errors"
	"io"
	"os"
)

// Exists ファイルまたはディレクトリの存在を確認する
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// ExistsDir ディレクトリの存在を確認する
func ExistsDir(path string) bool {
	fi, err := os.Stat(path)
	if err == nil && fi.IsDir() {
		return true
	}
	return false
}

// MkDir ディレクトリが存在しなければ作成する
func MkDir(path string) (err error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			// Nothing to do. It already exists.
		} else {
			return errors.New("it already exists and isn't directory. Cannot recover")
		}
	} else {
		err = os.Mkdir(path, 0777)
		if err != nil {
			return
		}
	}
	return nil
}

// Download ファイルをダウンロードする
func Download(path, url string) error {
	resp, err := Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
