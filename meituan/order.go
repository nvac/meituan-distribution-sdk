package meituan

import (
	"meituan-distribution-sdk/models"
)

// 下单前校验 / hotel.order.check

type HotelOrderCheckRequest struct {
	HotelId      int64  `json:"hotelIds,omitempty"`
	GoodsId      int64  `json:"goodsId,omitempty"`
	CheckinDate  string `json:"checkinDate,omitempty"`
	CheckoutDate string `json:"checkoutDate"`
	RoomNum      int    `json:"roomNum,omitempty"`
	TotalPrice   int    `json:"totalPrice"`
}

type HotelOrderCheckResponse struct {
	Code          int                 `json:"code,omitempty"`
	Desc          string              `json:"desc,omitempty"`
	PriceModels   []models.PriceModel `json:"priceModels,omitempty"`
	RemainRoomNum int                 `json:"remainRoomNum"`
}

func (c *Client) HotelOrderCheck(request *HotelShoppingMultiRequest) (*HotelShoppingMultiResponse, error) {
	e := envelope{
		Method: "hotel.order.check",
		Data:   request,
	}

	response := &HotelShoppingMultiResponse{}
	_, _, _, err := c.callV1(e, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// 预约下单 / hotel.order.booking

type HotelOrderBookingRequest struct {
	HotelId            int64  `json:"hotelId,omitempty"`
	GoodsId            int64  `json:"goodsId,omitempty"`
	PersonNames        string `json:"personNames,omitempty"`
	ContactName        string `json:"contactName,omitempty"`
	ContactPhone       string `json:"contactPhone,omitempty"`
	ArriveDate         string `json:"arriveDate,omitempty"`
	CheckinDate        string `json:"checkinDate,omitempty"`
	CheckoutDate       string `json:"checkoutDate,omitempty"`
	RoomNum            int    `json:"roomNum,omitempty"`
	TotalPrice         int    `json:"totalPrice,omitempty"`
	SettlePrice        int    `json:"settlePrice,omitempty"`
	DistributorOrderId string `json:"distributorOrderId,omitempty"`
	PersonIdentities   string `json:"personIdentities,omitempty"`
	NeedInvoice        int    `json:"needInvoice,omitempty"`
}

type HotelOrderBookingResponse struct {
	DistributorOrderId string `json:"distributorOrderId,omitempty"`
	MtOrderId          int64  `json:"mtOrderId,omitempty"`
	Code               int    `json:"code,omitempty"`
	Desc               string `json:"desc,omitempty"`
}

func (c *Client) HotelOrderBooking(request *HotelShoppingMultiRequest) (*HotelShoppingMultiResponse, error) {
	e := envelope{
		Method: "hotel.order.booking",
		Data:   request,
	}

	response := &HotelShoppingMultiResponse{}
	_, _, _, err := c.callV1(e, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// 取消订单 / hotel.order.cancel

type HotelOrderCancelRequest struct {
	DistributorOrderId string `json:"distributorOrderId,omitempty"`
	MtOrderId          int64  `json:"mtOrderId,omitempty"`
	CancelReason       string `json:"cancelReason"`
	CancelCheck        int    `json:"cancelCheck"`
}

type HotelOrderCancelResponse struct {
	DistributorOrderId string `json:"distributorOrderId,omitempty"`
	MtOrderId          int64  `json:"mtOrderId,omitempty"`
	Code               int    `json:"code,omitempty"`
	Desc               string `json:"desc,omitempty"`
}

func (c *Client) HotelOrderCancel(request *HotelOrderCancelRequest) (*HotelOrderCancelResponse, error) {
	e := envelope{
		Method: "hotel.order.cancel",
		Data:   request,
	}

	response := &HotelOrderCancelResponse{}
	_, _, _, err := c.callV1(e, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// 查询订单 / hotel.order.query

type HotelOrderQueryRequest struct {
	QueryParams []models.OrderQueryItem `json:"queryParams,omitempty"`
}

type HotelOrderQueryResponse struct {
	OrderInfos []models.OrderInfo `json:"orderInfos,omitempty"`
	Code       int                `json:"code,omitempty"`
	Desc       string             `json:"desc,omitempty"`
}

func (c *Client) HotelOrderQuery(request *HotelOrderCancelRequest) (*HotelOrderCancelResponse, error) {
	e := envelope{
		Method: "hotel.order.query",
		Data:   request,
	}

	response := &HotelOrderCancelResponse{}
	_, _, _, err := c.callV1(e, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// 酒店催单 / hotel.order.reminder

type HotelOrderReminderRequest struct {
	MtOrderId int64 `json:"mtOrderId,omitempty"`
}

type HotelOrderReminderResponse struct {
	MtOrderId  int64  `json:"mtOrderId,omitempty"`
	CreateTime string `json:"createTime"`
	Code       int    `json:"code"`
	Desc       string `json:"desc"`
}

func (c *Client) HotelOrderReminder(request *HotelOrderReminderRequest) (*HotelOrderReminderResponse, error) {
	e := envelope{
		Method: "hotel.order.reminder",
		Data:   request,
	}

	response := &HotelOrderReminderResponse{}
	_, _, _, err := c.callV1(e, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
