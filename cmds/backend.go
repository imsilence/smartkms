package cmds

import (
	"github.com/imsilence/smartkms/backend"
	"github.com/imsilence/smartkms/backend/migrations"
	"github.com/imsilence/smartkms/backend/routers"
	"github.com/spf13/cobra"
)

var backendCmd = &cobra.Command{
	Use:   "backend",
	Short: "Smark KMS Backend Server",
	Long:  "Smark KMS Backend Server",
	RunE: func(cmd *cobra.Command, args []string) error {
		app, err := backend.InitApp()
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
	rootCmd.AddCommand(backendCmd)
}
