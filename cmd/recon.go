/*
Copyright © 2026 JayyDnt
*/
package cmd

import (
	"context"
	"fmt"
	"hz/internal"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

var reconCmd = &cobra.Command{
	Use:   "recon",
	Short: "Subdomain recon wraper with go",
	Run: func(cmd *cobra.Command, args []string) {
		if err := args; err != nil {
			fmt.Print("Missing input :(")
			return
		}

		internal.RunRecon(args[0])
	},
}

func init() {
	rootCmd.AddCommand(reconCmd)
}

func Execute() {
	ctx, camcel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGALRM)
	defer camcel()

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
