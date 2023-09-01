package showstart

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
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
	Cookies *WapEncryptConfigV3
	Address *Address
	Client  *resty.Client
}

func NewShowStart(cookies *WapEncryptConfigV3, client *resty.Client) API {
	if client == nil {
		client = resty.New().SetRetryCount(1)
	}
	client.SetRetryCount(1)
	return &ShowStart{Cookies: cookies, Client: client}
}

func (s *ShowStart) GetAddressList(pageNo int64) ([]Address, error) {
	source := &SourceV3{
		URL:    "/wap/address/list.json",
		Method: "POST",
		Data: map[string]interface{}{
			"pageNo": pageNo,
		},
	}

	// POST Map, default is JSON content type. No need to set one
	var resp AddressResp
	err := s.SendRequest(context.TODO(), source, &resp)
	if err != nil {
		logx.Errorf("GetAddress 请求发生错误: %v", err)
		return nil, err
	}

	logx.Infof("addressResp is %+v", resp)

	return resp.Result, nil
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

	source := &SourceV3{
		URL:    "/wap/cp/list.json",
		Method: "POST",
		Data: map[string]interface{}{
			"pageNo": pageNo,
		},
	}

	// POST Map, default is JSON content type. No need to set one
	var resp CpResp
	err := s.SendRequest(context.TODO(), source, &resp)
	if err != nil {
		logx.Errorf("GetCpList 请求发生错误: %v", err)
		return nil, err
	}

	return resp.Result, nil
}

// GetTicketList 获取场次票种列表
func (s *ShowStart) GetTicketList(activityId string) ([]TicketListResult, error) {
	source := &SourceV3{
		URL:    "/wap/activity/V2/ticket/list",
		Method: "POST",
		Data: map[string]interface{}{
			"activityId": activityId,
			"coupon":     "", // 优惠卷 默认先为空吧
		},
	}

	var resp TicketResp
	err := s.SendRequest(context.TODO(), source, &resp)
	if err != nil {
		logx.Errorf("GetTicketList 请求发生错误: %v", err)
		return nil, err
	}

	if len(resp.Result) == 0 {
		return nil, fmt.Errorf("没有获取到对应场次票种列表")
	}

	return resp.Result, nil
}
func (s *ShowStart) GetWafToken() (*WafTokenResult, error) {
	// reset
	s.Cookies.AccessToken = ""
	s.Cookies.IdToken = ""

	source := &SourceV3{
		URL: "/waf/gettoken",
		Data: map[string]interface{}{
			"sign":    s.Cookies.Sign,
			"st_flpv": s.Cookies.StFlpv,
		},
	}

	var resp WafTokenResp
	err := s.SendRequest(context.TODO(), source, &resp)
	if err != nil {
		logx.Errorf("GetWafToken 请求发生错误: %v\n", err)
		return nil, err
	}

	s.Cookies.AccessToken = resp.Result.AccessToken.AccessToken
	s.Cookies.IdToken = resp.Result.IDToken.IDToken

	return &resp.Result, nil
}

func (s *ShowStart) SendRequest(ctx context.Context, source *SourceV3, response interface{}) error {
	request := s.Client.R().AddRetryCondition(func(resp *resty.Response, err error) bool {
		if err != nil {
			return false
		}
		var sr ResponseWrapper
		if err = json.Unmarshal(resp.Body(), &sr); err != nil {
			return false
		}
		// "state":"token-expire-at","sleep":0.4,"msg":"访问令牌已过期"
		// check token-expire-at
		if !strings.Contains(resp.Request.URL, "/waf/gettoken") && (getState(sr.State) == "token-expire-at" ||
			getState(sr.State) == "token-clean-at") {
			// get waf token then return true
			wafTokenResp, err := s.GetWafToken()
			if err != nil {
				logx.Errorf("get waf token err: %v", err)
				return false
			}
			//s.Infof("get waf token is: %+v", wafTokenResp)
			s.Cookies.AccessToken = wafTokenResp.AccessToken.AccessToken
			s.Cookies.IdToken = wafTokenResp.IDToken.IDToken
			// regen request
			w, err := NewWapEncryptV3(s.Cookies, source)
			if err != nil {
				logx.Errorf("retry gen encrypt v3 err: %v", err)
				return false
			}
			// set request
			resp.Request = resp.Request.SetBody(w.Source.Data).SetHeaders(w.Source.Headers)
			return true
		}
		return false
	})

	w, err := NewWapEncryptV3(s.Cookies, source)
	if err != nil {
		return err
	}

	resp, err := request.SetContext(ctx).SetHeaders(w.Source.Headers).SetBody(w.Source.Data).SetResult(&response).Post(w.GetRequestUrl())
	if err != nil {
		logx.Errorf("send request failed, resp: %v", resp.String())
		return errors.Wrap(err, "failed to send post request")
	}

	var sr ResponseWrapper
	if err = json.Unmarshal(resp.Body(), &sr); err != nil {
		logx.Errorf("resp.Body is %v\n", resp.String())
		return errors.Wrap(err, "json unmarshal")
	}

	return checkState(sr.State)
}
