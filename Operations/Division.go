package Operations

import "fmt"

func Parity() {

	const a int = 5
	const b int = 20

	const division = float32(a) / float32(b)

	var fact string = fmt.Sprintf("%d divided by %d is %.2f", a, b, division)

	fmt.Println(fact)

	fmt.Println(a << b)

	if a%2 == 0 {
		fmt.Println(a, "is Even number")
	} else {
		fmt.Println(a, "is Odd number")
	}
}
