package dfs

import (
	"bufio"
	"os"
)

func readFile(filename string) *[]string {
	readFile, err := os.Open(filename)

	if err != nil {
		return nil
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	return &fileLines
}
