package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func logmaxxing(source, file_path, name, typoe string) {
	c := colors[name]
	special := special_char[typoe]

	if count, ok := countTypes[typoe]; ok {
		*count++
	}

	tab := "\t"
	parentDir := filepath.Join(filepath.Base(filepath.Dir(source)), filepath.Base(source))
	padpar := fmt.Sprintf("%-80s", parentDir)

	println(tab, c, special, tab, padpar, "<| |=| |>", tab, file_path, reset_white)
}

func errormaxxing(str string) {
	fmt.Println("ERROR: ----------------------------------------------------------------------------------------------- ", str)
}

func removeTmp(path string) {
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
