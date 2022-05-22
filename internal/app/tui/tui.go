package tui

import (
	"gofm/internal/app/explorer"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TUI struct {
	app           *tview.Application
	leftPane      *tview.List
	leftExplorer  explorer.FileExplorer
	rightPane     *tview.List
	rightExplorer explorer.FileExplorer
	footer        *tview.TextView
}

func NewTUI(newExplorer explorer.FileExplorer) TUI {
	t := TUI{
		app:           tview.NewApplication(),
		leftExplorer:  newExplorer,
		rightExplorer: newExplorer,
	}

	t.initFooter()
	t.initLeftPane()
	t.initRightPane()
	t.initApp()

	return t
}

func (t *TUI) Run() error {
	return t.app.Run()
}

func (t *TUI) initFooter() {
	t.footer = tview.NewTextView().
		SetText("").
		SetTextColor(tcell.ColorGreen)
}

func (t *TUI) initLeftPane() {
	t.leftPane = tview.NewList().
		SetWrapAround(false).
		SetHighlightFullLine(true).
		SetSelectedFocusOnly(true).
		ShowSecondaryText(false).
		SetMainTextColor(tcell.ColorNavy).
		SetSelectedBackgroundColor(tcell.ColorGreen)

	t.leftPane.SetTitle(t.leftExplorer.Pwd()).
		SetBorder(true).
		SetBorderColor(tcell.ColorGreen)

	for _, item := range t.leftExplorer.Ls(t.leftExplorer.Pwd()) {
		t.leftPane.AddItem(item, "", 0, nil)
	}

	t.leftPane.SetSelectedFunc(t.leftPaneSelected).
		SetChangedFunc(t.leftPaneChanged).
		SetInputCapture(t.leftPaneEvents)
}

func (t *TUI) initRightPane() {
	t.rightPane = tview.NewList().
		SetWrapAround(false).
		SetHighlightFullLine(true).
		SetSelectedFocusOnly(true).
		ShowSecondaryText(false).
		SetMainTextColor(tcell.ColorNavy).
		SetSelectedBackgroundColor(tcell.ColorGreen)

	t.rightPane.SetTitle(t.rightExplorer.Pwd()).
		SetBorder(true).
		SetBorderColor(tcell.ColorGreen)

	for _, item := range t.rightExplorer.Ls(t.rightExplorer.Pwd()) {
		t.rightPane.AddItem(item, "", 0, nil)
	}

	t.rightPane.SetSelectedFunc(t.rightPaneSelected).
		SetChangedFunc(t.rightPaneChanged).
		SetInputCapture(t.rightPaneEvents)
}

func (t *TUI) initApp() {
	panes := tview.NewFlex().
		AddItem(t.leftPane, 0, 1, true).
		AddItem(t.rightPane, 0, 1, false)

	menu := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(panes, 0, 15, true).
		AddItem(t.footer, 0, 1, false)

	t.app.SetRoot(menu, true).
		EnableMouse(true).
		SetInputCapture(t.appEvents)
}
