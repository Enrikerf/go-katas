package main

func main() {
	println(removeFolder([]string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f", "/c/f/g"}, "c"))
}

func removeFolder(folderStructure []string, folderName string) (result string) {
	for _, element := range folderStructure {
		println(element)
	}
	return
}
