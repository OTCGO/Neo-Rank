package utils

import (
	"Neo-Rank/utils/base58"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
)

func AddressToScripthash(address string) []byte {

	de := base58.Decode(address)

	str := de[1 : len(de)-4]

	dst := make([]byte, hex.EncodedLen(len(str)))
	hex.Encode(dst, str)

	fmt.Printf("%s\n", dst)

	return dst
	// return binascii.hexlify(b58decode(address)[1:-4]).decode('utf-8')
}

func BigOrLittle(arr []byte) []byte {
	length := len(arr)
	for index := 0; index < length/2; index++ {
		if index%2 == 0 {
			arr[index], arr[length-2-index] = arr[length-2-index], arr[index]
		} else {
			arr[index], arr[length-index] = arr[length-index], arr[index]
		}
	}

	return arr
}

func HexToNumStr(fixied8 string, decimals int) float64 {
	hexStr := BigOrLittle([]byte(fixied8))

	x, _ := strconv.ParseInt(string(hexStr), 16, 64)

	result := float64(x) / math.Pow10(decimals)
	return result
}
