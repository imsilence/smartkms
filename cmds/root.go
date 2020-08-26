package cmds

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "smartkms",
	Short: "Smark KMS",
	Long:  "Smark KMS",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

// Execute 程序执行入口
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
