package meituan

// 查询跳转订单 / hotel.jump.order.query

type HotelJumpOrderRequest struct {
	QueryType     int    `json:"query_type,omitempty"`
	StartTime     string `json:"start_time,omitempty"`
	EndTime       string `json:"end_time,omitempty"`
	OrderStatus   int    `json:"orderStatus,omitempty"`
	PageSize      int    `json:"pageSize,omitempty"`
	PageNo        int    `json:"pageNo,omitempty"`
	PositionIndex string `json:"positionIndex,omitempty"`
	JumpType      bool   `json:"jumpType,omitempty"`
}

type HotelJumpOrderResponse struct {
}

func (c *Client) HotelJumpOrder(request *HotelJumpOrderRequest) (*HotelJumpOrderResponse, error) {
	e := envelope{
		Method: "hotel.realroom.info",
		Data:   request,
	}

	response := &HotelJumpOrderResponse{}
	_, _, _, err := c.callV1(e, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// 生成跳转推广链接 / hotel.jump.referral.link

type HotelJumpReferralLinkRequest struct {
	QueryType     int    `json:"query_type,omitempty"`
	StartTime     string `json:"start_time,omitempty"`
	EndTime       string `json:"end_time,omitempty"`
	OrderStatus   int    `json:"orderStatus,omitempty"`
	PageSize      int    `json:"pageSize,omitempty"`
	PageNo        int    `json:"pageNo,omitempty"`
	PositionIndex string `json:"positionIndex,omitempty"`
	JumpType      bool   `json:"jumpType,omitempty"`
}

type HotelJumpReferralLinkResponse struct {
}

func (c *Client) HotelJumpReferralLink(request *HotelJumpReferralLinkRequest) (*HotelJumpReferralLinkResponse, error) {
	e := envelope{
		Method: "hotel.realroom.info",
		Data:   request,
	}

	response := &HotelJumpReferralLinkResponse{}
	_, _, _, err := c.callV1(e, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
