package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/robopuff/go-workshop/kata/02_pubsub/internal/client"
)

var (
	app            *tview.Application
	consumersFlex  *tview.Flex
	loggerTextView *tview.TextView

	consumerUrl, producerUrl *string
	consumers                []*tview.TextView
)

// This is just demo subscriber, most of the logic
// in this function should be moved to a better place
// but due to limited amount of time that I had, it'll
// have to work ;)
func main() {
	var err error

	producerUrl = flag.String("producer", "ws://localhost:8080/ws", "Service that produces messages")
	consumerUrl = flag.String("consumer", "ws://localhost:8090/ws", "Service that consume messages")
	flag.Parse()

	if *consumerUrl == "" || *producerUrl == "" {
		fmt.Println("Consumer and producer have to be provided and cannot be empty")
		return
	}

	if err != nil {
		panic(err)
	}

	app = tview.NewApplication()
	flex := tview.NewFlex().SetDirection(tview.FlexRow)

	topFlex := tview.NewFlex().SetDirection(tview.FlexColumn)
	topFlex.SetBorder(false)

	inputField := tview.NewInputField().
		SetLabel("Message ").
		SetPlaceholder("Write anything you want to send to producer").
		SetFieldBackgroundColor(tcell.ColorBlack)

	app.SetRoot(flex, true).SetFocus(inputField)

	inputField.
		SetBorder(false)

	flex.AddItem(topFlex, 0, 1, false)
	flex.AddItem(inputField, 1, 1, false)

	loggerTextView = tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() { app.Draw() })

	loggerTextView.
		SetBorder(true).
		SetTitle("Logger").
		SetTitleColor(tcell.ColorBlue).
		SetTitleAlign(tview.AlignLeft)

	consumersFlex = tview.NewFlex().SetDirection(tview.FlexRow)
	newConsumer()

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlN:
			newConsumer()
		case tcell.KeyCtrlQ:
			if len(consumers) == 1 {
				return event
			}

			c := consumers[len(consumers)-1]
			consumers = consumers[:len(consumers)-1]
			consumersFlex.RemoveItem(c)
		}
		return event
	})

	topFlex.
		AddItem(loggerTextView, 0, 1, false).
		AddItem(consumersFlex, 0, 1, false)

	log("Welcome to demo client app")
	log("------------")
	log("[blue][red]CTRL+N[blue] to create new consumer[white]")
	log("[blue][red]CTRL+Q[blue] to close consumer[white]")
	log("[blue][red]CTRL+C[blue] to quit[white]")
	log("------------")

	p := client.NewWebsocketsClient(*producerUrl)
	defer p.Close()

	inputField.SetDoneFunc(func(key tcell.Key) {
		if key != tcell.KeyEnter {
			return
		}

		m := inputField.GetText()
		inputField.SetText("")

		err := p.Send(m)
		if err != nil {
			logError(err)
			return
		}
		log(fmt.Sprintf("[yellow]publisher [green]WRITE[white] `%s`", m))
	})

	if err := app.Run(); err != nil {
		panic(err)
	}
}

func newConsumer() {
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() { app.Draw() })

	textView.
		SetBorder(true).
		SetTitle("Consumer").
		SetTitleColor(tcell.ColorGreen).
		SetTitleAlign(tview.AlignLeft)

	consumers = append(consumers, textView)
	go func() {
		consumersFlex.AddItem(textView, 0, 1, false)
		addText(textView, fmt.Sprintf("Consumer starting ..."))

		c := client.NewWebsocketsClient(*consumerUrl)
		defer c.Close()
		for {
			m, err := c.Read()
			if err != nil {
				logError(err)
				continue
			}

			addText(textView, fmt.Sprintf("[green]READ[white] `%s`", m))
		}
	}()
}

func addText(v *tview.TextView, message string) {
	v.SetText(
		fmt.Sprintf(
			"%s[blue][%s][white] %s",
			v.GetText(false),
			time.Now().Format("15:04:05"),
			message,
		),
	)
}

func log(message string) {
	addText(loggerTextView, message)
}

func logError(err error) {
	log(fmt.Sprintf("[red]ERROR[white] %s", err.Error()))
}
