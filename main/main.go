package main

import (
	"cbrsuite/cbr"
	"cbrsuite/scan"
	"fmt"
	"os"
)

//TODO:
// pass paths by console (to isolate specific paths from code)
// upload to github!
// documentation + polish

func main() {
	fmt.Println(`program start`)

	// variables for scan (execute "go run .\main\main.go -s")
	scan_source := "C:\\Users\\ed\\Desktop\\new col\\"

	// variables for cbr (execute "go run .\main\main.go")
	cbr_source := "C:\\Users\\ed\\Desktop\\bbb\\"
	cbr_dest := "C:\\Users\\ed\\Desktop\\bbbCBR\\"

	if len(os.Args) > 1 {
		firstArg := os.Args[1]
		if firstArg == "-s" {
			scan.Scan(scan_source)
		}
	} else {
		cbr.Cbr(cbr_source, cbr_dest)
	}
}
