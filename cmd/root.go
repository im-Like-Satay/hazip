/*
Copyright © 2026 JayyDnt
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hz",
	Short: "Simple wraper recon tool",
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
