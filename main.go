package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	db := make(map[string]string)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("📝 Simple In-Memory DB (type EXIT to quit)")

	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err)
			continue
		}
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		parts := strings.Fields(line)
		cmd := strings.ToUpper(parts[0])

		switch cmd {
		case "EXIT":
			fmt.Println("Bye 👋")
			return

		case "SET":
			if len(parts) != 3 {
				fmt.Println("Usage: SET key value")
				continue
			}
			key := parts[1]
			value := parts[2]
			db[key] = value
			fmt.Println("OK")

		case "GET":
			if len(parts) != 2 {
				fmt.Println("Usage: GET key")
				continue
			}
			key := parts[1]
			value, ok := db[key]
			if ok {
				fmt.Println(value)
			} else {
				fmt.Println("(nil)")
			}

		case "DEL":
			if len(parts) != 2 {
				fmt.Println("Usage: DEL key")
				continue
			}
			key := parts[1]
			delete(db, key)
			fmt.Println("OK")

		case "KEYS":
			if len(db) == 0 {
				fmt.Println("(empty)")
				continue
			}
			for k := range db {
				fmt.Println(k)
			}

		default:
			fmt.Println("Unsupported command")
		}
	}
}
