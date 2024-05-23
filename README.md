
# GoMove

---

this is a program to categorize images by their width, height & ratio


Since, it is made for doing multiple directories in parallel

by default it is using this patern:

your-path/

├── subdir1

├── subdir2

└── subdir3


where "your-path/" is the default path ("D:/grapper" or "/mnt/d/grapper/") or the one specified  

it will **only** sort the subdirectories (in this example: subdir1, subdir2 & subdir3)


by default it will organize the category using this patern:


subdir/

├── bad_quality

│   ├── l

│   ├── p

│   └── s

├── other

├── square

├── video

└── wall

---

for my convenience, it will also open the File Explorer if it is running on Windows

---

# Usage


`go build` will build the binary file

`gomove <optional : choice> <optional : path> `


only using `go run main.go`

use the defaults

choice ; 0 : organize the files

choice ; 1 : scramble the files into their respecting root subdirectory 

if the path is specified, it will use it as the root dir (where all the subdirs are located)

however **NO VERIFICATION IS DONE** (and by that I mean that it will just exit if it cant find any available directories), make sure the path is correct



(the structure of the directories in this README were made using the 'tree' package)

