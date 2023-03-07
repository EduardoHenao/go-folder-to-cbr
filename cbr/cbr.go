package cbr

import (
	"archive/zip"
	"cbrsuite/common"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func isCandidate(entries []fs.DirEntry) bool {
	for _, entry := range entries {
		if entry.IsDir() {
			return false
		}
	}
	return true
}

func createCBRs(folders *[]common.Tuple) {
	for _, folder := range *folders {
		// create dir
		os.MkdirAll(folder.Dest, os.FileMode(0522))

		// create cbr file (it is a zip file, with a changed file extension to cbr)
		processedFilepath := strings.Join([]string{folder.Dest, "\\", folder.EntryName, ".cbr"}, "")
		archive, err := os.Create(processedFilepath)
		if err != nil {
			panic(err)
		}

		// create zip writer
		zipWriter := zip.NewWriter(archive)

		// Add some files to the archive.
		addFiles(zipWriter, folder.Source)

		// Make sure to check the error on Close.
		err = zipWriter.Close()
		if err != nil {
			fmt.Println(err)
		}

		// process errors
		err = archive.Close()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func addFiles(w *zip.Writer, basePath string) {
	// in theory, all files here are FILES, there are no dirs.
	common.Separator()
	println("addFiles [", basePath, "]")
	common.Separator()

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

func scan_recursive(source string, dest string, entryName string, folders *[]common.Tuple) {
	common.Print_scan_recursive(source, dest, entryName)

	// variables
	entries := []fs.DirEntry{}

	// Scan everything inside this source (dir)
	entries, err := os.ReadDir(source)
	if err != nil {
		log.Fatal(err)
	}

	// Detect if no subfolders
	if isCandidate(entries) {
		*folders = append(*folders, common.Tuple{Source: source, Dest: dest, EntryName: entryName})
	}

	// recursive part for directories
	for _, entry := range entries {
		entryName := entry.Name()
		if entry.IsDir() {
			scan_recursive(filepath.Join(source, entryName), filepath.Join(dest, entryName), entryName, folders)
		}
	}
}

func Cbr(source string, dest string) {
	//variables
	folders := []common.Tuple{}

	// main call
	scan_recursive(source, dest, "", &folders)

	// see folder scan results
	common.Print_folder_list(folders, "folders array")

	// compress and change ending for each folder
	createCBRs(&folders)
}
