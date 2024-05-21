package tutorials

import (
	"fmt"
	"os"
)

func ShowEcho1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("Echo1 Original:", s)

	var s1, sep1 string
	for i := 0; i < len(os.Args); i++ {
		s1 += sep1 + os.Args[i]
		sep = " "
	}
	fmt.Println("Echo1 Command Name:", s)

	for i, arg := range os.Args {
		fmt.Printf("Echo1 Index: %d, Value: %s\n", i, arg)
	}
}
