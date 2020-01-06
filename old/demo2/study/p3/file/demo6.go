// 按行读取文件内容和写文件

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, err := os.Open("input_file.txt")
	if err != nil {
		fmt.Printf("An error occurred on opening the inputFile:%s \n", err)
		return
	}
	defer inputFile.Close()

	outputFile, err := os.OpenFile("output_file.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("An error occurred on opening the outputFile:%s \n", err)
	}
	defer outputFile.Close()

	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(outputFile)

	for {
		inputString, inputError := reader.ReadString('\n')
		if inputError == io.EOF {
			break
		}

		writer.WriteString(inputString)

		fmt.Println(inputString)
	}
	writer.Flush()

}
