package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
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
		if !info.IsDir() {
			files = append(files, path)
		} else {
			fmt.Println(".. scanning ", info.Name())
		}
		return nil
	})
	return files, err
}
