package main

import (
	"fmt"
	"regexp"
)

func main() {
	content := "[27/Feb/2019:14:26:07 +0800]"

	rule := `\[([^\]]+)\]`

	rep := regexp.MustCompile(rule)

	result := rep.FindStringSubmatch(content)

	for i, v := range result {
		fmt.Printf("%d : %s\n", i, v)
	}
}
