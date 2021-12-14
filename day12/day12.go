package day12

import (
	"fmt"
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
	fmt.Println(caveMap)

	var paths []string
	var traversal []string
	traverse("start", traversal, caveMap, &paths)
	fmt.Println()
	fmt.Println(paths)

	return len(paths)
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
			if strings.ToLower(c) == c && alreadyContainedInTraversal(c, traversal) {
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
