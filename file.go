package toolbox

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type File struct {
	Mainpath string
}

// Creates a file with the content in Mode 0700
// only the creater can read.
func (a *File) StoreAStringInAFile(filename, content string) (err error) {
	contentInByte := []byte(content)
	return ioutil.WriteFile(a.Mainpath+filename, contentInByte, 0700)
}

func (a *File) GetContentOfAFile(filename string) (content string, err error) {
	bytes, err := ioutil.ReadFile(a.Mainpath + filename)
	content = string(bytes)
	return
}

// Rturns all Files with the ending of suffix. If suffix is nil all filnames will be in the Array
func (a *File) GetFileFromDir(directory string, suffix interface{}) (allFiles []string, err error) {
	files, err := ioutil.ReadDir(a.Mainpath + directory)
	var tmp string
	if suffix != nil {
		tmp = fmt.Sprint(suffix)
	}
	for _, file := range files {
		if suffix != nil {
			if strings.HasSuffix(file.Name(), tmp) || strings.HasSuffix(file.Name(), strings.ToLower(tmp)) || strings.HasSuffix(file.Name(), strings.ToUpper(tmp)) {
				allFiles = append(allFiles, file.Name())
			}
		} else {
			allFiles = append(allFiles, file.Name())
		}
	}
	return
}
