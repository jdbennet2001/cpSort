package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// this is a comment

func main() {

	// Get command line arguments
	numbPtr := flag.Int("size", 128, "Target bucket size")
	srcPtr := flag.String("src", "/Volumes/Seagate Expansion Drive/comics", "Source directory")
	destPtr := flag.String("dest", "/Users/igg/projects/tmp", "Destination directory")
	flag.Parse()

	fmt.Println("numb:", *numbPtr)
	fmt.Println("src:", *srcPtr)
	fmt.Println("dest:", *destPtr)

	fmt.Println("Hello World")

	files, err := FilePathWalkDir(*srcPtr)
	if err != nil {
		fmt.Println("Error")
	} else {
		fmt.Println(len(files), " files")
	}

}

func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		// Check if the parent directory has a '.' prefix
		parent := filepath.Dir(path)
		base := filepath.Base(parent)

		if info.IsDir() {
			fmt.Println(".. scanning ", info.Name())
		} else if strings.HasPrefix(base, ".") || strings.HasPrefix(base, " ") {
			fmt.Println(".. skipping ", path)
		} else if strings.HasSuffix(path, "cbz") {
			files = append(files, path)
		}

		return nil
	})
	return files, err
}
