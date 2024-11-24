package main

import (
	"encoding/json"
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/icellan/export-utxos/models"
	"github.com/icellan/export-utxos/process"
)

var w fyne.Window
var fileButton *widget.Button
var fileLabel *widget.Label
var selectedFile string
var progressBar *widget.ProgressBar
var output models.Output

func main() {
	a := app.New()
	w = a.NewWindow("Export UTXOs using WhatsOnChain.com")
	w.Resize(fyne.NewSize(640, 480))

	fileButton = widget.NewButton("Click here to select a file with addresses", func() {
		showFilePicker(w)
	})

	fileLabel = widget.NewLabel("")

	progressBar = widget.NewProgressBar()
	progressBar.Hide()

	addressList := "\n```\n1GoriLLa2bdsQ8fB1CA4JNkDX88mLqXf4u\n18VWHjMt4ixHddPPbs6righWTs3Sg2QNcn\n13B99QhnFUoA4cqjrdSyJkY1ASV4MeNN6Q\n\n```"
	explanationText := widget.NewRichTextFromMarkdown("> Upload a file with 1 address per line, like this:\n" + addressList)

	w.SetContent(container.NewVBox(
		fileButton,
		fileLabel,
		progressBar,
		explanationText,
	))

	w.ShowAndRun()
}

func processNewFile() {
	fileButton.Disabled()
	defer fileButton.Enable()

	if selectedFile == "" {
		dialog.ShowError(fmt.Errorf("no file selected"), w)
		return
	}

	// read addresses from file
	f, err := os.Open(selectedFile)
	if err != nil {
		dialog.ShowError(err, w)
		return
	}

	// read addresses from file
	addresses := make([]string, 0)
	for {
		var address string
		_, err = fmt.Fscanln(f, &address)
		if err != nil {
			break
		}

		addresses = append(addresses, address)
	}

	if len(addresses) == 0 {
		dialog.ShowError(fmt.Errorf("no addresses found in file"), w)
		return
	}

	fileLabel.SetText(fmt.Sprintf("Processing file %s\n", selectedFile))

	defer func() {
		fileLabel.SetText(fmt.Sprintf("Processing done, saved to file\n"))
	}()

	// process addresses
	progressBar.SetValue(0)
	progressBar.Show()

	output, err = process.Addresses(addresses, func(idx int) {
		progressBar.SetValue(float64(idx+1) / float64(len(addresses)))
	})
	if err != nil {
		dialog.ShowError(err, w)
		return
	}

	progressBar.Hide()

	fileSaveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err != nil {
			dialog.ShowError(err, w)
			return
		}

		if writer == nil {
			fileLabel.SetText(fmt.Sprintf("No file selected for saving..."))
			return
		}

		// marshal the output to JSON and save it to the file
		outputJSON, err := json.MarshalIndent(output, "", "  ")
		if err != nil {
			dialog.ShowError(fmt.Errorf("error marshalling JSON: %v", err), w)
			return
		}

		_, err = writer.Write(outputJSON)
		if err != nil {
			dialog.ShowError(fmt.Errorf("error writing JSON to file: %v", err), w)
			return
		}

		dialog.ShowInformation("Success", "Output written to file", w)
	}, w)

	fileSaveDialog.SetFileName("output.json")
	fileSaveDialog.Show()
}

func showFilePicker(w fyne.Window) {
	fileOpenDialog := dialog.NewFileOpen(func(f fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		if f == nil {
			return
		}

		selectedFile = f.URI().Path()

		processNewFile()
	}, w)

	fileLabel.SetText(fmt.Sprintf(""))
	// set the current working directory as the default location
	cwd, err := os.Getwd()
	if err != nil {
		dialog.ShowError(err, w)
		return
	}

	listableURI, err := storage.ListerForURI(storage.NewFileURI(cwd))
	if err != nil {
		dialog.ShowError(err, w)
		return
	}
	fileOpenDialog.SetLocation(listableURI)

	fileOpenDialog.Show()
}
