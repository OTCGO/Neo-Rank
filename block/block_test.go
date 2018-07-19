package block

import (
	"testing"
)

func TestGetNep5Balance(t *testing.T) {
	var b = &Block{}
	body, _ := b.GetNep5Balance()

	t.Log("body", body)
}
