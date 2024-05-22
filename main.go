package main

import (
	"fmt"
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

	"golang.org/x/term"
)

const (
	landscape_char string = "━"
	square_char    string = "■"
	portrait_char  string = "┃"
	video_char     string = "▶"
)

var (
	Choice          uint
	Count           uint
	Succeed         uint
	Failed          uint
	Landscape_count uint
	Portrait_count  uint
	Square___count  uint
	Video____count  uint
	reset_white     string = "\033[0m"

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
	special_char = map[string]string{
		"land":     landscape_char,
		"square":   square_char,
		"portrait": portrait_char,
		"video":    video_char,
	}

	countTypes = map[string]*uint{
		landscape_char: &Landscape_count,
		square_char:    &Square___count,
		portrait_char:  &Portrait_count,
		video_char:     &Video____count,
	}
)

func main() {
	defer timeminmaxxing()()

	/////////////////////////////////////////////////////////////////////////////
	// tak youw path & get the wuck out

	osName := runtime.GOOS
	path := "D:/grapper"

	if osName == "linux" {
		path = "/mnt/d/grapper/" // change this because fuck me
		// log.Fatal("This dir : \"", path, "\" is invalid")
	}
	if _, err := os.Stat(path); err != nil {
		println("dir not exist")
	}

	//////////////////////////////////////////////////////////////////////////////
	// determining what wheter to scramble or to mangle

	Choice = 0

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

	/////////////////////////////////////////////////////////////////////////////
	// get the fowdews

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
			switch Choice {
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
	println("Finished,")
	println("count     : ", Count)
	println("succeeded : ", Succeed)
	println("failed    : ", Failed)
	println("portrait  : ", Portrait_count)
	println("landscape : ", Landscape_count)
	println("square    : ", Square___count)
	println("video     : ", Video____count)

	removeTmp(path)

	///////////////////////////////////////////////////////////////////////////////
	if osName == "windows" {
		exec.Command("explorer.exe", "D:\\grapper\\").Run()
	}
	//////////////////////////////////////////////////////////////////////////////
}

func move_stuff(dir string) {
	entries, _ := os.ReadDir(dir)
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		dwall_______________ := filepath.Join(dir, "wall")
		dother______________ := filepath.Join(dir, "other")
		dsquare_____________ := filepath.Join(dir, "square")
		dbadquality_________ := filepath.Join(dir, "bad_quality")
		dbadqualitylandscape := filepath.Join(dbadquality_________, "l")
		dbadqualitysquare___ := filepath.Join(dbadquality_________, "s")
		dbadqualityportrait_ := filepath.Join(dbadquality_________, "p")
		dvideo______________ := filepath.Join(dir, "video")
		Destinations := []string{
			dwall_______________,
			dother______________,
			dsquare_____________,
			dbadquality_________,
			dvideo______________,
			dbadqualitylandscape,
			dbadqualitysquare___,
			dbadqualityportrait_,
		}
		createDirectories(Destinations)
		cMoveFile(filepath.Join(dir, e.Name()), Destinations)
	}
}

func cMoveFile(path string, dests []string) {
	d_wall := dests[0]
	d_other := dests[1]
	d_square := dests[2]
	d_badquality := dests[3]
	d_video := dests[4]
	d_landscape_badquality := dests[5]
	d_square_badquality := dests[6]
	d_portrait_badquality := dests[7]

	ext := filepath.Ext(path)
	var _ string = string(ext)

	if ext == ".mp4" {
		moveFiles(path, d_video, "yellow", "video")
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
				moveFiles(path, d_wall, "red", "land")
			} else if ar == 1 {
				moveFiles(path, d_square, "blue", "square")
			} else if ar < 1 {
				moveFiles(path, d_other, "green", "portrait")
			}
		} else {
			if ar > 1 {
				moveFiles(path, d_landscape_badquality, "cyan", "land")
			} else if ar == 1 {
				moveFiles(path, d_square_badquality, "magenta", "square")
			} else if ar < 1 {
				moveFiles(path, d_portrait_badquality, "purple", "portrait")
			} else {
				moveFiles(path, d_badquality, "grey", "")
			}
		}
	}
}

func moveFiles(source, dest, name, special_char string) error {
	Count++
	f_name := filepath.Base(source)
	d_path := filepath.Join(dest, f_name)

	err := os.Rename(source, d_path)
	if err != nil {
		Failed++
		errormaxxing(err.Error())
		return err
	} else {
		Succeed++
		logmaxxing(source, dest, name, special_char)
		return nil
	}

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

	println(tab, c, special, tab, padpar, "<| |=| |>", tab, file_path, reset_white)
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

func createDirectories(destinationsFolders []string) {
	for _, s := range destinationsFolders {
		err := os.MkdirAll(s, 0755)
		if err != nil {
			log.Fatal("error making directories")
		}
	}
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
