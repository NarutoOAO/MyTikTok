package util

import (
	"MyTikTok/conf"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

// UploadVideoToLocalStatic upload video
func UploadVideoToLocalStatic(file multipart.File, userId int64, filename string) (filePath string, err error) {
	finalName := fmt.Sprintf("%d_%s", userId, filename)
	filePath = filepath.Join(".", conf.VideoPath, finalName)
	basePath := "." + conf.VideoPath
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(filePath, content, 0666)
	if err != nil {
		return "", err
	}
	return finalName, err
}

// DirExistOrNot
func DirExistOrNot(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		log.Println(err)
		return false
	}
	return s.IsDir()
}

// CreateDir
func CreateDir(dirName string) bool {
	err := os.MkdirAll(dirName, 0755)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
