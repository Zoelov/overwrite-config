package main

import (
	"os"

	"github.com/apex/log"
	"github.com/spf13/viper"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	app := kingpin.New("overwrite", "overwrite json config if exist else add new")
	debug := app.Flag("debug", "enable debug mode").Bool()

	runCmd := app.Command("run", "load source and new config json, overwrite if exist else add new config").Alias("r")
	sourcePath := runCmd.Flag("src", "the source config path").PlaceHolder("old config path").Default(".").String()
	newConfigPath := runCmd.Flag("new", "the new config path").PlaceHolder("new config path").Default(".").String()
	sourceConfigName := runCmd.Flag("src-name", "the source config file name").PlaceHolder("config name").Default("default").String()
	newConfigName := runCmd.Flag("new-name", "the new config file name").PlaceHolder("config name").Default("default").String()

	cmd := kingpin.MustParse(app.Parse(os.Args[1:]))

	if *debug {
		log.SetLevel(log.DebugLevel)
	}

	switch cmd {
	case runCmd.FullCommand():
		log.Debug(*sourcePath)
		log.Debug(*newConfigPath)

		viper.AddConfigPath(*sourcePath)
		viper.SetConfigName(*sourceConfigName)
		viper.SetConfigType("json")

		if err := viper.ReadInConfig(); err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}

		viper.AddConfigPath(*newConfigPath)
		viper.SetConfigName(*newConfigName)
		if err := viper.MergeInConfig(); err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}

		if err := viper.WriteConfig(); err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}

		log.Info("write config success")
	}
}
