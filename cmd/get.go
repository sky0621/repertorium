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
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/sky0621/repertorium/static"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		listupCmd.Run(cmd, args)
		filterCmd.Run(cmd, args)
		checkoutCmd.Run(cmd, args)
	},
}

func init() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	home, err := homedir.Dir()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("current directory", zap.String("home", home))

	getCmd.PersistentFlags().String(static.FlagKeyListupOutputPath, filepath.Join(home, "listup.json"), "")
	getCmd.PersistentFlags().String(static.FlagKeyFilterOutputPath, filepath.Join(home, "filter.json"), "")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
