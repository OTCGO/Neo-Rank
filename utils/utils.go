package utils

import (
	"Neo-Rank/utils/base58"
	"encoding/hex"
	"fmt"
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

func BigOrLittle(arr []byte) string {
	length := len(arr)
	for index := 0; index < length/2; index++ {
		if index%2 == 0 {
			arr[index], arr[length-2-index] = arr[length-2-index], arr[index]
		} else {
			arr[index], arr[length-index] = arr[length-index], arr[index]
		}
	}

	// fmt.Printf("0x%s\n", arr)
	// fmt.Printf("%s\n", "0xa7274594ce215208c8e309e8f2fe05d4a9ae412b")

	return fmt.Sprintf("0x%s\n", string(arr))

}
