package jarome

import (
	"bufio"
	"fmt"
	"os"
)

func array_contains(a []string, t string) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == t {
			return true
		}
	}
	return false
}

func Init() {
	fmt.Println("=== STARTING JAROME ===")
	var exit_tokens = []string{"exit", "quit", "close", "q"}
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Printf("> ")
		scanner.Scan()
		input := scanner.Text()
		if array_contains(exit_tokens, input) {
			break
		}

	}
}
