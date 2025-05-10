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
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/patricksferraz/accounting-services/client/domain/service"
	"github.com/patricksferraz/accounting-services/client/infrastructure/external"
	"github.com/patricksferraz/accounting-services/service/common/pb"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var username string
var password string
var refreshToken string
var accessToken string

// authLoginCmd represents the auth command
var authLoginCmd = &cobra.Command{
	Use:   "authLogin",
	Short: "Login in auth service",
	Run: func(cmd *cobra.Command, args []string) {

		authConn, err := external.ConnectService(os.Getenv("AUTH_SERVICE_ADDR"), grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}

		defer authConn.Close()
		authGrpcService := pb.NewAuthServiceClient(authConn)
		authService := service.NewAuthService(authGrpcService)

		jwt, err := authService.Login(context.Background(), username, password)
		if err != nil {
			log.Fatal(err)
		}

		j, _ := json.MarshalIndent(jwt, "", "  ")
		log.Println(string(j))
	},
}

// authRefreshTokenCmd represents the auth command
var authRefreshTokenCmd = &cobra.Command{
	Use:   "authRefreshToken",
	Short: "Refresh Token in auth service",
	Run: func(cmd *cobra.Command, args []string) {

		authConn, err := external.ConnectService(os.Getenv("AUTH_SERVICE_ADDR"), grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}

		defer authConn.Close()
		authGrpcService := pb.NewAuthServiceClient(authConn)
		authService := service.NewAuthService(authGrpcService)

		jwt, err := authService.RefreshToken(context.Background(), refreshToken)
		if err != nil {
			log.Fatal(err)
		}

		j, _ := json.MarshalIndent(jwt, "", "  ")
		log.Println(string(j))
	},
}

// authEmployeeClaimsCmd represents the auth command
var authEmployeeClaimsCmd = &cobra.Command{
	Use:   "authEmployeeClaims",
	Short: "Find Employee Claims by Token in auth service",
	Run: func(cmd *cobra.Command, args []string) {

		authConn, err := external.ConnectService(os.Getenv("AUTH_SERVICE_ADDR"), grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}

		defer authConn.Close()
		authGrpcService := pb.NewAuthServiceClient(authConn)
		authService := service.NewAuthService(authGrpcService)

		employee, err := authService.FindEmployeeClaimsByToken(context.Background(), accessToken)
		if err != nil {
			log.Fatal(err)
		}

		e, _ := json.MarshalIndent(employee, "", "  ")
		log.Println(string(e))
	},
}

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/../../.env")

	if err != nil {
		log.Printf("Error loading .env files")
	}

	rootCmd.AddCommand(authLoginCmd)
	rootCmd.AddCommand(authRefreshTokenCmd)
	rootCmd.AddCommand(authEmployeeClaimsCmd)

	defaultUser := os.Getenv("AUTH_SERVICE_USERNAME")
	defaultPass := os.Getenv("AUTH_SERVICE_PASSWORD")

	authLoginCmd.Flags().StringVarP(&username, "username", "u", defaultUser, "Username for login in Auth Service")
	authLoginCmd.Flags().StringVarP(&password, "password", "p", defaultPass, "Password for login in Auth Service")

	authRefreshTokenCmd.Flags().StringVarP(&refreshToken, "token", "t", "", "Token for refresh in Auth Service")
	authRefreshTokenCmd.MarkFlagRequired("token")

	authEmployeeClaimsCmd.Flags().StringVarP(&accessToken, "token", "t", "", "Access Token for Find Employee Claims in Auth Service")
	authEmployeeClaimsCmd.MarkFlagRequired("token")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// authCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// authCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
