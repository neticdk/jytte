// Package cmd implements the command line interface for starting up 'jytte'
package cmd

import (
	"fmt"
	"os"

	"github.com/neticdk/jytte/pkg/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	addr        string
	tracingAddr string
	tracing     bool

	rootCmd = &cobra.Command{
		Use:   "jytte",
		Short: "jytte is a small http based application written for demo and testing purposes",
		Run: func(cmd *cobra.Command, args []string) {
			server.ListenAndServe()
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
	viper.SetEnvPrefix("JYTTE")
	viper.AutomaticEnv()
	viper.SetDefault("TRACING", true)
	viper.SetDefault("TRACING_ADDRESS", "localhost:4317")
	viper.SetDefault("LISTEN_ADDRESS", ":8080")

	rootCmd.Flags().StringVarP(&addr, "listen-address", "l", viper.GetString("LISTEN_ADDRESS"), "Listen address")
	rootCmd.Flags().BoolVarP(&tracing, "tracing", "t", viper.GetBool("TRACING"), "Enable tracing")
	rootCmd.Flags().StringVarP(&tracingAddr, "tracing-address", "a", viper.GetString("TRACING_ADDRESS"), "Tracing address")
	viper.BindPFlags(rootCmd.PersistentFlags())
}
