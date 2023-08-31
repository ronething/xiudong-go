package showstart

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

type WapEncryptConfig struct {
	Sign   string `json:"sign"`
	StFlpv string `json:"st_flpv"`
	Token  string `json:"token"`
	UserId uint32 `json:"user_id"`

	// AesKey Deprecated 字段弃用 默认值已写死
	AesKey string `json:"aes_key"`
}

type Source struct {
	URL        string                 `json:"url"`
	Method     string                 `json:"method"`
	Data       map[string]interface{} `json:"data"`
	APISource  string                 `json:"apiSource"`
	APIId      uint32                 `json:"apiId"`
	EventID    uint32                 `json:"eventId"`
	ParamsType string                 `json:"paramsType"`
	Headers    map[string]string      `json:"headers"`
}

// WapEncrypt 网页版加密
type WapEncrypt struct {
	Config *WapEncryptConfig
	Source *Source
}

// Deprecated: NewWapEncrypt use v3 instead
func NewWapEncrypt(config *WapEncryptConfig, source *Source) (*WapEncrypt, error) {
	if config.Token == "" { // fix issue #36
		config.Token = getRandStr(32)
	}
	w := &WapEncrypt{
		Config: config,
		Source: source,
	}
	err := w.InjectSource()
	if err != nil {
		return nil, err
	}

	return w, nil
}

// GetRequestUrl 获取请求 url
func (w *WapEncrypt) GetRequestUrl() string {
	if w.Source.APISource == "" {
		w.Source.APISource = "hw"
	}

	b := Base62{}

	apiIdBase62 := b.Encode(w.Source.APIId)
	prefix2 := fmt.Sprintf("00%s", apiIdBase62)
	prefix2 = prefix2[len(prefix2)-2:]

	eventIdBase62 := b.Encode(w.Source.EventID)
	prefix4 := fmt.Sprintf("0000%s", eventIdBase62)
	prefix4 = prefix4[len(prefix4)-4:]

	userIdBase62 := b.Encode(w.Config.UserId)
	prefix6 := fmt.Sprintf("000000%s", userIdBase62)
	prefix6 = prefix6[len(prefix6)-6:]

	apiUrlPrefix := "https://wap.showstart.com/api" // 先写死
	apiUrlSuffix := fmt.Sprintf("%s%s%s", prefix2, prefix4, prefix6)

	requestUrl := fmt.Sprintf("%s/%s/%s", apiUrlPrefix, w.Source.APISource, apiUrlSuffix)

	return requestUrl
}

func (w *WapEncrypt) InjectSource() error {
	data, err := w.getInjectData(nowTimeMs(), getRandStr(8))
	if err != nil {
		return err
	}
	w.Source.Data = data

	headers := w.getInjectHeaders(time.Now().Unix())
	w.Source.Headers = headers

	return nil
}

type sourceDataTmp struct {
	Action  string                 `json:"action"`
	Method  string                 `json:"method"`
	Query   map[string]interface{} `json:"query"`
	Body    map[string]interface{} `json:"body"`
	QTime   int64                  `json:"qtime"`
	RandStr string                 `json:"ranstr"` // 这个 key 是 ranstr
}

// 暂时不需要用到 仅作为标识
type sourceDataNew struct {
	Data     string `json:"data"`
	Sign     string `json:"sign"`
	AppId    string `json:"appid"`
	Terminal string `json:"terminal"`
	Version  string `json:"version"`
}

// getInjectData data 注入
func (w *WapEncrypt) getInjectData(qtime int64, randStr string) (map[string]interface{}, error) {
	if w.Source.Data == nil {
		w.Source.Data = make(map[string]interface{})
	}
	data := w.Source.Data
	_, ok := data["st_flpv"] // 如果已有就不用赋值
	if !ok {
		data["st_flpv"] = w.Config.StFlpv
	}
	_, ok = data["sign"] // 如果已有就不用赋值
	if !ok {
		data["sign"] = w.Config.Sign
	}
	data["trackPath"] = ""
	data["terminal"] = "wap"

	if w.Source.ParamsType == "" {
		w.Source.ParamsType = "query"
	}

	tmpData := sourceDataTmp{
		Action:  strings.TrimSuffix(w.Source.URL, ".json"),
		Method:  w.Source.Method,
		Query:   nil,
		Body:    nil,
		QTime:   qtime,
		RandStr: randStr,
	}

	// 理论上一定能走到的
	switch w.Source.ParamsType {
	case "query":
		tmpData.Query = data
	case "body":
		tmpData.Body = data
	}

	dumpRes, err := json.Marshal(tmpData)
	if err != nil {
		log.Printf("tmpData: %+v, 序列化失败: %v\n", tmpData, err)
		return nil, err
	}

	//log.Printf("序列化之后的数据: %v\n", string(dumpRes))

	dataMd5Lower := strings.ToLower(Md5SumByte(dumpRes))

	// 加解密
	aesCrypto := AESCrypto{Key: []byte("0RGF99CtUajPF0Ny")}
	encryptRes, err := aesCrypto.Encrypt(dumpRes)
	if err != nil {
		log.Printf("加密失败: %v", err)
		return nil, err
	}
	//decryptRes, err := aesCrypto.Decrypt(encryptRes)
	//if err != nil {
	//	log.Printf("解密失败: %v", err)
	//	return nil, err
	//}
	//log.Printf("解密后的数据: %v", decryptRes)

	newData := map[string]interface{}{
		"data":     encryptRes,
		"sign":     dataMd5Lower,
		"appid":    "wap",
		"terminal": "wap",
		"version":  "997",
	}

	//log.Printf("newData is %v", newData)

	return newData, nil
}

// getInjectHeaders 注入 request headers
// now 为秒级时间戳
func (w *WapEncrypt) getInjectHeaders(now int64) map[string]string {
	headers := map[string]string{
		"cookie":          fmt.Sprintf("Hm_lvt_da038bae565bb601b53cc9cb25cdca74=1638605081; Hm_lpvt_da038bae565bb601b53cc9cb25cdca74=%d", now),
		"content-type":    "application/json",
		"st_flpv":         w.Config.StFlpv,
		"sign":            w.Config.Sign,
		"terminal":        "wap",
		"user-agent":      "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1",
		"sec-fetch-dest":  "empty",
		"sec-fetch-mode":  "cors",
		"sec-fetch-site":  "same-origin",
		"accept-language": "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7",
		"accept-encoding": "gzip, deflate, br",
		"origin":          "https://wap.showstart.com",
	}
	key := "xVgXtOUSos6jzR3mqb4aLHYybqqPFFGfx12r"    // 之后看看是不是要作为配置 毕竟随时可能发生更改
	urlPrefix := strings.Split(w.Source.URL, "?")[0] // url 前缀如 /wap/V2/ticket/list.json
	headers["r"] = fmt.Sprintf("%d", now*1000)
	headers["cusystime"] = headers["r"]
	headers["s"] = strings.ToLower(Md5Sum(fmt.Sprintf("%d%s%s%s", now*1000, urlPrefix, w.Config.StFlpv, key)))
	headers["cusut"] = w.Config.Sign
	headers["cuuserref"] = w.Config.Token
	headers["ctrackpath"] = "" // trackReferer 先默认 ""
	headers["csourcepath"] = ""

	return headers
}

func nowTimeMs() int64 {
	return time.Now().Unix() * 1000 // 毫秒时间戳
}
