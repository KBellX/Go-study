/*
12.8
写一个程序读取 vcard.gob 文件，解码并打印它的内容。
*/
package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

var vc VCard

func main() {
	// 取文件句柄
	file, err := os.Open("vcard.gob")
	if err != nil {
		fmt.Printf("An error occurred on the inputfile: %s\n", err)
	}
	defer file.Close()

	// 用gob操作
	dec := gob.NewDecoder(file)
	err = dec.Decode(&vc)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(vc)

}
