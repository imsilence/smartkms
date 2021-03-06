package cmds

import (
	"github.com/imsilence/smartkms/backend"
	"github.com/imsilence/smartkms/backend/migrations"
	"github.com/imsilence/smartkms/backend/routers"
	"github.com/spf13/cobra"
)

var config string

var backendCmd = &cobra.Command{
	Use:   "backend",
	Short: "Smark KMS Backend Server",
	Long:  "Smark KMS Backend Server",
	RunE: func(cmd *cobra.Command, args []string) error {
		app, err := backend.InitAppWithConfigFile(config)
		if err != nil {
			return err
		}

		routers.Register(app.Engine)
		migrations.Migrate(app.Db)
		return app.Run()
	},
	Hidden: false,
}

func init() {
	backendCmd.Flags().StringVarP(&config, "config", "c", "etc/smartkms.yaml", "config file")
	rootCmd.AddCommand(backendCmd)
}
