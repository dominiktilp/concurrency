package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("concurrency")
}

func main() {
	rootCmd := &cobra.Command{Use: "concurrency"}
	rootCmd.AddCommand(NewRunCommand())
	rootCmd.Execute()
}
