package main

import (
	"fmt"
	"golang.org/x/term"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

const (
	land_s string = "━"
	squa_s string = "■"
	port_s string = "┃"
	vide_s string = "▶"
)

var (
	choice  uint
	Count   uint
	Succeed uint
	Failed  uint
	res     string = "\033[0m"

	colors = map[string]string{
		"red":     "\033[31m",
		"blue":    "\033[34m",
		"green":   "\033[32m",
		"cyan":    "\033[36m",
		"magenta": "\033[35m",
		"purple":  "\033[35;1m",
		"grey":    "\033[37m",
		"yellow":  "\033[33m",
	}
	specia = map[string]string{
		"land":     land_s,
		"square":   squa_s,
		"portrait": port_s,
		"video":    vide_s,
	}
)

func main() {
	defer trackTime()()

	//////////////////////////////////////////////////////////////////////////////

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "1":
			choice = 1
		default:
			log.Fatal("invalid option")
		}
	} else {
		choice = 0
	}

	/////////////////////////////////////////////////////////////////////////////
	osName := runtime.GOOS
	path := "D:/grapper"

	if osName == "linux" {
		path = "/mnt/d/grapper/"
	}
	if _, err := os.Stat(path); err != nil {
		println("dir not exist")
	}

	/////////////////////////////////////////////////////////////////////////////

	subs := get_folders(path)
	if subs == nil {
		println("could not get directories")
		return
	}
	for _, f := range subs {
		println(f)
	}

	//////////////////////////////////////////////////////////////////////////////

	width, _, _ := term.GetSize(0)
	for i := 0; i < width; i++ {
		print("_")
	}
	println()

	/////////////////////////////////////////////////////////////////////////////

	var wg sync.WaitGroup
	for _, item := range subs {
		wg.Add(1)
		go func(item string) {
			defer wg.Done()
			switch choice {
			case 0:
				move_stuff(item)
			case 1:
				scramble(item)
			default:
				log.Fatal("invalid option in choice... somehow")
			}
		}(item)
	}
	wg.Wait()

	//////////////////////////////////////////////////////////////////////////////
	for i := 0; i < width; i++ {
		print("_")
	}
	println("finished")
	println("count : ", Count)
	println("succe : ", Succeed)
	println("faile : ", Failed)
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
	var _ string = string(ext)

	if ext == ".mp4" {
		moveFile(path, dvideo, "yellow", "video")
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
				moveFile(path, dwall, "red", "land")
			} else if ar == 1 {
				moveFile(path, dsquare, "blue", "square")
			} else if ar < 1 {
				moveFile(path, dother, "green", "portrait")
			}
		} else {
			if ar > 1 {
				moveFile(path, dbadqualitylandscape, "cyan", "land")
			} else if ar == 1 {
				moveFile(path, dbadqualitysquare, "magenta", "square")
			} else if ar < 1 {
				moveFile(path, dbadqualityportrait, "purple", "portrait")
			} else {
				moveFile(path, dbadquality, "grey", "")
			}
		}
	}
}

func moveFile(source, dest, name, typoe string) error {
	Count++
	fileName := filepath.Base(source)
	destpath := filepath.Join(dest, fileName)

	err := os.Rename(source, destpath)
	if err != nil {
		Failed++
		errormaxxing(err.Error())
		return err
	} else {
		Succeed++
		logmaxxing(source, dest, name, typoe)
		return nil
	}

}

func logmaxxing(source, file_path, name, typoe string) {
	c := colors[name]
	special := specia[typoe]

	tab := "\t"
	parentDir := filepath.Join(filepath.Base(filepath.Dir(source)), filepath.Base(source))
	padpar := fmt.Sprintf("%-80s", parentDir)

	println(tab, c, special, tab, padpar, "<><><><><>", tab, file_path, res)
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
			log.Fatal("error making directories")
			//println("error dir ", err)
			//return
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

func errormaxxing(str string) {
	fmt.Println("ERROR: ----------------------------------------------------------------------------------------------- ", str)
}

func scramble(path string) {
	err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			Count++
			fileName := filepath.Base(p)
			err := os.Rename(p, filepath.Join(path, fileName))
			if err != nil {
				Failed++
				log.Fatal("cant move file")
			} else {
				Succeed++
				println("moved: ", fileName, "\t", p)
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal("error walking over eggs")
	}
}

func trackTime() func() {
	start := time.Now()
	return func() {
		log.Printf("Elapsed : %s", time.Since(start))
	}
}
