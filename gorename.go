package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/akamensky/argparse"
)

func readdirintobuffer(directory string) {
	// read contents of `directory` into `dir`
	dir, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}
	// all add files into a slice
	fileSlice := make([]string, 0)
	for _, file := range dir {
		fileSlice = append(fileSlice, file.Name())
	}
	// convert slice into string with new line separating each item
	filesAsString := strings.Join(fileSlice, "\n")
	// create buffer files and then write all items to it
	bufferFile, _ := os.Create("buffer.txt")
	bytes, err := bufferFile.WriteString(filesAsString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Wrote %d bytes to file\n", bytes)
}

func readLines(path string) ([]string, error) {
	// open file for reading
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	// check for empty file
	if len(lines) == 0 {
		return nil, errors.New("File was empty")
	}
	return lines, scanner.Err()
}

func rename(contentDir string) {
	bufferFile := "buffer.txt"
	// read all lines of buffer file into slice
	fileLines, err := readLines(bufferFile)
	if err != nil {
		log.Fatal(err)
	}
	// read all files in `contentDir` into slice for iteration
	dir, _ := ioutil.ReadDir(contentDir)
	contentFiles := make([]string, 0)
	for _, file := range dir {
		contentFiles = append(contentFiles, file.Name())
	}

	// use index to rename file from the current (in the directory) to the new (in the buffer file)
	for i := 0; i < len(fileLines); i++ {
		os.Rename(fmt.Sprintf("%s/%s", contentDir, contentFiles[i]), fmt.Sprintf("%s/%s", contentDir, fileLines[i]))
	}
	// remove buffer after renaming
	removeErr := os.Remove("buffer.txt")
	if removeErr != nil {
		log.Fatal(err)
	}

}

func main() {
	parser := argparse.NewParser("gorename", "Bulk renaming utility written in golang")
	command := parser.Selector("e", "command", []string{"rename", "prep"}, &argparse.Options{
		Required: true,
		Help:     "The command you want to run. `rename` uses a buffer file to rename files. `prep` creates a buffer file",
	})

	contentDir := parser.String("c", "contentDir", &argparse.Options{
		Required: false,
		Help:     "Directory of files to be renamed",
	})

	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
	} else {
		if *command == "rename" {
			rename(*contentDir)
		} else {
			readdirintobuffer(*contentDir)
		}

	}
}
