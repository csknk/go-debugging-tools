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
	fmt.Println(input)

	data := strings.Split(input, ",")
	var numSlice []uint8
	for _, s := range data {
		num, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal("NaN: ", err)
		}
		numSlice = append(numSlice, uint8(num))
	}
	printHex(&numSlice)
}

func printHex(data *[]uint8) {
	fmt.Printf("0x")
	for _, x := range *data {
		fmt.Printf("%02x", x)
	}
	fmt.Printf("\n")
}

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
