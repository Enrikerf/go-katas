package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%v \n", removeFolder([]string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f", "/c/f/g"}, "c"))
	fmt.Printf("%v \n", removeFolder([]string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f", "/c/f/g"}, "d"))
}

func removeFolder(folderStructure []string, folderName string) (result []string) {
	var newFolderStructure []string
	for _, element := range folderStructure {
		filteredElement := seekAndDestroy(element, folderName)
		if len(filteredElement) > 0 {
			newFolderStructure = append(newFolderStructure, filteredElement)
		}
	}
	result = newFolderStructure
	return
}

//🤘
func seekAndDestroy(element string, folderName string) string {
	firstOccurrenceIndex := strings.Index(element, folderName)
	if firstOccurrenceIndex > 0 {
		return element[:firstOccurrenceIndex-1]
	} else {
		return element
	}
}
