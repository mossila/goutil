package goutil

import (
	"bufio"
	"log"
	"os"
)

//Readlns from file result in chan
/**
for s := range Readlns(filename) {
    fmt.Println(s)
}
*/
func Readlns(path string) chan string {
	outChan := make(chan string)
	go func() {
		inFile, err := os.Open(path)
		defer func() {
			inFile.Close()
			close(outChan)
		}()
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(inFile)
		for scanner.Scan() {
			outChan <- scanner.Text()
		}
	}()
	return outChan
}
