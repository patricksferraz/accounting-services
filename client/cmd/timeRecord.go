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
	"time"

	appgrpc "github.com/c4ut/accounting-services/client/application/grpc"
	"github.com/c4ut/accounting-services/client/domain/service"
	"github.com/c4ut/accounting-services/client/infrastructure/external"
	"github.com/c4ut/accounting-services/service/common/pb"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var id string
var _time string
var description string
var fromDate string
var toDate string
var refusedReason string

func wraper(
	f func(s *service.TimeRecordService, ctx context.Context, vals ...interface{}) (interface{}, error),
	vals ...interface{},
) {
	authConn, err := external.ConnectService(os.Getenv("AUTH_SERVICE_ADDR"), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer authConn.Close()
	authGrpcService := pb.NewAuthServiceClient(authConn)
	authService := service.NewAuthService(authGrpcService)

	transportOption := grpc.WithInsecure()
	interceptor, err := appgrpc.NewAuthInterceptor(authService, username, password)
	if err != nil {
		log.Fatal(err)
	}

	trConn, err := external.ConnectService(
		os.Getenv("TIME_RECORD_SERVICE_ADDR"),
		transportOption,
		grpc.WithUnaryInterceptor(interceptor.Unary()),
		grpc.WithStreamInterceptor(interceptor.Stream()),
	)
	if err != nil {
		log.Fatal(err)
	}

	defer trConn.Close()
	trGrpcService := pb.NewTimeRecordServiceClient(trConn)
	trService := service.NewTimeRecordService(trGrpcService)

	res, err := f(trService, context.Background(), vals...)
	if err != nil {
		log.Fatal(err)
	}

	r, _ := json.MarshalIndent(res, "", "  ")
	log.Println(string(r))
}

// timeRecordRegisterCmd represents the timeRecordRegisterCmd command
var timeRecordRegisterCmd = &cobra.Command{
	Use:   "timeRecordRegister",
	Short: "Register a new time record",
	Run: func(cmd *cobra.Command, args []string) {

		t, err := time.Parse(time.RFC3339, _time)
		if err != nil {
			log.Fatal(err)
		}

		wraper(
			func(s *service.TimeRecordService, ctx context.Context, vals ...interface{}) (interface{}, error) {
				_time := vals[0].(time.Time)
				desc := vals[1].(string)
				res, err := s.Register(ctx, _time, desc)
				return res, err
			},
			t, description,
		)
	},
}

// timeRecordApproveCmd represents the timeRecordApproveCmd command
var timeRecordApproveCmd = &cobra.Command{
	Use:   "timeRecordApprove",
	Short: "Approve a time record",
	Run: func(cmd *cobra.Command, args []string) {

		wraper(
			func(s *service.TimeRecordService, ctx context.Context, vals ...interface{}) (interface{}, error) {
				id := vals[0].(string)
				res, err := s.Approve(ctx, id)
				return res, err
			},
			id,
		)
	},
}

// timeRecordRefuseCmd represents the timeRecordRefuseCmd command
var timeRecordRefuseCmd = &cobra.Command{
	Use:   "timeRecordRefuse",
	Short: "Refuse a time record",
	Run: func(cmd *cobra.Command, args []string) {

		wraper(
			func(s *service.TimeRecordService, ctx context.Context, vals ...interface{}) (interface{}, error) {
				id := vals[0].(string)
				refusedReason := vals[1].(string)
				res, err := s.Refuse(ctx, id, refusedReason)
				return res, err
			},
			id, refusedReason,
		)
	},
}

// timeRecordFindCmd represents the timeRecordFindCmd command
var timeRecordFindCmd = &cobra.Command{
	Use:   "timeRecordFind",
	Short: "Find a time record",
	Run: func(cmd *cobra.Command, args []string) {

		wraper(
			func(s *service.TimeRecordService, ctx context.Context, vals ...interface{}) (interface{}, error) {
				id := vals[0].(string)
				res, err := s.Find(ctx, id)
				return res, err
			},
			id,
		)
	},
}

// timeRecordSearchCmd represents the timeRecordSearchCmd command
var timeRecordSearchCmd = &cobra.Command{
	Use:   "timeRecordSearch",
	Short: "Find all time records by employee id",
	Run: func(cmd *cobra.Command, args []string) {

		fromDate, err := time.Parse(time.RFC3339, fromDate)
		if err != nil {
			log.Fatal(err)
		}

		toDate, err := time.Parse(time.RFC3339, toDate)
		if err != nil {
			log.Fatal(err)
		}

		wraper(
			func(s *service.TimeRecordService, ctx context.Context, vals ...interface{}) (interface{}, error) {
				id := vals[0].(string)
				fromDate := vals[1].(time.Time)
				toDate := vals[2].(time.Time)
				res, err := s.SearchTimeRecords(ctx, id, fromDate, toDate)
				return res, err
			},
			id, fromDate, toDate,
		)
	},
}

// timeRecordListCmd represents the timeRecordListCmd command
var timeRecordListCmd = &cobra.Command{
	Use:   "timeRecordList",
	Short: "Find all time records by employee id",
	Run: func(cmd *cobra.Command, args []string) {

		fromDate, err := time.Parse(time.RFC3339, fromDate)
		if err != nil {
			log.Fatal(err)
		}

		toDate, err := time.Parse(time.RFC3339, toDate)
		if err != nil {
			log.Fatal(err)
		}

		wraper(
			func(s *service.TimeRecordService, ctx context.Context, vals ...interface{}) (interface{}, error) {
				fromDate := vals[0].(time.Time)
				toDate := vals[1].(time.Time)
				res, err := s.ListTimeRecords(ctx, fromDate, toDate)
				return res, err
			},
			fromDate, toDate,
		)
	},
}

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/../../.env")

	if err != nil {
		log.Printf("Error loading .env files")
	}

	rootCmd.AddCommand(timeRecordFindCmd)
	rootCmd.AddCommand(timeRecordListCmd)
	rootCmd.AddCommand(timeRecordSearchCmd)
	rootCmd.AddCommand(timeRecordRefuseCmd)
	rootCmd.AddCommand(timeRecordApproveCmd)
	rootCmd.AddCommand(timeRecordRegisterCmd)

	defaultUser := os.Getenv("AUTH_SERVICE_USERNAME")
	defaultPass := os.Getenv("AUTH_SERVICE_PASSWORD")

	timeRecordFindCmd.Flags().StringVarP(&id, "id", "i", "", "Time record id to find")
	timeRecordFindCmd.Flags().StringVarP(&username, "username", "u", defaultUser, "Username for login in Auth Service")
	timeRecordFindCmd.Flags().StringVarP(&password, "password", "p", defaultPass, "Password for login in Auth Service")
	timeRecordFindCmd.MarkFlagRequired("id")

	timeRecordListCmd.Flags().StringVarP(&fromDate, "from", "f", time.Now().Format(time.RFC3339), "Date ('from') for search all time records")
	timeRecordListCmd.Flags().StringVarP(&toDate, "to", "t", time.Now().Format(time.RFC3339), "Date ('to') for search all time records")
	timeRecordListCmd.Flags().StringVarP(&username, "username", "u", defaultUser, "Username for login in Auth Service")
	timeRecordListCmd.Flags().StringVarP(&password, "password", "p", defaultPass, "Password for login in Auth Service")

	timeRecordSearchCmd.Flags().StringVarP(&id, "employeeID", "i", "", "Employee id to search")
	timeRecordSearchCmd.Flags().StringVarP(&fromDate, "from", "f", time.Now().Format(time.RFC3339), "Date ('from') for search all time records")
	timeRecordSearchCmd.Flags().StringVarP(&toDate, "to", "t", time.Now().Format(time.RFC3339), "Date ('to') for search all time records")
	timeRecordSearchCmd.Flags().StringVarP(&username, "username", "u", defaultUser, "Username for login in Auth Service")
	timeRecordSearchCmd.Flags().StringVarP(&password, "password", "p", defaultPass, "Password for login in Auth Service")
	timeRecordSearchCmd.MarkFlagRequired("employeeID")

	timeRecordRefuseCmd.Flags().StringVarP(&id, "id", "i", "", "Time record id to approve")
	timeRecordRefuseCmd.Flags().StringVarP(&refusedReason, "refusedReason", "r", "", "Reason for refusal")
	timeRecordRefuseCmd.Flags().StringVarP(&username, "username", "u", defaultUser, "Username for login in Auth Service")
	timeRecordRefuseCmd.Flags().StringVarP(&password, "password", "p", defaultPass, "Password for login in Auth Service")
	timeRecordRefuseCmd.MarkFlagRequired("id")
	timeRecordRefuseCmd.MarkFlagRequired("refusedReason")

	timeRecordApproveCmd.Flags().StringVarP(&id, "id", "i", "", "Time record id to approve")
	timeRecordApproveCmd.Flags().StringVarP(&username, "username", "u", defaultUser, "Username for login in Auth Service")
	timeRecordApproveCmd.Flags().StringVarP(&password, "password", "p", defaultPass, "Password for login in Auth Service")
	timeRecordApproveCmd.MarkFlagRequired("id")

	timeRecordRegisterCmd.Flags().StringVarP(&_time, "time", "t", "", "Time record time")
	timeRecordRegisterCmd.Flags().StringVarP(&description, "description", "d", "", "Time record description")
	timeRecordRegisterCmd.Flags().StringVarP(&username, "username", "u", defaultUser, "Username for login in Auth Service")
	timeRecordRegisterCmd.Flags().StringVarP(&password, "password", "p", defaultPass, "Password for login in Auth Service")
	timeRecordRegisterCmd.MarkFlagRequired("time")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timeRecordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timeRecordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
