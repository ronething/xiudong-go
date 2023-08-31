package showstart

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type WapEncryptConfigV3 struct {
	Sign        string `json:"sign"`
	StFlpv      string `json:"st_flpv"`
	Token       string `json:"token"`
	UserId      uint32 `json:"user_id"`
	AccessToken string `json:"accessToken"`
	IdToken     string `json:"idToken"`
}

type WapEncryptV3 struct {
	Config *WapEncryptConfigV3
	Source SourceV3
}

type SourceV3 struct {
	URL        string                 `json:"url"`
	Method     string                 `json:"method"`
	Data       map[string]interface{} `json:"data"`
	APISource  string                 `json:"apiSource"`
	APIId      uint32                 `json:"apiId"`
	EventID    string                 `json:"eventId"`
	ParamsType string                 `json:"paramsType"`
	Headers    map[string]string      `json:"headers"`
	Secret     bool                   `json:"secret"`
}

func NewWapEncryptV3(config *WapEncryptConfigV3, source *SourceV3) (*WapEncryptV3, error) {
	if config.Token == "" { // fix issue #36
		config.Token = getRandStr(32)
	}
	w := &WapEncryptV3{
		Config: config,
		Source: *source,
	}
	err := w.InjectSource()
	if err != nil {
		return nil, err
	}

	return w, nil
}

func (w *WapEncryptV3) InjectSource() error {
	w.Source.URL = strings.TrimSuffix(w.Source.URL, ".json")
	if w.Source.APISource == "appnj" {
		w.Source.APISource = "nj"
	}
	if len(w.Source.APISource) > 0 { // need concat
		w.Source.URL = "/" + w.Source.APISource + w.Source.URL
	}
	if w.Source.Data == nil {
		w.Source.Data = map[string]interface{}{}
	}
	w.Source.Data["st_flpv"] = w.Config.StFlpv
	w.Source.Data["sign"] = w.Config.Sign
	w.Source.Data["trackPath"] = ""

	headers, err := w.getInjectHeaders(time.Now().Unix())
	if err != nil {
		return err
	}
	w.Source.Headers = headers

	return nil
}

func (w *WapEncryptV3) GetRequestUrl() string {
	apiUrlPrefix := "https://wap.showstart.com/v3"
	return apiUrlPrefix + w.Source.URL
}

const deviceInfo = "%7B%22vendorName%22:%22%22,%22deviceMode%22:%22iPhone%22,%22deviceName%22:%22%22,%22systemName%22:%22ios%22,%22systemVersion%22:%2213.2.3%22,%22cpuMode%22:%22%20%22,%22cpuCores%22:%22%22,%22cpuArch%22:%22%22,%22memerySize%22:%22%22,%22diskSize%22:%22%22,%22network%22:%22UNKNOWN%22,%22resolution%22:%22390*844%22,%22pixelResolution%22:%22%22%7D"

func (w *WapEncryptV3) getInjectHeaders(now int64) (map[string]string, error) {
	headers := map[string]string{
		"Cookie":       fmt.Sprintf("Hm_lvt_da038bae565bb601b53cc9cb25cdca74=1689569273; Hm_lpvt_da038bae565bb601b53cc9cb25cdca74=%d", now),
		"Content-Type": "application/json",
		"st_flpv":      w.Config.StFlpv,
		"User-Agent":   "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1",
		"Origin":       "https://wap.showstart.com",
		"Referer":      "https://wap.showstart.com",
	}

	headers["CUSAT"] = valueOrNil(w.Config.AccessToken)
	headers["CUSUT"] = valueOrNil(w.Config.Sign)
	headers["CUSIT"] = valueOrNil(w.Config.IdToken)
	headers["CUSID"] = valueOrNil(fmt.Sprintf("%d", w.Config.UserId))
	headers["CUSNAME"] = "nil"
	headers["CDEVICENO"] = w.Config.Token
	headers["CUUSERREF"] = w.Config.Token
	headers["CVERSION"] = "997"
	headers["CTERMINAL"] = "wap"
	headers["CSAPPID"] = "wap"

	headers["CDEVICEINFO"] = deviceInfo
	headers["CRTRACEID"] = fmt.Sprintf("%s%d", getRandStr(32), now*1000)

	headers["CTRACKPATH"] = ""
	headers["CSOURCEPATH"] = ""

	// 生成签名
	dumpRes, err := json.Marshal(w.Source.Data)
	if err != nil {
		return nil, errors.Wrap(err, "marshal source data again")
	}
	// accessToken + sign + idToken + userId + "wap" + token + I + n['url'] + "997" + "wap" + traceId
	signValue := w.Config.AccessToken + w.Config.Sign + w.Config.IdToken + fmt.Sprintf("%d", w.Config.UserId) + "wap" +
		w.Config.Token + string(dumpRes) + w.Source.URL + "997" +
		"wap" + headers["CRTRACEID"]
	headers["CRPSIGN"] = Md5Sum(signValue)

	return headers, nil
}

func valueOrNil(a string) string {
	if len(a) > 0 {
		return a
	}
	return "nil"
}
