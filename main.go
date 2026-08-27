package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/Pastalikek65/dotman/store"
	"github.com/Pastalikek65/dotman/tui"
)

var version = "0.1.0"

func main() {
	root := &cobra.Command{Use: "dotman", Short: "Dotfiles manager for Termux — git sync, TUI", Version: version}
	root.AddCommand(&cobra.Command{
		Use: "add [file]", Args: cobra.ExactArgs(1), Short: "add dotfile",
		RunE: func(cmd *cobra.Command, args []string) error {
			dst, err := store.Add(args[0])
			if err != nil { return err }
			fmt.Println(dst)
			return nil
		},
	})
	root.AddCommand(&cobra.Command{
		Use: "list", Short: "list dotfiles",
		RunE: func(cmd *cobra.Command, args []string) error {
			list, _ := store.List()
			for _, f := range list { fmt.Println(f) }
			return nil
		},
	})
	root.AddCommand(&cobra.Command{
		Use: "restore [file] [dst]", Args: cobra.RangeArgs(1,2), Short: "restore dotfile",
		RunE: func(cmd *cobra.Command, args []string) error {
			dst := ""
			if len(args)==2 { dst=args[1] }
			return store.Restore(args[0], dst)
		},
	})
	root.AddCommand(&cobra.Command{
		Use: "tui", Short: "open TUI",
		RunE: func(cmd *cobra.Command, args []string) error {
			m:=tui.New()
			p:=tea.NewProgram(m, tea.WithAltScreen())
			_,err:=p.Run()
			return err
		},
	})
	if err:=root.Execute(); err!=nil { os.Exit(1) }
}
