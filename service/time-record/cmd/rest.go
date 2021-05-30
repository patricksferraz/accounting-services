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
	"context"
	"log"
	"os"

	"github.com/c4ut/accounting-services/service/common/pb"
	"github.com/c4ut/accounting-services/service/time-record/application/rest"
	"github.com/c4ut/accounting-services/service/time-record/infrastructure/db"
	"github.com/c4ut/accounting-services/service/time-record/infrastructure/external"
	"github.com/c4ut/accounting-services/utils"
	"github.com/spf13/cobra"
)

// NewRestCmd represents the rest command
func NewRestCmd() *cobra.Command {
	var restPort int
	var uri string
	var dbName string

	var restCmd = &cobra.Command{
		Use:   "rest",
		Short: "Run rest Service",

		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			database, err := db.NewMongo(ctx, uri, dbName)
			if err != nil {
				log.Fatal(err)
			}

			err = database.Test(ctx)
			if err != nil {
				log.Fatal(err)
			}

			if utils.GetEnv("DB_MIGRATE", "false") == "true" {
				err = database.Migrate()
				if err != nil {
					log.Fatal(err)
				}
			}

			defer database.Close(ctx)

			authServiceAddr := os.Getenv("AUTH_SERVICE_ADDR")
			conn, err := external.ConnectAuthService(authServiceAddr)
			if err != nil {
				log.Fatal(err)
			}

			defer conn.Close()
			authdb := pb.NewAuthServiceClient(conn)

			rest.StartRestServer(database, authdb, restPort)
		},
	}

	dUri := utils.GetEnv("DB_URI", "mongodb://localhost")
	dDbName := utils.GetEnv("DB_NAME", "time_record_service")

	restCmd.Flags().IntVarP(&restPort, "port", "p", 8080, "rest server port")
	restCmd.Flags().StringVarP(&uri, "uri", "u", dUri, "db uri")
	restCmd.Flags().StringVarP(&dbName, "dbName", "", dDbName, "db name")

	return restCmd
}

func init() {
	rootCmd.AddCommand(NewRestCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// restCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// restCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
