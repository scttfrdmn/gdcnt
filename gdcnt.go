package main

import (
    "fmt"
    "log"
    "os"
)

func count(dirname string) (int, int) {
    dir := 0
    fil := 0

    f, err := os.Open(dirname)
    if err != nil {
        log.Fatal(err)
    }
    files, err := f.Readdir(-1)
    f.Close()
    if err != nil {
        log.Fatal(err)
    }
    
    for _, file := range files {
	if file.IsDir() {
//		fmt.Println("dirname:", dirname )
//		fmt.Println("Directory Name:", file.Name() )
		dir++
                dd, ff := count( dirname + "/" + file.Name() )
		dir += dd
		fil += ff
	} else {
		fil++
	}
    }
    fmt.Println(dirname)
    fmt.Println("  total:", dir+fil, "files:", fil, "dirs:", dir)
    return dir, fil
}

func main() {
    args := os.Args[1:]

    dirname := args[0]
    _, _ = count( dirname )
}
