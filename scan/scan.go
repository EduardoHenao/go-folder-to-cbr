package scan

import (
	"cbrsuite/common"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func isEmpty(source string, entries []fs.DirEntry, extensions []string) bool {
	// if no files, then is an empty folder
	if len(entries) == 0 {
		return true
	}

	// if it has some suspicious types inside, so no empty
	for _, entry := range entries {
		if !entry.IsDir() {
			temp_file_path := filepath.Join(source, entry.Name())
			extension := filepath.Ext(temp_file_path)
			if common.Contains(extensions, extension) {
				return true
			}
		}
	}

	return false
}

func scan_recursive_empty(source string, entryName string, folders *[]common.Tuple) {
	// variables
	entries := []fs.DirEntry{}

	// Scan everything inside this source (dir)
	entries, err := os.ReadDir(source)
	if err != nil {
		log.Fatal(err)
	}

	// Detect if no subfolders
	extensions := []string{}
	if isEmpty(source, entries, extensions) == true {
		*folders = append(*folders, common.Tuple{Source: source, Dest: "", EntryName: entryName})
	}

	// recursive part for directories
	for _, entry := range entries {
		entryName := entry.Name()
		if entry.IsDir() {
			scan_recursive_empty(filepath.Join(source, entryName), entryName, folders)
		}
	}
}

func Scan(source string) {
	//variables
	folders := []common.Tuple{}

	// main call
	scan_recursive_empty(source, "", &folders)

	// see folder scan results
	common.Print_folder_list(folders, "empty or suspicious folders array")
}
