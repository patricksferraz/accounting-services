/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"log"

	"github.com/patricksferraz/accounting-services/service/time-record/infrastructure/db"
	_ "github.com/patricksferraz/accounting-services/service/time-record/infrastructure/db/migrations"
	migrate "github.com/patricksferraz/mongo-migrate"
	"github.com/spf13/cobra"
)

var n int

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate [up|down]",
	Short: "A brief description of your command",
	Args: func(cmd *cobra.Command, args []string) error {

		if len(args) < 1 {
			return errors.New("requires up or down argument")
		}
		return nil

	},
	Run: func(cmd *cobra.Command, args []string) {

		db, err := db.ConnectMongoDB()
		if err != nil {
			log.Fatal(err)
		}
		migrate.SetDatabase(db)

		switch args[0] {
		case "up":
			err = migrate.Up(n)
			if err != nil {
				log.Fatal(err)
			}
		case "down":
			err = migrate.Down(n)
			if err != nil {
				log.Fatal(err)
			}
		default:
			log.Fatal("requires up or down argument")
		}

	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().IntVarP(&n, "n", "n", migrate.AllAvailable, "amount of migrations to UP or DOWN")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
