package day12

import (
	"strings"
)

func PossiblePaths(input string) int {
	caveMap := make(map[string][]string)
	for _, line := range strings.Split(input, "\n") {
		c1 := strings.Split(line, "-")[0]
		c2 := strings.Split(line, "-")[1]
		caveMap[c1] = append(caveMap[c1], c2)
		caveMap[c2] = append(caveMap[c2], c1)
	}

	var paths []string
	var traversal []string
	traverse("start", traversal, caveMap, &paths)

	return len(paths)
}

func PossiblePaths2(input string) int {
	caveMap := make(map[string][]string)
	for _, line := range strings.Split(input, "\n") {
		c1 := strings.Split(line, "-")[0]
		c2 := strings.Split(line, "-")[1]
		caveMap[c1] = append(caveMap[c1], c2)
		caveMap[c2] = append(caveMap[c2], c1)
	}

	var paths []string
	var traversal []string
	traverse2("start", traversal, caveMap, &paths)

	return len(paths)
}

func traverse2(cave string, traversal []string, caveMap map[string][]string, paths *[]string) {
	traversal = append(traversal, cave)
	if cave == "end" {
		traversalStr := strings.Join(traversal, ",")
		if !alreadyContainedInPaths(traversalStr, paths) {
			*paths = append(*paths, traversalStr)
		}
	} else {
		for _, c := range caveMap[cave] {
			if c == "start" {
				continue
			}
			if isLowerCase(c) && !isAllowed(c, traversal) {
				continue
			}
			traverse2(c, traversal, caveMap, paths)
		}
	}
}

func isAllowed(cave string, traversal []string) bool {
	lowerCaveCount := make(map[string]int)
	for _, c := range traversal {
		if isLowerCase(c) {
			lowerCaveCount[c] += 1
		}
	}
	if alreadyContainedInTraversal(cave, traversal) {
		for _, v := range lowerCaveCount {
			if v > 1 {
				return false
			}
		}
	}
	return true
}

func traverse(cave string, traversal []string, caveMap map[string][]string, paths *[]string) {
	traversal = append(traversal, cave)
	if cave == "end" {
		traversalStr := strings.Join(traversal, ",")
		if !alreadyContainedInPaths(traversalStr, paths) {
			*paths = append(*paths, traversalStr)
		}
	} else {
		for _, c := range caveMap[cave] {
			if c == "start" {
				continue
			}
			if isLowerCase(c) && alreadyContainedInTraversal(c, traversal) {
				continue
			}
			traverse(c, traversal, caveMap, paths)
		}
	}
}

func alreadyContainedInPaths(traversal string, paths *[]string) bool {
	for _, p := range *paths {
		if p == traversal {
			return true
		}
	}
	return false
}

func alreadyContainedInTraversal(cave string, traversal []string) bool {
	for _, c := range traversal {
		if c == cave {
			return true
		}
	}
	return false
}

func isLowerCase(s string) bool {
	return strings.ToLower(s) == s
}
