/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/konan0802/dakoku/handler"
	"github.com/konan0802/dakoku/infra"
	"github.com/konan0802/dakoku/usecase"

	"github.com/spf13/cobra"
)

// tasksCmd represents the tasks command
var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "固定タスクとAsanaタスクの一覧を表示",
	Long:  "固定タスクとAsanaタスクの一覧を表示",
	Run: func(cmd *cobra.Command, args []string) {
		// DI
		arp := infra.NewAsanaRepository()
		trp := infra.NewTogglRepository()
		uc := usecase.NewUsecase(arp, trp)
		hdr := handler.NewHandler(uc)

		// handler
		ret, err := hdr.Tasks(cmd, args)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		fmt.Println(ret)
	},
}

func init() {
	rootCmd.AddCommand(tasksCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tasksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tasksCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
