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
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/spf13/cobra"
)

// listupCmd represents the listup command
var listupCmd = &cobra.Command{
	Use:   "listup",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger, _ := zap.NewProduction()
		defer logger.Sync()
		logger.Info("listup called")

		owner := viper.GetString("get.listup.owner")
		accessToken := viper.GetString("get.listup.accessToken")
		maxPage := viper.GetInt("get.listup.maxPage")

		listupOutputPath, err := cmd.PersistentFlags().GetString(static.FlagKeyListupOutputPath)
		if err != nil {
			logger.Error("@PersistentFlags.Get", zap.String("flag.key", static.FlagKeyListupOutputPath))
			return
		}

		logger.Info("[settings]", zap.String("owner", owner), zap.Int("maxPage", maxPage), zap.String("listupOutputPath", listupOutputPath))

		err = service.Listup(owner, accessToken, maxPage, listupOutputPath)
		if err != nil {
			logger.Error("@service.Listup", zap.Error(err))
			return
		}
	},
}

func init() {
	getCmd.AddCommand(listupCmd)
	// TODO if I wanna add other command(f.e. set) and use same listupCmd, add this.
	// setCmd.AddCommand(listupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
