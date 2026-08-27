package store

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func RepoDir() string {
	if xdg := os.Getenv("XDG_CACHE_HOME"); xdg != "" {
		return filepath.Join(xdg, "dotman", "repo")
	}
	home, _ := os.UserHomeDir()
	if home == "" { home = os.Getenv("HOME") }
	if home == "" { home = "/tmp" }
	return filepath.Join(home, ".cache", "dotman", "repo")
}

func EnsureRepo() error {
	dir := RepoDir()
	if err := os.MkdirAll(dir, 0700); err != nil { return err }
	if _, err := os.Stat(filepath.Join(dir, ".git")); err != nil {
		cmd := exec.Command("git", "init")
		cmd.Dir = dir
		if err := cmd.Run(); err != nil { return err }
	}
	return nil
}

func Add(src string) (string, error) {
	if err := EnsureRepo(); err != nil { return "", err }
	if _, err := os.Stat(src); err != nil { return "", err }
	// copy file to repo, preserve basename
	base := filepath.Base(src)
	dst := filepath.Join(RepoDir(), base)
	data, err := os.ReadFile(src)
	if err != nil { return "", err }
	if err := os.WriteFile(dst, data, 0600); err != nil { return "", err }
	// git add + commit
	cmd := exec.Command("git", "add", base)
	cmd.Dir = RepoDir()
	cmd.Run()
	cmd = exec.Command("git", "-c", "user.name=dotman", "-c", "user.email=dotman@local", "commit", "-m", "add "+base)
	cmd.Dir = RepoDir()
	// ignore commit error if nothing to commit
	cmd.Run()
	return dst, nil
}

func List() ([]string, error) {
	if err := EnsureRepo(); err != nil { return nil, err }
	ents, err := os.ReadDir(RepoDir())
	if err != nil { return nil, err }
	var out []string
	for _, e := range ents {
		if e.Name()==".git" { continue }
		out = append(out, e.Name())
	}
	return out, nil
}

func Restore(name, dst string) error {
	src := filepath.Join(RepoDir(), name)
	if _, err := os.Stat(src); err != nil { return fmt.Errorf("not found %s", name) }
	if dst == "" {
		// try to restore to original location: if name is termux.properties, restore to ~/.termux/
		home, _ := os.UserHomeDir()
		if strings.Contains(name, "termux") {
			dst = filepath.Join(home, ".termux", name)
		} else {
			dst = filepath.Join(home, name)
		}
	}
	if err := os.MkdirAll(filepath.Dir(dst), 0700); err != nil { return err }
	data, err := os.ReadFile(src)
	if err != nil { return err }
	return os.WriteFile(dst, data, 0600)
}
