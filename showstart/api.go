package showstart

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

// API showstart server api
type API interface {
	GetAddressList(pageNo int64) ([]Address, error)
	GetAddress() (*Address, error)
	GetCpList(pageNo int64) ([]CpItem, error)
	GetTicketList(activityId string) ([]TicketListResult, error)
}

type ShowStart struct {
	Cookies *WapEncryptConfig
	Address *Address
	Client  *resty.Client
}

const OKShowstartState = "1"

func NewShowStart(cookies *WapEncryptConfig, client *resty.Client) API {
	if client == nil {
		client = resty.New()
	}
	return &ShowStart{Cookies: cookies, Client: client}
}

func (s *ShowStart) GetAddressList(pageNo int64) ([]Address, error) {
	w, err := NewWapEncrypt(s.Cookies, &Source{
		URL:    "/wap/address/list.json",
		Method: "POST",
		Data: map[string]interface{}{
			"pageNo": pageNo,
		},
	})
	if err != nil {
		return nil, err
	}

	// POST Map, default is JSON content type. No need to set one
	resp, err := s.Client.R().
		SetBody(w.Source.Data).
		SetHeaders(w.Source.Headers).
		SetResult(&AddressResp{}).
		Post(w.GetRequestUrl())
	if err != nil {
		logx.Errorf("GetAddress 请求发生错误: %v", err)
		return nil, err
	}

	v, ok := resp.Result().(*AddressResp)
	if !ok {
		logx.Errorf("AddressResp 断言失败")
		return nil, fmt.Errorf("AddressResp 断言失败")
	}

	logx.Infof("addressResp is %+v", v)
	if v.State != OKShowstartState {
		return nil, fmt.Errorf("showstart 状态码错误: %v", v.State)
	}

	return v.Result, nil
}

// GetAddress 获取用户地址列表
func (s *ShowStart) GetAddress() (*Address, error) {
	resp, err := s.GetAddressList(1)
	if err != nil {
		return nil, err
	}

	if len(resp) == 0 {
		return nil, fmt.Errorf("请补充地址列表")
	}
	s.Address = &resp[0]

	return s.Address, nil
}

// GetCpList 获取观演人列表
func (s *ShowStart) GetCpList(pageNo int64) ([]CpItem, error) {
	// 目前看来 pageNo 貌似没有特别的作用 例如传了 1 和 2 都是一样的返回结果 不知道秀动服务端如何处理的

	w, err := NewWapEncrypt(s.Cookies, &Source{
		URL:    "/wap/cp/list.json",
		Method: "POST",
		Data: map[string]interface{}{
			"pageNo": pageNo,
		},
	})
	if err != nil {
		return nil, err
	}

	// POST Map, default is JSON content type. No need to set one
	resp, err := s.Client.R().
		SetBody(w.Source.Data).
		SetHeaders(w.Source.Headers).
		SetResult(&CpResp{}).
		Post(w.GetRequestUrl())
	if err != nil {
		logx.Errorf("GetCpList 请求发生错误: %v", err)
		return nil, err
	}

	v, ok := resp.Result().(*CpResp)
	if !ok {
		logx.Errorf("CpResp 断言失败")
		return nil, fmt.Errorf("CpResp 断言失败")
	}

	logx.Infof("GetCpList is %+v\n", v)
	if v.State != OKShowstartState {
		return nil, fmt.Errorf("showstart 状态码错误: %v", v.State)
	}

	return v.Result, nil
}

// GetTicketList 获取场次票种列表
func (s *ShowStart) GetTicketList(activityId string) ([]TicketListResult, error) {
	w, err := NewWapEncrypt(s.Cookies, &Source{
		URL:    "/wap/activity/V2/ticket/list",
		Method: "POST",
		Data: map[string]interface{}{
			"activityId": activityId,
			"coupon":     "", // 优惠卷 默认先为空吧
		},
	})
	if err != nil {
		return nil, err
	}

	// POST Map, default is JSON content type. No need to set one
	resp, err := s.Client.R().
		SetBody(w.Source.Data).
		SetHeaders(w.Source.Headers).
		SetResult(&TicketResp{}).
		Post(w.GetRequestUrl())
	if err != nil {
		logx.Errorf("GetTicketList 请求发生错误: %v", err)
		return nil, err
	}

	v, ok := resp.Result().(*TicketResp)
	if !ok {
		logx.Errorf("TicketResp 断言失败")
		return nil, fmt.Errorf("TicketResp 断言失败")
	}

	//logx.Infof("GetTicketList is %+v\n", v)
	if v.State != OKShowstartState {
		logx.Errorf("获取票种列表状态发生错误: %v", v.State) // 可以反应出是否是帐号的凭证过期之类
		return nil, fmt.Errorf("showstart 状态码错误: %v", v.State)
	}

	if len(v.Result) == 0 {
		return nil, fmt.Errorf("没有获取到对应场次票种列表")
	}

	return v.Result, nil
}
