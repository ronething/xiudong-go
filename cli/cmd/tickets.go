package cmd

import (
	"fmt"
	"log"

	"github.com/modood/table"
	"github.com/spf13/cobra"

	"xiudong/showstart"
)

type TicketItem struct {
	TicketId     string `table:"票种标识"`
	TicketType   string `table:"票种"`
	SellingPrice string `table:"售价"`
}

var ticketsCmd = &cobra.Command{
	Use:   "tickets",
	Short: "列出指定场次 ticketId 列表",
	Long:  `根据秀动场次 id(activityId) 列出指定场次 ticketId 列表`,
	Run: func(cmd *cobra.Command, args []string) {
		//参数校验
		if globalConfig == nil {
			fmt.Println("nil config")
			return
		}
		activityId, err := cmd.Flags().GetString("activityId")
		if err != nil {
			log.Println("获取 activityId 发生错误,", err)
			return
		}
		if activityId == "" { // 为空则 return
			log.Println("activityId 场次 id 为空，请检查")
			return
		}

		s := showstart.NewShowStart(globalConfig, globalClient)
		got, err := s.GetTicketList(activityId)
		if err != nil {
			log.Println("获取观演人列表失败:", err)
			return
		}

		for _, item := range got {
			fmt.Printf("演出名称: %+v\n", item.SessionName)
			fmt.Printf("场次 ID: %+v\n", activityId)
			res := make([]TicketItem, len(item.TicketList))
			for index, ticket := range item.TicketList {
				res[index] = TicketItem{
					TicketId:     ticket.TicketId,
					TicketType:   ticket.TicketType,
					SellingPrice: ticket.SellingPrice,
				}
			}
			table.Output(res)
		}
	},
}

func init() {
	rootCmd.AddCommand(ticketsCmd)

	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	ticketsCmd.Flags().StringP("activityId", "a", "", "秀动场次 id")
}
