package config

import (
	"fmt"
	"runtime"

	"github.com/spf13/viper"
)

var (
	// Offset is offset row numbers on Excel
	Offset int
	// LimitOfParallelDownload is limit of parallel downloads from ufo catcher
	LimitOfParallelDownload int
	// LimitOfParallelProcess is limit of parallel processes
	LimitOfParallelProcess int

	// UpdateCheck is update check
	UpdateCheck bool
	// RealTime is realtime flag
	RealTime bool

	// DBPath is db path
	DBPath string
	// HTMLRoot is root folder of generated html files
	HTMLRoot string
)

func init() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")    // for package testing
	viper.AddConfigPath("../..") // for package testing
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.SetDefault("offset", 2)
	viper.SetDefault("real-time", true)
	viper.SetDefault("update-check", true)

	viper.SetDefault("limit.download", 1)
	viper.SetDefault("limit.process", runtime.NumCPU())

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	Offset = viper.GetInt("offset")
	RealTime = viper.GetBool("real-time")
	UpdateCheck = viper.GetBool("update-check")

	LimitOfParallelDownload = viper.GetInt("limit.download")
	LimitOfParallelProcess = viper.GetInt("limit.process")

	DBPath = viper.GetString("path.db")
	HTMLRoot = viper.GetString("path.html")
}
