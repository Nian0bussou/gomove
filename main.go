package main

import (
	"log"
	"os/exec"
	"runtime"
	"sync"
)

func main() {
	defer timeminmaxxing()()

	path := get_path()
	var Choice int
	Choice, path = get_choice(path)
	subs := get_folders(path)
	if subs == nil {
		println("could not get directories")
		return
	}
	for _, f := range subs {
		println(f)
	}

	line_print()

	//////////////////////////////////////////////////////////////////////////////
	var wg sync.WaitGroup
	for _, item := range subs {
		wg.Add(1)
		go func(item string) {
			defer wg.Done()
			switch Choice {
			case 0:
				move_wrapper(item)
			case 1:
				scramble(item)
			default:
				log.Fatal("invalid option")
			}
		}(item)
	}
	wg.Wait()
	//////////////////////////////////////////////////////////////////////////////

	removeTmp(path)
	line_print()

	println("Finished,")
	println("count     : ", Count)
	println("succeeded : ", Succeed)
	println("failed    : ", Failed)
	println("portrait  : ", Portrait_count)
	println("landscape : ", Landscape_count)
	println("square    : ", Square___count)
	println("video     : ", Video____count)

	///////////////////////////////////////////////////////////////////////////////
	if runtime.GOOS == "windows" {
		exec.Command("explorer.exe", "D:\\grapper\\").Run()
	}
	//////////////////////////////////////////////////////////////////////////////

	return
}
