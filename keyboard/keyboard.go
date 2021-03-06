package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const olovo int = 45
const Witcher int = 34

func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil

}

func PrintConstant() {
	fmt.Println("Konstanata iz paketa", olovo)
	fmt.Println("bla bla")
}
