package aoc22

import (
	"cmp"
	"math"
	"slices"
	"strings"

	. "github.com/zeroibot/aoc-go/aoc"
	"github.com/zeroibot/fn/dict"
	"github.com/zeroibot/fn/list"
	"github.com/zeroibot/fn/number"
	"github.com/zeroibot/fn/str"
)

type FileSystem = map[string]*Item

type Item struct {
	isDir    bool
	name     string
	path     string
	size     int
	parent   *Item
	children []*Item
}

var (
	cmdCD    string = "$ cd"
	cmdLS    string = "$ ls"
	pathGlue string = "/"
)

func Day07() Solution {
	fs := data07(true)

	// Part 1
	total := 0
	for _, item := range fs {
		if item.isDir && item.size <= 100_000 {
			total += item.size
		}
	}

	// Part 2
	totalSize := 70_000_000
	required := 30_000_000
	free := totalSize - fs["/"].size
	minimum := required - free

	minSize := math.MaxInt
	for _, item := range fs {
		if item.isDir && item.size >= minimum {
			minSize = min(minSize, item.size)
		}
	}

	return NewSolution(total, minSize)
}

func data07(full bool) FileSystem {
	fs := make(FileSystem)
	var cwd *Item = nil
	for _, line := range ReadLines(22, 7, full) {
		if strings.HasPrefix(line, cmdCD) {
			name := str.SpaceSplit(line)[2]
			if name == ".." && cwd != nil {
				cwd = cwd.parent
			} else {
				cwd, _ = getDir(fs, name, cwd)
			}
		} else if line == cmdLS {
			continue
		} else {
			p := str.SpaceSplit(line)
			head, tail := p[0], p[1]
			var item *Item
			var isNew bool
			if head == "dir" {
				item, isNew = getDir(fs, tail, cwd)
			} else {
				item, isNew = getFile(fs, tail, number.ParseInt(head), cwd)
			}
			if isNew && cwd != nil {
				cwd.children = append(cwd.children, item)
			}
		}
	}

	dirPaths := list.Filter(dict.Keys(fs), func(path string) bool {
		return fs[path].isDir
	})
	slices.SortFunc(dirPaths, func(p1, p2 string) int {
		len1 := len(strings.Split(p1, pathGlue))
		len2 := len(strings.Split(p2, pathGlue))
		return cmp.Compare(len2, len1)
	})
	for _, path := range dirPaths {
		item := fs[path]
		item.size = item.computeSize()
		fs[path] = item
	}
	return fs
}

func getDir(fs FileSystem, name string, parent *Item) (*Item, bool) {
	path := name
	if parent != nil {
		path = parent.path + name + pathGlue
	}
	item, ok := fs[path]
	if ok {
		return item, false
	}
	item = &Item{
		isDir:    true,
		name:     name,
		path:     path,
		size:     0,
		parent:   parent,
		children: []*Item{},
	}
	fs[path] = item
	return item, true
}

func getFile(fs FileSystem, name string, size int, parent *Item) (*Item, bool) {
	path := name
	if parent != nil {
		path = parent.path + name
	}
	item, ok := fs[path]
	if ok {
		return item, false
	}
	item = &Item{
		isDir:    false,
		name:     name,
		path:     path,
		size:     size,
		parent:   parent,
		children: nil,
	}
	fs[path] = item
	return item, true
}

func (i Item) computeSize() int {
	if len(i.children) == 0 {
		return 0
	}
	size := 0
	for _, child := range i.children {
		size += child.size
	}
	return size
}
