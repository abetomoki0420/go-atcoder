package atcorder

import (
	"fmt"
	"strconv"
	"strings"
)

// ABC086 A. Product
func Product() {
	sc := GetScanner()

	input := GetNextLine(sc)

	arr := strings.Split(input, " ")

	a, _ := strconv.Atoi(arr[0])
	b, _ := strconv.Atoi(arr[1])

	ans := productLogic(a,b)

	fmt.Println(ans)
}

func productLogic( a , b int) string {
	if ( a * b ) % 2 == 0 {
		// is even
		return "Even"
	}  else {
		// is odd
		return "Odd"
	}
}

