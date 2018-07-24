package block

import (
	"Neo-Rank/utils"
	"testing"
)

func TestGetNep5Balance(t *testing.T) {
	var b = &Block{}
	body, _ := b.GetNep5Balance("0xecc6b20d3ccac1ee9ef109af5a7cdb85706b1df9", utils.BigOrLittle(utils.AddressToScripthash("AKibPRzkoZpHnPkF6qvuW2Q4hG9gKBwGpR")))

	t.Log("body", body)

}
