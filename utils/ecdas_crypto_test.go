package utils

import (
	"bytes"
	"strconv"
	"testing"
)

func TestHash160(t *testing.T) {
	// ad := Hash160([]byte("AKibPRzkoZpHnPkF6qvuW2Q4hG9gKBwGpR")) // 0xa7274594ce215208c8e309e8f2fe05d4a9ae412b

	// t.Log("ad", string(ad))
	// fmt.Println("11", ad)

	e := &ECDsaCrypto{}
	ad := e.Hash160([]byte("AKibPRzkoZpHnPkF6qvuW2Q4hG9gKBwGpR"))

	// dst := make([]byte, hex.EncodedLen(len(ad)))
	// hex.Encode(dst, ad)
	// fmt.Printf("%s\n", dst)
	//t.Log("ad", string(dst))

	// var address string = "AKibPRzkoZpHnPkF6qvuW2Q4hG9gKBwGpR"
	// ripemd160, _, err := base58.CheckDecode(address)
	// if err != nil {
	// 	fmt.Printf("Failed to decode address %s: %s\n", address, err)

	// }

	// fmt.Printf("HASH160 %x from address %s\n", ripemd160, address)

	// 遍历, 转为16进制
	buffer := new(bytes.Buffer)
	for _, b := range ad {

		s := strconv.FormatInt(int64(b&0xff), 16)
		if len(s) == 1 {
			buffer.WriteString("0")
		}
		buffer.WriteString(s)
	}
	// 转化为字符串
	// fmt.Println(buffer.String())
}
