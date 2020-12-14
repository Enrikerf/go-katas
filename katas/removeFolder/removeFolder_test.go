package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveFolder(testing *testing.T) {
	// requisites: folders name contains only letters
	// factors
	//  * folder don't exists in folder structure
	//  * folder exists in folder structure
	// equivalence classes
	//  * all FolderStructure without folderName
	//  * all FolderStructure with one folderName 
	// limit values: 
	//   * folderName its root, 
	//   * folderName in middle, 
	//   * folderName its end, 
	var valuesToTest = [] struct {
		folderStructure []string
		folderName string
		result []string
	}{
		{[]string{"/a","/a/b","/c/d","/c/d/e","/c/f", "/c/f/g"},"c",[]string{"/a","/a/b"}}, // example case of cassidoo test
		{[]string{"/a","/folderName/b","/d/folderName"},"c",[]string{"/a","/folderName/b","/d/folderName"}}, // without folderName in FolderStructure
		{[]string{"/a","/c/b/a","/r/g/b/c","/d/folderName/c/v"},"c",[]string{"/a","/r/g/b","/d/folderName"}}, // with folderName in FolderStructure all limit cases
	}
	for i := 0; i < len(valuesToTest); i++ {
		assert.Equal(testing, valuesToTest[i].result, removeFolder(valuesToTest[i].folderStructure,valuesToTest[i].folderName), "should be equal")
	}

}
