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
