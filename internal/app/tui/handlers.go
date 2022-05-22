package tui

import "github.com/gdamore/tcell/v2"

func (t *TUI) appEvents(event *tcell.EventKey) *tcell.EventKey {
	if event.Rune() == 'q' {
		t.app.Stop()
	}

	return event
}

func (t *TUI) rightPaneEvents(event *tcell.EventKey) *tcell.EventKey {
	if event.Rune() == rune(tcell.KeyTab) {
		t.app.SetFocus(t.leftPane)
		return nil
	}

	if event.Key() == 259 {
		return nil
	}

	if event.Key() == 260 {
		return nil
	}

	return event
}

func (t *TUI) leftPaneEvents(event *tcell.EventKey) *tcell.EventKey {
	if event.Rune() == rune(tcell.KeyTab) {
		t.app.SetFocus(t.rightPane)
		return nil
	}

	if event.Key() == 259 {
		return nil
	}

	if event.Key() == 260 {
		return nil
	}

	return event
}

func (t *TUI) rightPaneSelected(i int, s string, s2 string, r rune) {
	if !t.rightExplorer.IsDir(s) {
		return
	}

	t.rightPane.Clear()

	t.rightExplorer.Cd(s)

	t.rightPane.SetTitle(t.rightExplorer.Pwd())

	list := t.rightExplorer.Ls("")
	for _, item := range list {
		t.rightPane.AddItem(item, "", 0, nil)
	}
}

func (t *TUI) rightPaneChanged(index int, mainText string, secondaryText string, shortcut rune) {
	t.footer.SetText(t.rightExplorer.FileInfoString(mainText))
}

func (t *TUI) leftPaneSelected(i int, s string, s2 string, r rune) {
	if !t.leftExplorer.IsDir(s) {
		return
	}

	t.leftPane.Clear()

	t.leftExplorer.Cd(s)

	t.leftPane.SetTitle(t.leftExplorer.Pwd())

	list := t.leftExplorer.Ls("")
	for _, item := range list {
		t.leftPane.AddItem(item, "", 0, nil)
	}
}

func (t *TUI) leftPaneChanged(index int, mainText string, secondaryText string, shortcut rune) {
	t.footer.SetText(t.leftExplorer.FileInfoString(mainText))
}
