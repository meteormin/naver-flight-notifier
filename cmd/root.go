/*
Copyright © 2025 meteormin <miniyu97@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "naver-flight-notifier",
	Short: "네이버 항공 알림이",
	Long:  `가성비 일본 여행을 위해 디스코드로 원하는 조건의 항공권 알림을 받아볼 수 있습니다.`,
	RunE: func(cmd *cobra.Command, args []string) error {
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	var cfgFile string
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "./config.toml", "config file (default is ./config.toml)")
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
