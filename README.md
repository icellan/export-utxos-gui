# Export UTXOs using WhatsOnChain.com

This is a GUI application built with the Fyne framework in Go. The application allows users to upload a file containing Bitcoin addresses and processes these addresses to export UTXOs (Unspent Transaction Outputs) using the WhatsOnChain.com API.

## Features

- Select a file containing Bitcoin addresses.
- Process the addresses to retrieve UTXOs.
- Display a progress bar during processing.
- Save the output to a JSON file.

## Prerequisites

- Go 1.16 or later
- Fyne 2.0 or later

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/icellan/export-utxos.git
    cd export-utxos
    ```

2. Install the dependencies:

    ```sh
    go mod tidy
    ```

## Usage

1. Run the application:

    ```sh
    go run main.go
    ```

2. The application window will open. Click the "Click here to select a file with addresses" button to open a file picker dialog.

3. Select a file containing Bitcoin addresses, one per line.

4. The application will process the addresses and display a progress bar.

5. Once processing is complete, a file save dialog will appear. Choose a location to save the output JSON file.

## File Format

The input file should contain one Bitcoin address per line, for example:

```
1Pwyh2BK4ez9uyDMiq5bDXdHdLHVFa13Xt
1Nvazaharko2cPtVndBVnsRt73ATgyydZF
1AdZzMusaum92gHEcTDdDv7Lt733yii6EN
156xqGjcFVaQMAZBcdmM2fy7SHyNrZPpbM
```

## Project Structure

- `main.go`: The main entry point of the application.
- `models/`: Contains the data models used in the application.
- `process/`: Contains the logic for processing the addresses and retrieving UTXOs.
- `.gitignore`: Specifies files and directories to be ignored by Git.

## Dependencies

- [Fyne](https://fyne.io/) - A cross-platform GUI toolkit for Go.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
