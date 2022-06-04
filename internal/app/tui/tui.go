package tui

import (
	"log/syslog"

	"gofm/internal/app/explorer"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TUI struct {
	logger *syslog.Writer //TODO temporary for debugging

	app       *tview.Application
	leftPane  Pane
	rightPane Pane
	footer    *tview.TextView
}

func NewTUI(newExplorer explorer.FileExplorer, logger *syslog.Writer) TUI {
	t := TUI{
		app:       tview.NewApplication(),
		leftPane:  NewPane(tview.NewList(), newExplorer),
		rightPane: NewPane(tview.NewList(), newExplorer),
		logger:    logger,
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
		SetText(t.leftPane.CurrentFileInfo()).
		SetTextColor(tcell.ColorGreen)
}

func (t *TUI) initLeftPane() {
	t.leftPane.Init()

	t.leftPane.List().
		SetSelectedFunc(t.leftPaneSelected).
		SetChangedFunc(t.leftPaneChanged).
		SetInputCapture(t.leftPaneEvents)
}

func (t *TUI) initRightPane() {
	t.rightPane.Init()

	t.rightPane.List().
		SetSelectedFunc(t.rightPaneSelected).
		SetChangedFunc(t.rightPaneChanged).
		SetInputCapture(t.rightPaneEvents)
}

func (t *TUI) initApp() {
	panes := tview.NewFlex().
		AddItem(t.leftPane.List(), 0, 1, true).
		AddItem(t.rightPane.List(), 0, 1, false)

	menu := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(panes, 0, 15, true).
		AddItem(t.footer, 0, 1, false)

	t.app.SetRoot(menu, true).
		EnableMouse(true).
		SetInputCapture(t.appEvents)
}
