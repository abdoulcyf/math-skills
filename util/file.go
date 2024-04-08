package util

import (
	"io"
	"os"
)

func ReadFileToStr(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "Error Opening file", err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	//dataArr := make([]string, 0)

	data, err := io.ReadAll(f)
	if err != nil {
		return "Error Reading file", err
	}

	//for _, line := range strings.Split(string(data), "\n") {
	//	dataArr = append(dataArr, line)
	//}
	return string(data), nil
}

//func ConvertToStrArr() ([]string, error) {
//
//	data, err := ReadFileToStr("data.txt")
//	if err != nil {
//		return nil, err
//	}
//	dataArr := make([]string, 0)
//	for i, line := range strings.Split(string(data), "\n") {
//		dataArr = append(dataArr, line)
//	}
//}
