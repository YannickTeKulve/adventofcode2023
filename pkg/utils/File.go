package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadFile(filename string, line func(value string)) {
	file, errOpen := os.Open(filename)
	if errOpen != nil {
		log.Fatal(errOpen.Error())
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
