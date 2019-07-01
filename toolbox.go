package toolbox

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

// Version: 01.07.2019

// Creates a actually TempusString in the Format: "day-month-year_hour-min-sec"
func TempusString() string {
	now := time.Now()
	y, m, d := now.Date()
	h, min, s := now.Clock()
	return (fmt.Sprint(d, "-", m, "-", y, "_", h, "-", min, "-", s))
}

// Creates a file with the content in Mode 0700
// only the creater can read.
func StoreAStringInAFile(filename, content string) (err error) {
	contentInByte := []byte(content)
	return ioutil.WriteFile("/home/ubuntu/brooker/"+filename, contentInByte, 0700)
}

func GetContentOfAFile(filename string) (content string, err error) {
	bytes, err := ioutil.ReadFile("/home/ubuntu/brooker/" + filename)
	content = string(bytes)
	return
}

// Rturns all Files with the ending of suffix. If suffix is nil all filnames will be in the Array
func GetFileFromDir(directory string, suffix interface{}) (allFiles []string, err error) {
	files, err := ioutil.ReadDir("/home/ubuntu/brooker/" + directory)
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

func TestingNil(err error, ret bool) {
	if err != nil {
		fmt.Println(err)
		if ret {
			return
		}
	}
}

func TestingNilAndLog(l *log.Logger, err error, ret bool) {
	if err != nil {
		l.Println(err)
		if ret {
			return
		}
	}
}
