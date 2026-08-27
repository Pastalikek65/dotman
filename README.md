# dotman

[![Go 1.25](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go)](https://go.dev)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Termux](https://img.shields.io/badge/Termux-friendly-brightgreen?logo=android)](https://termux.dev)

Dotfiles manager for Termux — git sync, TUI, 5M.

I built `dotman` because `provision` installs packages but doesn't version your `~/.termux/termux.properties` and `~/.bashrc`. This does one thing: `add` → `list` → `restore`, with a `tui` on top.

```bash
go install github.com/Pastalikek65/dotman@latest
dotman add ~/.termux/termux.properties
dotman list
dotman tui   # j/k nav, q quit
dotman restore termux.properties ~/.termux/termux.properties
```

Repo at `~/.cache/dotman/repo` (`0700`, git), survives reboot.

## Install (Termux)

```bash
git clone https://github.com/Pastalikek65/dotman.git
cd dotman
CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o dotman .
./dotman add ~/.bashrc
```

## How it works

`store` → `RepoDir()` `XDG_CACHE_HOME` or `~/.cache/dotman/repo` `0700`, `git init` if needed, `Add` copies file + `git add/commit`, `List` reads dir, `Restore` copies back. No daemon.

## Dev

```bash
go vet ./...
go test ./... -count=1 -timeout 30s -cover
CGO_ENABLED=0 go build -o dotman .
```

## License

MIT
