package edinet

import (
	"fmt"
	"testing"
)

func TestIsUpdated(t *testing.T) {
	codes := []string{"4235", "6200"}
	for _, code := range codes {
		updated, err := IsUpdatedIn3Days(code)
		if err != nil {
			t.Error(err)
			continue
		}
		if updated {
			fmt.Printf("%s が3日以内に更新されています！\n", code)
		}
	}
	for _, code := range codes {
		updated, err := IsUpdatedIn6Months(code)
		if err != nil {
			t.Error(err)
			continue
		}
		if updated {
			fmt.Printf("%s が6ヶ月以内に更新されています！\n", code)
		} else {
			t.Error("あるべき更新が検知できません。")
			continue
		}
	}
}

func TestSessionKeyIsLive(t *testing.T) {
	updated, err := IsUpdatedIn6Months("6200")
	if err != nil {
		t.Error(err)
		return
	}
	if !updated {
		t.Error("SESSIONKEYの有効期限切れと考えられます。更新してください。")
	}
}
