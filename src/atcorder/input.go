package atcorder

import (
	"bufio"
	"os"
)


func GetScanner() *bufio.Scanner {
	sc := bufio.NewScanner(os.Stdin)

	return sc
}

func GetNextLine( scanner *bufio.Scanner ) string {
	scanner.Scan()
	return scanner.Text()
}