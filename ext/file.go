package ext

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// MAXLENGTH Maximum length of file name
const MAXLENGTH = 80

// FileSize return the file size of the specified path file
func FileSize(filePath string) (int64, bool, error) {
	file, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, false, nil
		}
		return 0, false, err
	}
	return file.Size(), true, nil
}

// LimitLength Handle overly long strings
func LimitLength(s string, length int) string {
	const ELLIPSES = "..."
	str := []rune(s)
	if len(str) > length {
		return string(str[:length-len(ELLIPSES)]) + ELLIPSES
	}
	return s
}

// FileName Converts a string to a valid filename
func FileName(name string, ext string) string {
	rep := strings.NewReplacer("\n", " ", "/", " ", "|", "-", ": ", "：", ":", "：", "'", "’")
	name = rep.Replace(name)
	if runtime.GOOS == "windows" {
		rep = strings.NewReplacer("\"", " ", "?", " ", "*", " ", "\\", " ", "<", " ", ">", " ")
		name = rep.Replace(name)
	}
	limitedName := LimitLength(name, MAXLENGTH)
	if ext == "" {
		return limitedName
	} else {
		return fmt.Sprintf("%s.%s", limitedName, ext)
	}
}

// PathIsExist  files and folders exist
func PathIsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

// MkDir create directory
func MkDir(path string) (bool, error) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return false, err
	}
	return true, nil
}
