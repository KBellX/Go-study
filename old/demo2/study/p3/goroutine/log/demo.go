package main

import (
	"fmt"
	"regexp"
)

func main() {
	// content := "192.168.211.1 - - [27/Feb/2019:14:26:07 +0800] http \"GET /index.php HTTP/1.1\" 200 2024 \"-\" \"cba\" 1.001 2.02 "
	// rule := `([\d\.]+)\s+([^ \[]+)\s+([^ \[]+)\s+\[([^\]]+)\]\s+([a-z]+)\s+\"([^"]+)\"\s+(\d{3})\s+(\d+)\s+\"([^"]+)\"\s+\"([^"]+)\"\s+([\d\.]+)\s+([\d\.]+)`

	content := "192.168.211.1 - - [27/Feb/2019:14:26:07 +0800] http \"GET /index.php HTTP/1.1\" 200 2024 \"-\" \"cba\" 1.001 2.02 "
	rule := `([\d\.]+)\s+([^ \[]+)\s+([^ \[]+)\s+\[([^\]]+)\]\s+([a-z]+)\s+\"([^"]+)\"\s+(\d{3})\s+(\d+)\s+\"([^"]+)\"\s+\"([^"]+)\"\s+([\d\.]+)\s+([\d\.]+)`

	rep := regexp.MustCompile(rule)

	result := rep.FindStringSubmatch(content)

	// fmt.Println(result)

	// fmt.Println(result[len(result)-1])

	for i, v := range result {
		fmt.Printf("%d : %s\n", i, v)
	}
}
