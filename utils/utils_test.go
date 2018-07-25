package utils

import (
	"math"
	"strconv"
	"testing"
)

func TestAddressToScripthash(t *testing.T) {

	expected := "0xa7274594ce215208c8e309e8f2fe05d4a9ae412b"
	actual := BigOrLittle(AddressToScripthash("AKibPRzkoZpHnPkF6qvuW2Q4hG9gKBwGpR"))

	if actual != expected {
		t.Errorf("Expected the  to be %s but instead got %s!", expected, actual)
	}
}

func TestHexToNumStr(t *testing.T) {
	const expect = "20987b86fe00"

	hexStr := BigOrLittle([]byte(expect))

	// a := hex.EncodeToString([]byte(hexStr))

	// t.Log("a", a)
	// num, _ := strconv.Atoi("0x0bebc200")
	// t.Log("hexStr", hexStr)

	// t.Log("num", num)

	x, _ := strconv.ParseInt(hexStr, 16, 64)
	t.Log("x", x)
	r := float64(x) / math.Pow10(8)

	t.Log("r", r)
}
