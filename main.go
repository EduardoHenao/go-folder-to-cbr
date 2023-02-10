package main

import (
	"archive/zip"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Tuple struct {
	source, dest, entryName string
}

func separator() {
	separator := "------------------------------------------------------"
	fmt.Println(separator)
}

func isCandidate(entries []fs.DirEntry) bool {
	for _, entry := range entries {
		if entry.IsDir() {
			return false
		}
	}
	// fmt.Println("this is a CANDIDATE FOR CBR")
	return true
}

func print_scan_recursive(source string, dest string, entryName string) {
	fmt.Println("")
	separator()
	fmt.Printf("scan_recursive source[%v] dest[%v] entryName[%v]", source, dest, entryName)
	fmt.Println("")
	separator()
}

func scan_recursive(source string, dest string, entryName string, folders *[]Tuple) {
	print_scan_recursive(source, dest, entryName)

	// variables
	entries := []fs.DirEntry{}

	// Scan everything inside this source (dir)
	entries, err := os.ReadDir(source)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("len(entries) [%v]", len(entries))
	// fmt.Println("")

	// Detect if no subfolders
	if isCandidate(entries) {
		*folders = append(*folders, Tuple{source, dest, entryName})
	}

	// recursive part for directories
	for _, entry := range entries {
		entryName := entry.Name()
		if entry.IsDir() {
			scan_recursive(filepath.Join(source, entryName), filepath.Join(dest, entryName), entryName, folders)
		}
	}
}

func addFiles(w *zip.Writer, basePath string) {
	// in theory, all files here are FILES, there are no dirs.
	separator()
	println("addFiles [", basePath, "]")
	separator()

	// Open the Directory
	files, err := os.ReadDir(basePath)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		// local variables
		completeFilename := basePath + "\\" + file.Name()
		fmt.Println(completeFilename)

		if !file.IsDir() {
			dat, err := os.ReadFile(completeFilename)
			if err != nil {
				fmt.Println(err)
			}

			// Add some files to the archive.
			f, err := w.Create(file.Name())
			if err != nil {
				fmt.Println(err)
			}
			_, err = f.Write(dat)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func main() {
	fmt.Println(`program start`)

	//variables
	folders := []Tuple{}

	// source := ".\\source\\"
	// source := "G:\\ed\\material\\02 p processed to depurate\\"
	source := "G:\\ed\\material\\03 depurated to cbr\\"
	dest := ".\\dest\\"

	// main call
	scan_recursive(source, dest, "", &folders)

	// see folder scan results
	println("")
	separator()
	println("folders array")
	separator()
	for _, folder := range folders {
		println(folder.entryName) // print all obtained directories
	}
	separator()
	println()

	// compress and change ending for each folder
	for _, folder := range folders {
		// create dir
		os.MkdirAll(folder.dest, os.FileMode(0522))

		// create file
		processedFilepath := strings.Join([]string{folder.dest, "\\", folder.entryName, ".cbr"}, "")
		archive, err := os.Create(processedFilepath)
		if err != nil {
			panic(err)
		}

		// create zip writer
		zipWriter := zip.NewWriter(archive)

		// Add some files to the archive.
		addFiles(zipWriter, folder.source)

		// Make sure to check the error on Close.
		err = zipWriter.Close()
		if err != nil {
			fmt.Println(err)
		}

		err = archive.Close()
		if err != nil {
			fmt.Println(err)
		}
	}
}
