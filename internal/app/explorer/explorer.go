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

func (fe *FileExplorer) Ls(path string) ([]string, error) {
	if path == "" {
		return fe.ls(fe.Pwd())
	}

	if path == "../" {
		dir, _ := filepath.Split(fe.Pwd())
		return fe.ls(dir)
	}

	return fe.ls(fe.Abs(path))
}

func (fe *FileExplorer) ls(path string) ([]string, error) {
	list := []string{}

	if path != "/" {
		list = append(list, "../")
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if f.IsDir() {
			list = append(list, f.Name()+"/")
		} else {
			list = append(list, f.Name())
		}
	}

	return list, nil
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
