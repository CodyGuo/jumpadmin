package cmd

import (
	"context"

	"github.com/CodyGuo/glog"
	"github.com/CodyGuo/jumpadmin/global"
	"github.com/CodyGuo/jumpadmin/internal/model"
	"github.com/CodyGuo/jumpadmin/pkg/setting"
	"github.com/spf13/cobra"
)

var cfg string
var ctx = context.Background()

var rootCmd = &cobra.Command{
	Short: "jumpadmin tools",
	Long:  "jumpadmin tools",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfg, "config", "c", "/etc/jumpadmin/config.yaml", "jumpadmin config file")
	rootCmd.AddCommand(orgCmd)
}

func initConfig() {
	set, err := setting.NewSetting(cfg)
	if err != nil {
		glog.Fatal(err)
	}
	err = set.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		glog.Fatal(err)
	}
	err = setupDBEngine()
	if err != nil {
		glog.Fatalf("init.setupDBEngine error: %v", err)
	}
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}
