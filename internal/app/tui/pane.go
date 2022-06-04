package tui

import (
	"gofm/internal/app/explorer"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Pane struct {
	currentFile string
	list        *tview.List
	explorer    explorer.FileExplorer
}

func NewPane(pane *tview.List, explorer explorer.FileExplorer) Pane {
	return Pane{
		list:     pane,
		explorer: explorer,
	}
}

func (p *Pane) Init() {
	p.list.SetWrapAround(false).
		SetHighlightFullLine(true).
		SetSelectedFocusOnly(true).
		ShowSecondaryText(false).
		SetMainTextColor(tcell.ColorNavy).
		SetSelectedBackgroundColor(tcell.ColorGreen).
		SetTitle(p.explorer.Pwd()).
		SetBorder(true).
		SetBorderColor(tcell.ColorGreen)

	fileList := p.explorer.Ls(p.explorer.Pwd())

	for _, item := range fileList {
		p.list.AddItem(item, "", 0, nil)
	}

	p.SetCurrentFile(fileList[0])
}

func (p *Pane) List() *tview.List {
	return p.list
}

func (p *Pane) ChangeDirectoryIfNeed(path string) {
	if !p.explorer.IsDir(path) {
		return
	}

	p.list.Clear()

	p.explorer.Cd(path)

	p.list.SetTitle(p.explorer.Pwd())

	list := p.explorer.Ls("")
	for _, item := range list {
		p.list.AddItem(item, "", 0, nil)
	}
}

func (p *Pane) SetCurrentFile(name string) {
	p.currentFile = name
}

func (p *Pane) GetCurrentFile() string {
	return p.currentFile
}

func (p *Pane) FileInfo(fileName string) string {
	return p.explorer.FileInfoString(fileName)
}

func (p *Pane) CurrentFileInfo() string {
	return p.explorer.FileInfoString(p.currentFile)
}
