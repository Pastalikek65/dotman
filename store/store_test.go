package store

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAddAndList(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("XDG_CACHE_HOME", dir)
	src := filepath.Join(dir, "test.conf")
	os.WriteFile(src, []byte("hello"), 0600)
	dst, err := Add(src)
	if err != nil { t.Fatal(err) }
	if _, err := os.Stat(dst); err != nil { t.Fatal(err) }
	list, _ := List()
	if len(list) != 1 || list[0] != "test.conf" { t.Fatalf("list %v", list) }
}

func TestRestore(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("XDG_CACHE_HOME", dir)
	src := filepath.Join(dir, "a.conf")
	os.WriteFile(src, []byte("data"), 0600)
	Add(src)
	dst := filepath.Join(dir, "restore.conf")
	if err := Restore("a.conf", dst); err != nil { t.Fatal(err) }
	if data, _ := os.ReadFile(dst); string(data) != "data" { t.Fatalf("restore %s", data) }
}
