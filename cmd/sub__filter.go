// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/sky0621/repertorium/service"
	"github.com/sky0621/repertorium/static"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// filterCmd represents the filter command
var filterCmd = &cobra.Command{
	Use:   "filter",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger, _ := zap.NewProduction()
		defer logger.Sync()
		logger.Info("filter called")

		name := viper.GetString("get.filter.name")
		language := viper.GetString("get.filter.language")

		listupOutputPath, err := cmd.PersistentFlags().GetString(static.FlagKeyListupOutputPath)
		if err != nil {
			logger.Error("@PersistentFlags.Get", zap.String("flag.key", static.FlagKeyListupOutputPath))
			return
		}

		filterOutputPath, err := cmd.PersistentFlags().GetString(static.FlagKeyFilterOutputPath)
		if err != nil {
			logger.Error("@PersistentFlags.Get", zap.String("flag.key", static.FlagKeyFilterOutputPath))
			return
		}

		logger.Info("[settings]", zap.String("name", name), zap.String("language", language), zap.String("listupOutputPath", listupOutputPath), zap.String("filterOutputPath", filterOutputPath))

		err = service.Filter(name, language, listupOutputPath, filterOutputPath)
		if err != nil {
			logger.Error("@service.Filter", zap.Error(err))
			return
		}
	},
}

func init() {
	getCmd.AddCommand(filterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// filterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// filterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
