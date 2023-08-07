package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type fs struct {
	size   int
	fsTree map[string]*fs
}

func sol1() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	sc := bufio.NewScanner(f)
	sc.Scan()
	var fileDirectory = &fs{size: 0, fsTree: make(map[string]*fs)}
	recurseDir(fileDirectory, sc)

}

func recurseDir(dir *fs, lines *bufio.Scanner) {
	for lines.Scan() {
		fmt.Println(lines.Text())
		switch lines.Text()[0] {
		case '$':
			var cmd, opts string
			fmt.Sscanf(lines.Text(), "$ %s %s", &cmd, &opts)
			// fmt.Println(cmd, opts)
			if cmd == "cd" {
				if opts != ".." {
					recurseDir(dir.fsTree[opts], lines)
				} else {
					return
				}
			}
		case 'd':
			var dirName string
			fmt.Sscanf(lines.Text(), "dir %s", &dirName)
			dir.fsTree[dirName] = &fs{size: 0, fsTree: make(map[string]*fs)}
		default:
			var size int
			fmt.Sscanf(lines.Text(), "%d %s", &size)
			dir.size += size
		}
	}
}

func traverseFs(dir *fs) {
	for k, v := range fs.fsTree {
		traverseFs(v)
	}
}
