package main

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"sync"

	"golang.org/x/term"
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

	return
}
