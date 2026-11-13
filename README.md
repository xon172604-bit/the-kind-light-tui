# The Kind Light (terminal)

A short horror story you play in your terminal. Type a number to choose, `q` to quit.

Built for Hack Club's Terminal Craft.

## Screenshots / demo

_TODO: add a screenshot or GIF before submitting._

## Play

Download the binary for your system from the GitHub Releases page, then run it:

- **Windows:** `the-kind-light.exe`
- **macOS / Linux:** `chmod +x the-kind-light-<os>-<arch> && ./the-kind-light-<os>-<arch>`

## Build from source (all platforms)

You only need [Go](https://go.dev/dl/) 1.21 or newer. No other dependencies.

```sh
git clone https://github.com/xon172604-bit/the-kind-light-tui
cd the-kind-light-tui
go build -o the-kind-light .
./the-kind-light
```

Cross-compile every platform:

```sh
GOOS=windows GOARCH=amd64 go build -o dist/the-kind-light-windows.exe .
GOOS=linux   GOARCH=amd64 go build -o dist/the-kind-light-linux-amd64 .
GOOS=darwin  GOARCH=arm64 go build -o dist/the-kind-light-macos-arm64 .
GOOS=darwin  GOARCH=amd64 go build -o dist/the-kind-light-macos-amd64 .
```

## Adding scenes

The story is plain data in `main.go` (the `story` map). Add a `Scene` with text and choices; the engine needs no changes.

## License

MIT
