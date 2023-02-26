package showstart

type AddressResp struct {
	Status  int       `json:"status"`
	State   string    `json:"state"`
	Result  []Address `json:"result"`
	TraceID string    `json:"traceId"`
}

type Address struct {
	ID           int    `json:"id"`
	Address      string `json:"address"`
	PostCode     string `json:"postCode"`
	Consignee    string `json:"consignee"`
	Telephone    string `json:"telephone"`
	IsDefault    int    `json:"isDefault"`
	ProvinceCode string `json:"provinceCode"`
	CityCode     string `json:"cityCode"`
	UserID       int    `json:"userId"`
	AreaCode     string `json:"areaCode"`
	CreateTime   string `json:"createTime"`
	ModifyTime   string `json:"modifyTime"`
	ProvinceName string `json:"provinceName"`
	CityName     string `json:"cityName"`
}

type CpResp struct {
	Status  int      `json:"status"`
	State   string   `json:"state"`
	Result  []CpItem `json:"result"`
	TraceId string   `json:"traceId"`
}

type CpItem struct {
	Id                 int    `json:"id"`
	UserId             int    `json:"userId"`
	Name               string `json:"name"`
	DocumentType       int    `json:"documentType"`
	DocumentTypeStr    string `json:"documentTypeStr"`
	ShowDocumentNumber string `json:"showDocumentNumber"`
	IsSelf             int    `json:"isSelf"`
	UpdateAuditStatus  int    `json:"updateAuditStatus"`
}

type TicketResp struct {
	Status  int                `json:"status"`
	State   string             `json:"state"`
	Result  []TicketListResult `json:"result"`
	TraceId string             `json:"traceId"`
}

type TicketListResult struct {
	SessionName          string       `json:"sessionName"`
	SessionId            int          `json:"sessionId"`
	IsConfirmedStartTime int          `json:"isConfirmedStartTime"`
	TicketList           []TicketItem `json:"ticketList"`
}

type TicketItem struct {
	TicketId              string `json:"ticketId"`
	TicketType            string `json:"ticketType"`
	SellingPrice          string `json:"sellingPrice"`
	CostPrice             string `json:"costPrice"`
	TicketNum             int    `json:"ticketNum"`
	ValidateType          int    `json:"validateType"`
	Time                  string `json:"time"`
	Instruction           string `json:"instruction"`
	Countdown             int    `json:"countdown"`
	RemainTicket          int    `json:"remainTicket"`
	SaleStatus            int    `json:"saleStatus"`
	ActivityId            int    `json:"activityId"`
	GoodType              int    `json:"goodType"`
	Telephone             string `json:"telephone"`
	AreaCode              string `json:"areaCode"`
	LimitBuyNum           int    `json:"limitBuyNum"`
	CanBuyNum             int    `json:"canBuyNum"`
	CityName              string `json:"cityName"`
	UnPayOrderNum         int    `json:"unPayOrderNum"`
	Type                  int    `json:"type"`
	BuyType               int    `json:"buyType"`
	CanAddGoods           int    `json:"canAddGoods"`
	TicketRecordStatus    int    `json:"ticketRecordStatus"`
	StartSellNoticeStatus int    `json:"startSellNoticeStatus"`
	ShowRuleTip           bool   `json:"showRuleTip"`
	StartTime             int64  `json:"startTime"`
	ShowTime              string `json:"showTime"`
	MemberNum             int    `json:"memberNum"`
}
