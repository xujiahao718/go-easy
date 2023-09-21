/*
Copyright © 2023 xujiahao <1787619881@qq.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xujiahao718/go-easy/common"
)

var (
	// cnfFile is the config file
	cnfFile string

	// serveCmd represents the serve command
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "start a web server",
		Long:  `start a web server`,
		PreRun: func(cmd *cobra.Command, args []string) {
			initConfig()
		},
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringVar(&cnfFile, "config", "", "config file path, default config.yaml in \".\" and \"./configs\"")
}

func initConfig() {
	if cnfFile != "" {
		viper.SetConfigFile(cnfFile)
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath("./configs")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}
	viper.WatchConfig()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("error when read config file: ", err)
		os.Exit(1)
	}
	fmt.Printf("using config file: %s\n", viper.ConfigFileUsed())

	viper.AutomaticEnv()

	viper.Unmarshal(&common.Configs)
	// when config file changed, should unmarshal again
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed, update.")
		viper.Unmarshal(&common.Configs)
	})
}

func run() {
	fmt.Println("server is running")

	fmt.Println(common.Configs)
}
