package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	DbStringKey = "db_string"
)

func init() {
	rootCmd.PersistentFlags().StringP(DbStringKey, "d", "", "Provide postgres connection string")
	viper.BindPFlag(DbStringKey, rootCmd.PersistentFlags().Lookup(DbStringKey))

	rootCmd.AddCommand(csvCmd)
}

var rootCmd = &cobra.Command{
	Use:   "import",
	Short: "imports meter logs to the database",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() error {
	return rootCmd.Execute()
}
