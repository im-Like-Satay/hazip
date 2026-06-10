/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"hz/internal"

	"github.com/spf13/cobra"
)

// reconCmd represents the recon command
var reconCmd = &cobra.Command{
	Use:   "recon",
	Short: "Subdomain recon wraper with go",
	Run: func(cmd *cobra.Command, args []string) {
		internal.RunRecon()
	},
}

func init() {
	rootCmd.AddCommand(reconCmd)
}
