package tui

import (
	"github.com/gdamore/tcell/v2"
)

const prevDirectory = "../"

func (t *TUI) appEvents(event *tcell.EventKey) *tcell.EventKey {
	if event.Rune() == 'q' {
		t.app.Stop()
	}

	return event
}

func (t *TUI) leftPaneEvents(event *tcell.EventKey) *tcell.EventKey {
	if event.Rune() == rune(tcell.KeyTab) {
		t.app.SetFocus(t.rightPane.List())
		t.footer.SetText(t.rightPane.CurrentFileInfo())
		return nil
	}

	if event.Key() == 259 { // ->
		err := t.leftPane.ChangeDirectoryIfNeed(t.leftPane.GetCurrentFile())
		if err != nil {
			t.showAlert(err.Error())
		}

		return nil
	}

	if event.Key() == 260 { // <-
		err := t.leftPane.ChangeDirectoryIfNeed(prevDirectory)
		if err != nil {
			t.showAlert(err.Error())
		}

		return nil
	}

	return event
}

func (t *TUI) leftPaneSelected(_ int, fileName string, _ string, _ rune) {
	err := t.leftPane.ChangeDirectoryIfNeed(fileName)
	if err != nil {
		t.showAlert(err.Error())
	}
}

func (t *TUI) leftPaneChanged(_ int, fileName string, _ string, _ rune) {
	t.leftPane.SetCurrentFile(fileName)
	t.footer.SetText(t.leftPane.FileInfo(fileName))
}

func (t *TUI) rightPaneEvents(event *tcell.EventKey) *tcell.EventKey {
	if event.Rune() == rune(tcell.KeyTab) {
		t.app.SetFocus(t.leftPane.List())
		t.footer.SetText(t.leftPane.CurrentFileInfo())

		return nil
	}

	if event.Key() == 259 { // ->
		err := t.rightPane.ChangeDirectoryIfNeed(t.rightPane.GetCurrentFile())
		if err != nil {
			t.showAlert(err.Error())
		}

		return nil
	}

	if event.Key() == 260 { // <-
		err := t.rightPane.ChangeDirectoryIfNeed(prevDirectory)
		if err != nil {
			t.showAlert(err.Error())
		}

		return nil
	}

	return event
}

func (t *TUI) rightPaneSelected(_ int, fileName string, _ string, _ rune) {
	err := t.rightPane.ChangeDirectoryIfNeed(fileName)
	if err != nil {
		t.showAlert(err.Error())
	}

}

func (t *TUI) rightPaneChanged(_ int, fileName string, _ string, _ rune) {
	t.rightPane.SetCurrentFile(fileName)
	t.footer.SetText(t.rightPane.FileInfo(fileName))
}
