package showstart

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

var errState = fmt.Errorf("state != 1")

func checkState(state interface{}) error {
	switch state.(type) {
	case float64: // 变成了 float64
		if state == float64(1) {
			return nil
		}
	case string:
		if state == "1" {
			return nil
		}
	}

	logx.Errorf("err state is : %v", state)
	// 其他情况都是返回错误
	return errState
}

func getState(state interface{}) string {
	return fmt.Sprintf("%v", state)
}
