package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var promethuesURL string
var rootCmd = &cobra.Command{
	Use:   "cpam",
	Short: "CPAM is a CLI tool for HPC cluster monitoring",
	Long:  `A high-performance monitoring tool that fetches metrics from Prometheus for Nodes, Kafka, and OpenSearch.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&promethuesURL, "url", "u", "http://localhost:9090", "Prometheus URL")
}
