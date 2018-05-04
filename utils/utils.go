package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func EncodeMD5(str string) string {
	fmt.Println("md5", str)
	h := md5.New()
	h.Write([]byte(str)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	fmt.Println(cipherStr)

	fmt.Printf("%s\n", strings.ToUpper(hex.EncodeToString(cipherStr))) // 输出加密结果
	return strings.ToUpper(hex.EncodeToString(cipherStr))
}
