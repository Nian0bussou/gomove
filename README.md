
# GoMove


this is a program to categorize images by their width, height & ratio


Since, it is made for doing multiple directories in parallel by default it is using this patern:


your-path/

├── subdir1

├── subdir2

└── subdir3


where "your-path/" is the default path ("D:/grapper" or "/mnt/d/grapper/") or the one specified  

it will __only__ sort the subdirectories (in this example: subdir1, subdir2 & subdir3)


by default it will organize the category using this patern:


subdir1/

├── bad_quality

│   ├── l

│   ├── p

│   └── s

├── other

├── square

├── video

└── wall

subdir2/
...

---

# Usage


`go build` will build the binary file

`gomove <optional : choice> <optional : path> `

(necessary to provide 'choice' if providing the path)

(need to be in the order as written)



### choice
- 0 => sort the files
- 1 => scramble the files into their respective root subdir


### path
if the path is specified, it will use it as the root dir (where all the subdirs are located)

 however **NO VERIFICATION IS DONE** (and by that I mean that it will just exit if it cant find any available directories), make sure the path is correct




