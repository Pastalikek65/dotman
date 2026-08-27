# Contributing to dotman

```bash
git clone https://github.com/Pastalikek65/dotman.git
cd dotman
go vet ./...
go test ./... -count=1 -timeout 30s -cover
CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o dotman .
./dotman add ~/.bashrc && ./dotman list
```

PRs: fork → feat/name → test → open PR.
