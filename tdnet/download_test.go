package tdnet

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/tkido/gostock/my"
)

func TestDownload(t *testing.T) {
	dls, err := Download([]string{"6200"}, false)
	if err != nil {
		t.Error(err)
	}
	path := filepath.Join(os.TempDir(), "downloaded.html")
	err = my.WriteFileForCopyPaste(path, dls)
	if err != nil {
		t.Error(err)
	}
}
