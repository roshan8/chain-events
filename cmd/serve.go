/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "API Server to collect and store events",
	Long: `Provides a set of APIs to collect events from different sources 
and store them in a database. It supports multiple event sources and ensures efficient storage and retrieval..`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
