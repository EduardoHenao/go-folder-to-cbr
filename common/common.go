package common

import (
	"fmt"
	"strconv"
)

////////////////////////////////////////////////////////////////////////////////////////////////
// struct defs
////////////////////////////////////////////////////////////////////////////////////////////////

type Tuple struct {
	Source, Dest, EntryName string
}

////////////////////////////////////////////////////////////////////////////////////////////////
// print methods
////////////////////////////////////////////////////////////////////////////////////////////////

func Separator() {
	separator := "------------------------------------------------------"
	fmt.Println(separator)
}

func Print_folder_list(folders []Tuple, legend string) {
	println("")
	Separator()
	println(legend + "[" + strconv.Itoa(len(folders)) + "]")
	Separator()
	for _, folder := range folders {
		println("[" + folder.Source + "]") // print all obtained directories
	}
	Separator()
	println()
}

func Print_scan_recursive(source string, dest string, entryName string) {
	fmt.Println("")
	Separator()
	fmt.Printf("scan_recursive source[%v] dest[%v] entryName[%v]", source, dest, entryName)
	fmt.Println("")
	Separator()
}

////////////////////////////////////////////////////////////////////////////////////////////////
// array helpers
////////////////////////////////////////////////////////////////////////////////////////////////

func Contains(array_of_strings []string, str string) bool {
	for _, v := range array_of_strings {
		if v == str {
			return true
		}
	}
	return false
}
