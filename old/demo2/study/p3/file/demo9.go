/*
input:
"The ABC of Go";25.5;1500
"Functional Programming with Go";56;280
"Go for It";45.9;356
"The Go Way";55;500

output
{"The ABC of Go" 25.5 150}
{"Functional Programming with Go" 56 28}
{"Go for It" 45.900001525878906 35}
{"The Go Way" 55 50}

*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	title    string
	price    float64
	quantity int
}

func main() {
	file := "products.txt"
	// 声明切片，然后一直append
	// bks := []Book{}
	var bks []Book

	inputFile, err := os.Open(file)
	if err != nil {
		fmt.Printf("An error occurred:%s\n", err)
	}
	defer inputFile.Close()

	reader := bufio.NewReader(inputFile)

	for {
		inputString, inputError := reader.ReadString('\n')
		if inputError == io.EOF {
			break
		}

		inputArr := strings.Split(inputString[:len(inputString)-1], ";")

		book := &Book{}
		book.title = inputArr[0]
		book.price, _ = strconv.ParseFloat(inputArr[1], 32)
		book.quantity, _ = strconv.Atoi(inputArr[2])

		bks = append(bks, *book)
	}

	for _, book := range bks {
		fmt.Println(book)
	}
}
