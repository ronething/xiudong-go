package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"xiudong/showstart"
)

var addressCmd = &cobra.Command{
	Use:   "address",
	Short: "查询个人地址",
	Long:  `查询个人地址 如果没有需要自行到秀动 app 进行补充`,
	Run: func(cmd *cobra.Command, args []string) {
		//参数校验
		if globalConfig == nil {
			fmt.Println("nil config")
			return
		}
		s := showstart.NewShowStart(globalConfig, globalClient)
		addr, err := s.GetAddress()
		if err != nil {
			log.Println("获取地址失败:", err)
			return
		}

		log.Printf("地址: %+v\n", addr)
	},
}

func init() {
	rootCmd.AddCommand(addressCmd)
}
