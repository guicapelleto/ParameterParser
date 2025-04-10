package custom

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetInput(msg string) (retorno string) {
	buf := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	ret, _ := buf.ReadString('\n')
	retorno = strings.Trim(ret, "\r\n")
	return
}
