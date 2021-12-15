package input

import (
	"bufio"
	"os"
)

func WaitForInputText() string {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	return input.Text()
}
