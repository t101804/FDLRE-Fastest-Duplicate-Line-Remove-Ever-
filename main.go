// @CalMeRep Copyright 2022
// @IndonesianDarknet , @RepProject , fb : callmerep.real

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/cheggaaa/pb/v3"
)

func main() {
	file, err := os.Open("domains-amazon.txtrep") // change domains-amazon.textrep to what ever file and extension that you want to delete duplicate line 
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	lines := make(map[string]bool)
	scanner := bufio.NewScanner(file)

	// Ngitung total line
	var totalLines int
	for scanner.Scan() {
		totalLines++
	}

	// Reset point
	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)

	// Register Loading bar buat progress
	bar := pb.StartNew(totalLines)

	// Penunggu finish biar ga deadlock
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			line = strings.ToLower(line)
			if _, ok := lines[line]; !ok {
				lines[line] = true
			}
			bar.Increment()
		}
		wg.Done()
	}()

	wg.Wait()
	bar.Finish()
	newFile, err := os.Create("name.txt") // results file name , you can edit 
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer newFile.Close()

	for line := range lines {
		fmt.Fprintln(newFile, line)
	}
}
