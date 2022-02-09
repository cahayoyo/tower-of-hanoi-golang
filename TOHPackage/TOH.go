package TOHPackage

import "fmt"

func TOH(n int, A int, B int, C int) {
	if n > 0 {
		TOH(n-1, A, C, B)
		fmt.Println("(", A, ",", C, ")")
		TOH(n-1, B, A, C)
	}
}
