package stuff

import (
	"bytes"
	"errors"
	"os"
)

func OpenFile() ([]byte, error) {
	var err error
	if len(os.Args) < 2 {
		err = errors.New("please provide a file name")
		panic(err)
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		err = errors.New("cant read file")
		panic(err)
	}
	return data, err
}
func SaveFile(lines []string) error {
	data := StrToByte(lines)
	if len(os.Args) < 2 {
		err := errors.New("please provide a file name")
		panic(err)
	}
	err := os.WriteFile(os.Args[1], data, 0644)
	if err != nil {
		err = errors.New("cant write file")
		panic(err)
	}
	return nil
}
func FileConvert(data []byte) []string {
	//convert file to array of strings
	lines := make([]string, 0)
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		lines = append(lines, string(line))
	}

	return lines

}
func StrToByte(lines []string) []byte {
	//convert array of strings to byte array
	data := make([]byte, 0)
	for i := 0; i < len(lines); i++ {
		data = append(data, []byte(lines[i])...)
		data = append(data, '\n')
	}
	return data

}