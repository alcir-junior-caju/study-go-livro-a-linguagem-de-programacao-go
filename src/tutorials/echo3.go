package tutorials

import (
	"fmt"
	"os"
	"strings"
)

func ShowEcho3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
