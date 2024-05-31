package main

import (
	"log"
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

	exit_message()
}
