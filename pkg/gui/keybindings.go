package gui

import "github.com/gdamore/tcell"

func (g *Gui) setKeybind() {
	g.sourceFilesItem.SetKeybinds(g.grobalKeybind)
}

func (g *Gui) grobalKeybind(event *tcell.EventKey) {
	switch event.Rune() {
	case 'q':
		g.application.Stop()
	case 'r':
		// TODO: Run golangci-lint against the directories marked as selected.
	}
}
