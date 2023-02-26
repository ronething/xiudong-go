package showstart

import (
	"crypto/md5"
	"fmt"
)

func Md5Sum(value string) string {
	data := []byte(value)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func Md5SumByte(value []byte) string {
	return fmt.Sprintf("%x", md5.Sum(value))
}
