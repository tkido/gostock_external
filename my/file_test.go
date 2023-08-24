package my

import (
	"testing"
)

func TestExists(t *testing.T) {
	if !Exists(".") {
		t.Error("current dir must Exists!!")
	}
	if Exists("./file_not_exists") {
		t.Error("this file must not Exists!!")
	}
}

func TestExistsDir(t *testing.T) {
	if !ExistsDir(".") {
		t.Error("current dir must Exists!!")
	}
	if ExistsDir("./dir_not_exists") {
		t.Error("this file must not Exists!!")
	}
}
