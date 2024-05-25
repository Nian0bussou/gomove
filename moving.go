package main

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
)

var (
	dwall                string
	dother               string
	dsquare              string
	dbadquality          string
	dbadqualitylandscape string
	dbadqualitysquare    string
	dbadqualityportrait  string
	dvideo               string
)

func move_wrapper(dir string) {
	entries, _ := os.ReadDir(dir)
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		dwall = filepath.Join(dir, "wall")
		dother = filepath.Join(dir, "other")
		dsquare = filepath.Join(dir, "square")
		dbadquality = filepath.Join(dir, "bad_quality")
		dbadqualitylandscape = filepath.Join(dbadquality, "l")
		dbadqualitysquare = filepath.Join(dbadquality, "s")
		dbadqualityportrait = filepath.Join(dbadquality, "p")
		dvideo = filepath.Join(dir, "video")
		Destinations := []string{
			dwall,
			dother,
			dsquare,
			dbadquality,
			dvideo,
			dbadqualitylandscape,
			dbadqualitysquare,
			dbadqualityportrait,
		}
		createDirectories(Destinations)
		move_file(filepath.Join(dir, e.Name()))
	}
}

func move_file(path string) {

	ext := filepath.Ext(path)
	var _ string = string(ext) // not work otherwise

	if ext == ".mp4" {
		moveFiles(path, dvideo, "yellow", "video")
		return
	}
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
			moveFiles(path, dwall, "red", "land")
		} else if ar == 1 {
			moveFiles(path, dsquare, "blue", "square")
		} else if ar < 1 {
			moveFiles(path, dother, "green", "portrait")
		}
	} else {
		if ar > 1 {
			moveFiles(path, dbadqualitylandscape, "cyan", "land")
		} else if ar == 1 {
			moveFiles(path, dbadqualitysquare, "magenta", "square")
		} else if ar < 1 {
			moveFiles(path, dbadqualityportrait, "purple", "portrait")
		} else {
			moveFiles(path, dbadquality, "grey", "")
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
