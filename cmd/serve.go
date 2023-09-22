/*
Copyright © 2023 xujiahao <1787619881@qq.com>
*/
package cmd

import (
	"fmt"
	"os"
	"reflect"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xujiahao718/go-easy/common/config"
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
			initConfig(cmd)
		},
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringVar(&cnfFile, "config", "", "config file path, default config.yaml in \".\" and \"./configs\"")
	m := config.GetFlagMap()
	usage := "overwrite config file"
	for k, v := range m {
		switch v.Type().Kind() {
		case reflect.Bool:
			serveCmd.Flags().Bool(k, v.Bool(), usage)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			serveCmd.Flags().Int64(k, v.Int(), usage)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			serveCmd.Flags().Uint64(k, v.Uint(), usage)
		case reflect.Float32, reflect.Float64:
			serveCmd.Flags().Float64(k, v.Float(), usage)
		case reflect.String:
			serveCmd.Flags().String(k, v.String(), usage)
		case reflect.Slice:
			et := v.Type().Elem()
			switch et.Kind() {
			case reflect.Bool:
				serveCmd.Flags().BoolSlice(k, []bool{}, usage)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
				reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				serveCmd.Flags().Int64Slice(k, []int64{}, usage)
			case reflect.Float32, reflect.Float64:
				serveCmd.Flags().Float64Slice(k, []float64{}, usage)
			case reflect.String:
				serveCmd.Flags().StringSlice(k, []string{}, usage)
			}
		}
	}
}

func initConfig(cmd *cobra.Command) {
	viper.Set("isok", false)
	viper.BindPFlags(cmd.Flags())

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

	viper.Unmarshal(&config.Configs)
	// when config file changed, should unmarshal again
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed, update.")
		viper.Unmarshal(&config.Configs)
	})
}

func run() {
	fmt.Println("server is running")

	fmt.Println("welcome to ", config.Configs.Application.Name)
	fmt.Println("version: ", config.Configs.Application.Version)
}
