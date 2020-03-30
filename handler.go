package main

import (
	"MyCrypto/sm4"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"log"
	"path"
)

const (
	BlockSize = 16
)

func Encrypt(file, key, iv string) {
	res := make([]byte, 0)
	fileSuffix := path.Ext(file) //获取文件后缀

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	IV, err := hex.DecodeString(iv)
	if err != nil {
		log.Fatal("iv decode err :" + err.Error())
	}
	res = append(res, IV...) // 加密文件[0,16)表示iv

	KEY, err := hex.DecodeString(key)
	if err != nil {
		log.Fatal("key decode err :" + err.Error())
	}
	sum := sha256.Sum256(KEY)
	for i := 0; i < BlockSize; i++ { // 加密文件[16,32)表示hash之后的key
		res = append(res, sum[i])
	}

	data = sm4.Padding(data, BlockSize) // 尾部填充
	n := len(data)
	for i := 0; i < n; i += BlockSize { // CBC加密
		text := data[i : i+BlockSize]
		process(text, IV)
		IV = sm4.Encrypt(text, KEY)
		res = append(res, IV...)
	}
	err = ioutil.WriteFile("tmp"+fileSuffix, res, 0644) //创建加密后的文件tmp
	if err != nil {
		log.Fatal(err)
	}
}

func Decrypt(file, key string) {
	res := make([]byte, 0)
	fileSuffix := path.Ext(file) //获取文件后缀

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	n := len(data)

	IV := data[:BlockSize]
	KEY, err := hex.DecodeString(key)
	if err != nil {
		log.Fatal("key decode err :" + err.Error())
	}

	sum := sha256.Sum256(KEY)
	for i := BlockSize; i < BlockSize*2; i++ { // 判断加解密的key是否相同
		if sum[i-BlockSize] != data[i] {
			log.Fatal("you use wrong key")
		}
	}

	for i := 2 * BlockSize; i < n; i += BlockSize { // CBC解密
		sText := data[i : i+BlockSize]
		text := sm4.Decrpty(sText, KEY)
		process(text, IV)
		IV = sText
		res = append(res, text...)
	}

	err = ioutil.WriteFile("result"+fileSuffix, sm4.UnPadding(res), 0644) //创建解密后的文件result
	if err != nil {
		log.Fatal(err)
	}
}

func process(text, iv []byte) { // IV异或
	for i := 0; i < BlockSize; i++ {
		text[i] ^= iv[i]
	}
}
