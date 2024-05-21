package tutorials

import (
	"fmt"
	"os"
)

func ShowEcho2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
