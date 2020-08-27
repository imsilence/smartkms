package cmds

import (
	"fmt"

	"github.com/imsilence/smartkms/backend/utils"
	"github.com/spf13/cobra"
)

var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "Smark KMS Root Key Generate",
	Long:  "Smark KMS Root Key Generate",
	RunE: func(cmd *cobra.Command, args []string) error {
		key, err := utils.RandAesKey()
		if err != nil {
			return err
		}
		fmt.Println("key:", key)
		return nil
	},
	Hidden: false,
}

func init() {
	rootCmd.AddCommand(keyCmd)
}
