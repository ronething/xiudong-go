package cmd

import (
	"fmt"
	"log"

	"github.com/modood/table"
	"github.com/spf13/cobra"

	"xiudong/showstart"
)

type CpItem struct {
	Id                 int    `table:"观演人标识"`
	Name               string `table:"姓名"`
	DocumentTypeStr    string `table:"类型"`
	ShowDocumentNumber string `table:"号码"`
	IsSelf             int    `table:"isSelf"`
}

var idCardCmd = &cobra.Command{
	Use:   "idCard",
	Short: "查询已绑定观演人 id",
	Long:  `查询已经绑定的观演人 如果没有需要自行到秀动 app 进行补充`,
	Run: func(cmd *cobra.Command, args []string) {
		//参数校验
		if globalConfig == nil {
			fmt.Println("nil config")
			return
		}
		s := showstart.NewShowStart(globalConfig, globalClient)
		cpList, err := s.GetCpList(1)
		if err != nil {
			log.Println("获取观演人列表失败:", err)
			return
		}
		if len(cpList) == 0 {
			log.Println("信息为空，请补充观演人信息")
			return
		}

		res := make([]CpItem, len(cpList))
		for i := 0; i < len(cpList); i++ {
			item := cpList[i]
			res[i] = CpItem{
				Id:                 item.Id,
				Name:               item.Name,
				DocumentTypeStr:    item.DocumentTypeStr,
				ShowDocumentNumber: item.ShowDocumentNumber,
				IsSelf:             item.IsSelf,
			}
		}

		table.Output(res)

	},
}

func init() {
	rootCmd.AddCommand(idCardCmd)
}
