package sm4

import (
	"bytes"
)

func Encrypt(text, key []byte) []byte {
	x := initX(text)      // 初始化x
	rk := expandKey(key)  // 生成rk
	tmp := encrypt(x, rk) // 加密明文
	res := make([]byte, 0)
	for _, v := range tmp {
		res = append(res, Uint32ToBytes(v)...)
	}
	return res
}

func Decrpty(text, key []byte) []byte {
	x := initX(text)      // 初始化x
	rk := expandKey(key)  // 生成rk
	tmp := decrypt(x, rk) // 解密密文
	res := make([]byte, 0)
	for _, v := range tmp {
		res = append(res, Uint32ToBytes(v)...)
	}
	return res
}

func Uint32ToBytes(n uint32) []byte { // 将一个字转为四个字节
	return []byte{

		byte(n >> 24),
		byte(n >> 16),
		byte(n >> 8),
		byte(n),
	}
}

func Padding(text []byte, blockSize int) []byte { // 尾部填充
	length := blockSize - len(text)%blockSize
	padText := bytes.Repeat([]byte{byte(length)}, length)
	return append(text, padText...)
}

func UnPadding(text []byte) []byte { // 删除填充
	length := len(text)
	padLength := int(text[length-1])
	return text[:(length - padLength)]
}
