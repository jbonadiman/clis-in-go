package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")
	fileNames := flag.String("file", "", "provide a file name or multiple file names separated by comma")
	flag.Parse()

	target := []io.Reader{os.Stdin}

	if *fileNames != "" {
		var err error

		target, err = parseFiles(*fileNames)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	fmt.Println(count(*lines, *bytes, target...))
}

func parseFiles(fileNames string) ([]io.Reader, error) {
	parsedFileNames := strings.Split(fileNames, ",")

	files := []io.Reader{}

	for _, name := range parsedFileNames {
		file, err := os.Open(name)
		if err != nil {
			return nil, err
		}

		files = append(files, file)
	}

	return files, nil
}

func count(countLines bool, countBytes bool, r ...io.Reader) int {
	wc := 0

	for _, reader := range r {
		scanner := bufio.NewScanner(reader)

		if countBytes {
			scanner.Split(bufio.ScanBytes)
		} else if !countLines {
			scanner.Split(bufio.ScanWords)
		}

		for scanner.Scan() {
			wc++
		}
	}

	return wc
}
