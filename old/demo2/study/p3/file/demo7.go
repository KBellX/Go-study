// 整个文件读取&写入

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputFile := "input_file.txt"
	outputFile := "output_file.txt"

	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File error: %s\n", err)
	}

	fmt.Printf("%s\n", string(buf))

	// 八进制777权限
	err = ioutil.WriteFile(outputFile, buf, 0777)

	if err != nil {
		panic(err.Error())
	}

}
