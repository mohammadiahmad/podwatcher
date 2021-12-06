package main

import (
	"fmt"
	cmd "github.com/mohammadiahmad/podwatcher/cmd/watcher"

	"github.com/spf13/cobra"
)

const (
	short = "short description"
	long  = `long description`
)

func main() {
	root := &cobra.Command{Short: short, Long: long}
	root.AddCommand(cmd.Watcher())

	if err := root.Execute(); err != nil {
		fmt.Println(err.Error())
	}
}
