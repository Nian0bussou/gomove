package main

import (
	"fmt"
	"golang.org/x/term"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func get_path() string {
	path := "D:/grapper"
	if runtime.GOOS == "linux" {
		path = "/mnt/d/grapper/"
	}
	if _, err := os.Stat(path); err != nil {
		log.Fatal("dir not exist")
	}
	return path
}

func line_print() {
	println()
	width, _, _ := term.GetSize(0)
	for i := 0; i < width; i++ {
		print("_")
	}
	println()

}

func get_choice(path string) (int, string) {
	Choice := 0

	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "1":
			Choice = 1
		default:
			log.Fatal("invalid option")
		}
	} else if len(os.Args) == 3 {
		switch os.Args[1] {
		case "0":
			Choice = 0
		case "1":
			Choice = 1
		default:
			log.Fatal("invalid option")
		}
		path = os.Args[2]
	}
	return Choice, path
}

func logmaxxing(source, file_path, name, typoe string) {
	c := colors[name]
	special := special_char[typoe]

	if count, ok := countTypes[typoe]; ok {
		*count++
	}

	tab := "\t"
	parentDir := filepath.Join(filepath.Base(filepath.Dir(source)), filepath.Base(source))
	padpar := fmt.Sprintf("%-80s", parentDir)

	println(tab, c, special, tab, padpar, "|====|", tab, file_path, reset_white)
}

func errormaxxing(str string) {
	fmt.Println("ERROR: ----------------------------------------------------------------------------------------------- ", str)
}

func removeTmp(path string) {
	println("Removing abandonned tmp files")
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".tmp" {
			os.Remove(path)
		}
		return nil
	})
	if err != nil {
		log.Fatal("error walking over eggs")
	}
}

func timeminmaxxing() func() {
	start := time.Now()
	return func() {
		log.Printf("Elapsed : %s", time.Since(start))
	}
}
