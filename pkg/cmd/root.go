// Package cmd implements the command line interface for starting up 'jytte'
package cmd

import (
	"fmt"
	"os"

	"github.com/neticdk/jytte/pkg/server"
	"github.com/spf13/cobra"
)

var (
	addr    string
	rootCmd = &cobra.Command{
		Use:   "jytte",
		Short: "jytte is a small http based application written for demo and testing purposes",
		Run: func(cmd *cobra.Command, args []string) {
			server.ListenAndServe(addr)
		},
	}
)

// Execute root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&addr, "address", "a", ":8080", "Listen address")
}
