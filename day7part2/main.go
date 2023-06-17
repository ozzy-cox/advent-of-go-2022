package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type (
	stack []frame
	item  struct {
		name  string
		_type string
		size  int
	}
	frame struct {
		currentDir string
		size       int
		contents   map[string]item
	}
)

func (s *stack) Push(v frame) {
	*s = append(*s, v)
}

func (s *stack) Pop() (frame, error) {
	if len(*s) == 0 {
		return frame{}, errors.New("empty")
	}
	l := len(*s)
	ret := (*s)[l-1]
	*s = (*s)[:l-1]
	return ret, nil
}

func (s *stack) Peek() (frame, error) {
	if len(*s) == 0 {
		return frame{}, errors.New("empty")
	}
	return (*s)[len(*s)-1], nil
}

const (
	THRESHOLD            int = 100_000
	TOTAL_CAP            int = 70_000_000
	NECESSARY_FOR_UPDATE int = 30_000_000
)

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	frames := make(stack, 0)

	dirsBelowThreshold := make([]item, 0)
	allDirs := make([]item, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "$ cd .." {
			// Step out, pop frame, return sum of contents
			poppedFrame, err := frames.Pop()
			if err != nil {
				panic(err)
			}
			currentFrame, err := frames.Peek()
			if err != nil {
				panic("error")
			}

			size := 0

			for _, value := range poppedFrame.contents {
				size += value.size
			}

			if item, ok := currentFrame.contents[poppedFrame.currentDir]; ok {

				// Then we modify the copy
				item.size = size
				allDirs = append(allDirs, item)
				if size < THRESHOLD {
					dirsBelowThreshold = append(dirsBelowThreshold, item)
				}

				// Then we reassign map entry
				currentFrame.contents[poppedFrame.currentDir] = item
			}

			continue
		}

		if line[0:4] == "$ cd" {
			_frame := frame{
				currentDir: strings.Split(line, " ")[2],
				size:       0,
				contents:   make(map[string]item),
			}

			frames.Push(_frame)
			continue
		} else if line[0:4] == "$ ls" {
			continue
		}

		// Add the dir contents to the current frame
		currentFrame, err := frames.Peek()
		if err != nil {
			continue
		}

		lineParts := strings.Split(line, " ")
		isDir := lineParts[0] == "dir"

		_type := "dir"
		size := 0
		name := lineParts[1]
		if !isDir {
			_type = "file"
			size, _ = strconv.Atoi(lineParts[0])
		}

		currentFrame.contents[name] = item{
			name:  name,
			_type: _type,
			size:  size,
		}
	}

	var poppedFrame frame
	var currentFrame frame
	var err error
	for len(frames) > 0 {
		sum := 0
		for _, value := range dirsBelowThreshold {
			sum += value.size
		}
		// Step out, pop frame, return sum of contents
		poppedFrame, _ = frames.Pop()
		currentFrame, err = frames.Peek()
		if err != nil {
			break
		}
		size := 0
		for _, value := range poppedFrame.contents {
			size += value.size
		}
		if item, ok := currentFrame.contents[poppedFrame.currentDir]; ok {
			// Then we modify the copy
			item.size = size
			if size < THRESHOLD {
				dirsBelowThreshold = append(dirsBelowThreshold, item)
			}
			// Then we reassign map entry
			currentFrame.contents[poppedFrame.currentDir] = item
		}
	}

	totalSize := 0

	for _, value := range poppedFrame.contents {
		totalSize += value.size
	}
	fmt.Println(poppedFrame.currentDir)
	fmt.Println(poppedFrame.contents)

	spaceToClear := totalSize - (TOTAL_CAP - NECESSARY_FOR_UPDATE)
	fmt.Println("spacetoclear", spaceToClear)
	for _, value := range poppedFrame.contents {
		// fmt.Println(value)
		allDirs = append(allDirs, value)
	}

	var minSizeDir item
	for _, value := range allDirs[1:] {
		if value.size > spaceToClear {
			if minSizeDir == (item{}) {
				minSizeDir = value
			} else if value.size < minSizeDir.size {
				minSizeDir = value
			}
		}
	}
	fmt.Println(minSizeDir)
}
