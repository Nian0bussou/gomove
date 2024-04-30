package main

import (
	"fmt"
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
	fmt.Println(path)

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

	println("_______________")

	for _, f := range subs {
		println(f)
	}
	println("_______________")

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
		checkThenMove(dir, filepath.Join(dir, e.Name()))
	}

}

func checkThenMove(dir, path string) {

	dest_wall := filepath.Join(dir, "wall")
	dest_othe := filepath.Join(dir, "other")
	dest_squa := filepath.Join(dir, "square")
	dest_badq := filepath.Join(dir, "bad_quality")
	dest_bad_land := filepath.Join(dest_badq, "l")
	dest_bad_squa := filepath.Join(dest_badq, "s")
	dest_bad_port := filepath.Join(dest_badq, "p")
	dest_vide := filepath.Join(dir, "video")
	do(dest_wall)
	do(dest_othe)
	do(dest_squa)
	do(dest_badq)
	do(dest_vide)
	do(dest_bad_land)
	do(dest_bad_squa)
	do(dest_bad_port)

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

	ext := filepath.Ext(path)

	if ext == ".mp4" {
		moveFile(path, dest_vide, "yellow")
	} else {
		if w >= 1080 && h >= 1080 {
			if ar > 1 {
				// wallp
				moveFile(path, dest_wall, "red")
			} else if ar == 1 {
				// square
				moveFile(path, dest_squa, "blue")
			} else if ar < 1 {
				// other
				moveFile(path, dest_othe, "green")
			}
		} else {
			// badquality
			if ar > 1 {
				moveFile(path, dest_bad_land, "cyan")
			} else if ar == 1 {
				moveFile(path, dest_bad_squa, "magenta")
			} else if ar < 1 {
				moveFile(path, dest_bad_port, "purple")
			} else {
				moveFile(path, dest_badq, "grey")
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
	}
	Succeed_++
	logmaxxing(source, dest, name)
	return nil

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

func do(s string) {
	err := os.MkdirAll(s, 0755)
	if err != nil {
		println("error making dir: ", err)
		return
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
