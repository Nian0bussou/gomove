package main

import (
	"fmt"
	"golang.org/x/term"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sync"
)

var (
	Count___ uint
	Succeed_ uint
	Failed__ uint
)

func main() {
	osName := runtime.GOOS
	path := "D:/grapper"
	if osName == "linux" {
		path = "/mnt/d/grapper/"
	}
	if len(os.Args) > 1 {
		cf := countFiles(path)
		println(cf)
	}
	if _, err := os.Stat(path); err != nil {
		println("dir not exist")
	}
	subs := get_folders(path)
	if subs == nil {
		println("could not get directories")
		return
	}
	for _, f := range subs {
		println(f)
	}
	width, _, _ := term.GetSize(0)
	for i := 0; i < width; i++ {
		print("_")
	}
	var wg sync.WaitGroup
	for _, item := range subs {
		wg.Add(1)
		go func(item string) {
			defer wg.Done()
			move_stuff(item)
		}(item)
	}
	wg.Wait()
	fmt.Println("finished")
	println("count : ", Count___)
	println("succe : ", Succeed_)
	println("faile : ", Failed__)
	if osName == "windows" {
		exec.Command("explorer.exe", "D:\\grapper\\").Run()
	}
}

func move_stuff(dir string) {
	entries, _ := os.ReadDir(dir)
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		dwall := filepath.Join(dir, "wall")
		dother := filepath.Join(dir, "other")
		dsquare := filepath.Join(dir, "square")
		dbadquality := filepath.Join(dir, "bad_quality")
		dbadqualitylandscape := filepath.Join(dbadquality, "l")
		dbadqualitysquare := filepath.Join(dbadquality, "s")
		dbadqualityportrait := filepath.Join(dbadquality, "p")
		dvideo := filepath.Join(dir, "video")
		dests := []string{
			dwall,
			dother,
			dsquare,
			dbadquality,
			dvideo,
			dbadqualitylandscape,
			dbadqualitysquare,
			dbadqualityportrait,
		}
		do(dests)
		checkThenMove(dir, filepath.Join(dir, e.Name()))
	}
}

func checkThenMove(dir, path string) {
	// change that so it doesnt do it every loop
	dwall := filepath.Join(dir, "wall")
	dother := filepath.Join(dir, "other")
	dsquare := filepath.Join(dir, "square")
	dbadquality := filepath.Join(dir, "bad_quality")
	dbadqualitylandscape := filepath.Join(dbadquality, "l")
	dbadqualitysquare := filepath.Join(dbadquality, "s")
	dbadqualityportrait := filepath.Join(dbadquality, "p")
	dvideo := filepath.Join(dir, "video")
	ext := filepath.Ext(path)
	var estr string = string(ext)
	println(estr)
	if ext == ".mp4" {
		moveFile(path, dvideo, "yellow")
	} else {
		file, err := os.Open(path)
		if err != nil {
			return
		}
		defer file.Close()
		m, _, e := image.Decode(file)
		if e != nil {
			return
		}
		b := m.Bounds()
		w := b.Dx()
		h := b.Dy()
		ar := float32(w) / float32(h)
		if w >= 1080 && h >= 1080 {
			if ar > 1 {
				moveFile(path, dwall, "red")
			} else if ar == 1 {
				moveFile(path, dsquare, "blue")
			} else if ar < 1 {
				moveFile(path, dother, "green")
			}
		} else {
			if ar > 1 {
				moveFile(path, dbadqualitylandscape, "cyan")
			} else if ar == 1 {
				moveFile(path, dbadqualitysquare, "magenta")
			} else if ar < 1 {
				moveFile(path, dbadqualityportrait, "purple")
			} else {
				moveFile(path, dbadquality, "grey")
			}
		}
	}
}

func moveFile(source, dest, name string) error {
	Count___++
	fileName := filepath.Base(source)
	destpath := filepath.Join(dest, fileName)
	err := os.Rename(source, destpath)
	if err != nil {
		Failed__++
		errorMaxxing(err.Error())
		return err
	} else {
		Succeed_++
		logmaxxing(source, dest, name)
		return nil
	}

}

func logmaxxing(source, file_path, name string) {
	reset := "\033[0m"
	var colors = map[string]string{
		"red":     "\033[31m",
		"blue":    "\033[34m",
		"green":   "\033[32m",
		"cyan":    "\033[36m",
		"magenta": "\033[35m",
		"purple":  "\033[35;1m",
		"grey":    "\033[37m",
		"yellow":  "\033[33m",
		"reset":   "\033[0m", // aka white
	}
	c := colors[name]
	tab := "\t"
	parentDir := filepath.Join(filepath.Base(filepath.Dir(source)), filepath.Base(source))
	padpar := fmt.Sprintf("%-80s", parentDir)
	fmt.Println(tab, c, tab, padpar, "=>", tab, file_path, reset)
}

func get_folders(path string) []string {
	files, err := os.ReadDir(path)
	if err != nil {
		println("error: ", err)
		return nil
	}
	var subdir []string
	for _, file := range files {
		if file.IsDir() {
			subdir = append(subdir, filepath.Join(path, file.Name()))
		}
	}
	return subdir
}

func do(destinationsFolders []string) {
	for _, s := range destinationsFolders {
		err := os.MkdirAll(s, 0755)
		if err != nil {
			println("error dir ", err)
			return
		}
	}
}

func countFiles(path string) uint {
	var nFiles uint
	var size float64
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			nFiles++
			size += getFileSize(path)
			if nFiles%100 == 0 {
				fmt.Printf("%d\t%.1f\n", nFiles, size)
			}
		}
		return nil
	})
	if err != nil {
		println("error counting")
		return 0
	}
	return nFiles
}

func getFileSize(filename string) float64 {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return 0
	}
	return float64(fileInfo.Size()) / (1024 * 1024 * 1024)
}

func errorMaxxing(str string) {
	fmt.Println("ERROR: ----------------------------------------------------------------------------------------------- ", str)
}
