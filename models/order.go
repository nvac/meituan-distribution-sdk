package models

type OrderQueryItem struct {
	DistributorOrderId string `json:"distributorOrderId,omitempty"`
	MtOrderId          int64  `json:"mtOrderId,omitempty"`
}

type OrderInfo struct {
	BaseInfo        OrderBaseInfo       `json:"baseInfo"`
	AptInfo         OrderAptInfo        `json:"aptInfo"`
	RoomNights      []RoomNightInfo     `json:"roomNights,omitempty"`
	RefundInfos     []OrderRefundInfo   `json:"refundInfos,omitempty"`
	OrderSplitInfos HotelOrderSplitInfo `json:"orderSplitInfos"`
}

type OrderBaseInfo struct {
	MtOrderId          int64  `json:"mtOrderId,omitempty"`
	GoodsId            int    `json:"goodsId,omitempty"`
	TotalPrice         int    `json:"totalPrice,omitempty"`
	SettlePrice        int    `json:"settlePrice,omitempty"`
	CreateTime         string `json:"createTime,omitempty"`
	OrderStatus        int    `json:"orderStatus,omitempty"`
	FixRoom            bool   `json:"fixRoom,omitempty"`
	GoodsType          int    `json:"goodsType,omitempty"`
	ConfirmationNumber string `json:"confirmationNumber,omitempty"`
	Misc               string `json:"misc,omitempty"`
}

type OrderAptInfo struct {
	MtOrderId    int64  `json:"mtOrderId,omitempty"`
	CheckinTime  string `json:"checkinTime,omitempty"`
	CheckoutTime string `json:"checkoutTime,omitempty"`
	ArriveTime   string `json:"arriveTime,omitempty"`
	Comment      string `json:"comment,omitempty"`
	RoomName     string `json:"roomName,omitempty"`
	RoomId       int    `json:"roomId,omitempty"`
	RoomCount    int    `json:"roomCount,omitempty"`
	HotelId      int64  `json:"hotelId,omitempty"`
	PoiName      string `json:"poiName,omitempty"`
	PersonNames  string `json:"personNames,omitempty"`
	ContactName  string `json:"contactName,omitempty"`
	ContactPhone string `json:"contactPhone,omitempty"`
}

type RoomNightInfo struct {
	BizDate           string `json:"bizDate,omitempty"`
	AppointmentStatus int    `json:"appointmentStatus,omitempty"`
	PayStatus         int    `json:"payStatus,omitempty"`
	SellPrice         int    `json:"sellPrice,omitempty"`
	SubPrice          int    `json:"subPrice,omitempty"`
}

type OrderRefundInfo struct {
	RefundId          int    `json:"refundId,omitempty"`
	RefundPrice       int    `json:"refundPrice,omitempty"`
	RefundSettlePrice int    `json:"refundSettlePrice,omitempty"`
	RefundTime        string `json:"refundTime,omitempty"`
}

type HotelOrderSplitInfo struct {
	Date           string `json:"date,omitempty"`
	RoomCount      int    `json:"roomCount,omitempty"`
	MtPrice        int    `json:"mtPrice,omitempty"`
	SettlePrice    int    `json:"settlePrice,omitempty"`
	SubPrice       int    `json:"subPrice,omitempty"`
	DtorSubRatio   int    `json:"dtorSubRatio,omitempty"`
	OrderSplitTime string `json:"orderSplitTime,omitempty"`
}
