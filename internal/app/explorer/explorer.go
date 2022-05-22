package explorer

import (
	"fmt"
	"gofm/pkg/utils"
	"io/ioutil"
	"log/syslog"
	"os"
	"path/filepath"
)

type FileExplorer struct {
	currentPath string
	logger      *syslog.Writer
}

func NewFileExplorer(logger *syslog.Writer) FileExplorer {
	return FileExplorer{
		logger: logger,
	}
}

func (fe *FileExplorer) Cd(path string) string {
	newPath := fe.Abs(path)

	fe.currentPath = filepath.Dir(newPath)

	return fe.currentPath
}

func (fe *FileExplorer) Abs(path string) string {
	return fmt.Sprintf("%s/%s", fe.Pwd(), path)
}

func (fe *FileExplorer) Pwd() string {
	if fe.currentPath != "" {
		return fe.currentPath
	}

	path, err := os.Getwd()
	if err != nil {
		fe.logger.Err(err.Error())
		return ""
	}

	fe.currentPath = path

	return path
}

func (fe *FileExplorer) IsDir(path string) bool {
	absPath := fe.Abs(path)

	fileInfo, err := os.Stat(absPath)
	if err != nil {
		fe.logger.Err(err.Error())
		return false
	}

	return fileInfo.IsDir()
}

func (fe *FileExplorer) Ls(path string) []string {
	if path != "" {
		return fe.ls(path)
	}

	return fe.ls(fe.Pwd())
}

func (fe *FileExplorer) ls(path string) []string {
	list := []string{}

	if path != "/" {
		list = append(list, "../")
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		fe.logger.Err(err.Error())
		return nil
	}

	for _, f := range files {
		if f.IsDir() {
			list = append(list, f.Name()+"/")
		} else {
			list = append(list, f.Name())
		}
	}

	return list
}

func (fe *FileExplorer) FileInfoString(path string) string {
	absPath := fe.Abs(path)

	fileInfo, err := os.Stat(absPath)
	if err != nil {
		fe.logger.Err(err.Error())
		return ""
	}

	size, unit := utils.ConvertBytes(fileInfo.Size())

	return FileInfo{
		Mode:       fileInfo.Mode(),
		ModeTime:   fileInfo.ModTime(),
		Size:       size,
		MemoryUnit: unit,
	}.String()
}
