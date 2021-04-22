package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := openStdinOrFile()
	data := strings.Split(input, ",")
	var result []rune
	for _, s := range data {
		num, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal("NaN: ", err)
		}
		result = append(result, rune(num))
	}
	fmt.Println(string(result))
}

//func printHex(data *[]uint8) {
//	fmt.Printf("0x")
//	for _, x := range *data {
//		fmt.Printf("%02x", x)
//	}
//	fmt.Printf("\n")
//}

func openStdinOrFile() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter an Array of bytes in decimal format: ")
	scanner.Scan()
	input := scanner.Text()
	return input
}
