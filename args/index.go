package args

import (
	"os"
	"fmt"
)

var App string

func GetAppName() {
	if len(os.Args) < 2 {
		fmt.Println("Try: gore <binary-name>")
		fmt.Println("Binary must be located at ~/release/<binary-name>")
	}
	App = os.Args[1]
}
