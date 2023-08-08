package main

import (
	"bufio"
	"fmt"
)

type fs struct {
	size   int
	fsTree map[string]*fs
}

func recurseDir(dir *fs, lines *bufio.Scanner) int {
	for lines.Scan() {
		switch lines.Text()[0] {
		case '$':
			var cmd, opts string
			fmt.Sscanf(lines.Text(), "$ %s %s", &cmd, &opts)
			if cmd == "cd" {
				if opts != ".." {
					dir.size += recurseDir(dir.fsTree[opts], lines)
				} else {
					return dir.size
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
	return dir.size
}
func main() {
	// sol1()
	sol2()
}
