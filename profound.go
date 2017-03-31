package main

import (
	ui "github.com/gizak/termui"
	"github.com/kenan-rhoton/profound/sword"
	"github.com/kenan-rhoton/profound/widgets"
	"log"
	"os"
	"strings"
)

var app struct {
	verse    *sword.V
	selected int
	text     *ui.Par
	exp      *ui.Par
	logfile  *os.File
}

func VerseMenu() {
	ui.DefaultEvtStream.ResetHandlers()
	defer SetupHandlers()

	i := widgets.NewInput()
	i.BorderLabel = "Filter"
	i.SetY(ui.TermHeight() - i.Height)
	i.Data = app.verse.Name
	ui.Render(i)

	// refresh container rows on input
	stream := i.Stream()
	go func() {
		for _ = range stream {
			ui.Render(i)
		}
	}()

	i.InputHandlers()
	ui.Handle("/sys/kbd/<escape>", func(ui.Event) {
		ui.StopLoop()
	})
	ui.Handle("/sys/kbd/<enter>", func(ui.Event) {
		var err error
		app.verse, err = sword.Verse(i.Data)
		if err != nil {
			ui.Close()
			log.Println(err)
		}
		app.selected = 0
		ui.StopLoop()
	})
	ui.Loop()
}

func SetupHandlers() {
	ui.DefaultEvtStream.ResetHandlers()
	ui.Handle("sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("sys/kbd/<escape>", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("sys/kbd/<left>", func(ui.Event) {
		app.selected--
		if app.selected < 0 {
			app.selected = 0
		}
		Display()
	})

	ui.Handle("sys/kbd/<right>", func(ui.Event) {
		app.selected++
		if app.selected >= len(app.verse.Words) {
			app.selected = len(app.verse.Words) - 1
		}
		Display()
	})

	ui.Handle("sys/kbd/f", func(ui.Event) {
		VerseMenu()
		Display()
	})
}

func Setup() {
	var err error
	app.logfile, err = os.OpenFile("pro.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		ui.Close()
		log.Fatal(err)
	}
	log.SetOutput(app.logfile)
	SetupHandlers()
	app.verse, err = sword.Verse("John 3:16")

	if err != nil {
		ui.Close()
		log.Println(err)
	}

	app.text = ui.NewPar("")
	app.exp = ui.NewPar("")
	app.selected = 0
	ui.Body.AddRows(ui.NewRow(ui.NewCol(12, 0, app.text, app.exp)))
}

func Display() {
	var displaytext []string

	for i, w := range app.verse.Words {
		if i == app.selected {
			displaytext = append(displaytext, "["+w+"](fg-red)")
		} else {
			displaytext = append(displaytext, w)
		}
	}

	app.text.Text = strings.Join(displaytext, " ")
	app.text.Height = 3
	app.text.Width = ui.TermWidth()

	app.exp.Text = app.verse.Ref[app.selected]
	app.exp.Height = ui.TermHeight() - 3
	app.exp.Width = ui.TermWidth()

	ui.Body.Align()

	ui.Render(ui.Body)

}

func main() {
	err := ui.Init()
	if err != nil {
		ui.Close()
		log.Println(err)
	}
	defer ui.Close()
	Setup()
	defer app.logfile.Close()

	Display()

	ui.Loop()
}
