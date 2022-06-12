package meituan

import (
	"meituan-distribution-sdk/models"
)

// 查询酒店id列表 / hotel.poi.list

type HotelPoiListRequest struct {
	MaxId    int64 `json:"hotelId"`
	PageSize int   `json:"pageSize"`
}

type HotelPoiListResponse struct {
	MaxId    int64   `json:"maxId"`
	HotelIds []int64 `json:"hotelIds"`
}

func (c *Client) HotelPoiList(request *HotelPoiListRequest) (*HotelPoiListResponse, error) {
	e := envelope{
		Method: "hotel.poi.list",
		Data:   request,
	}

	response := &HotelPoiListResponse{}
	_, _, _, err := c.callV1(e, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// 查询酒店详情 / hotel.detail

type HotelDetailRequest struct {
	Strategy int     `json:"strategy"`
	HotelIds []int64 `json:"hotelIds"`
}

type HotelDetailResponse struct {
	BaseInfo   models.HotelBaseInfo   `json:"baseInfo"`
	ExtendInfo models.HotelExtendInfo `json:"extendInfo"`
	RoomInfos  []models.RoomInfo      `json:"roomInfos"`
	PoiImages  []models.PoiImage      `json:"poiImages"`
	HotelId    int64                  `json:"hotelId"`
}

func (c *Client) HotelDetail(request *HotelDetailRequest) (*HotelDetailResponse, error) {
	e := envelope{
		Method: "hotel.detail",
		Data:   request,
	}

	response := &HotelDetailResponse{}
	_, _, _, err := c.callV1(e, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// 查询酒店物理房型信息 / hotel.realroom.info

type HotelRealroomInfoRequest struct {
	HotelIds []int64 `json:"hotelIds"`
}

type HotelRealroomInfoResponse struct {
	RealRoomInfos map[int][]models.RealRoomInfo `json:"realRoomInfos"`
}

func (c *Client) HotelRealroomInfo(request *HotelRealroomInfoRequest) (*HotelRealroomInfoResponse, error) {
	e := envelope{
		Method: "hotel.realroom.info",
		Data:   request,
	}

	response := &HotelRealroomInfoResponse{}
	_, _, _, err := c.callV1(e, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// 查询酒店下产品列表 / hotel.goods.rp

type HotelGoodsRpRequest struct {
	HotelIds            []int64  `json:"hotelIds,omitempty"`
	CheckinDate         string   `json:"checkinDate,omitempty"`
	CheckoutDate        string   `json:"checkoutDate,omitempty"`
	GoodsType           int      `json:"goodsType,omitempty"`
	GoodsIds            []int64  `json:"goodsIds,omitempty"`
	QueryInfoDimensions []string `json:"queryInfoDimensions,omitempty"`
}

type HotelGoodsRpResponse struct {
	HotelGoods []models.HotelGoods `json:"hotelGoods"`
}

func (c *Client) HotelGoodsRp(request *HotelGoodsRpRequest) (*HotelGoodsRpResponse, error) {
	e := envelope{
		Method: "hotel.realroom.info",
		Data:   request,
	}

	response := &HotelGoodsRpResponse{}
	_, _, _, err := c.callV1(e, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// 查询产品价格日历 / hotel.goods.price

type HotelGoodsPriceRequest struct {
	GoodsIds  []int64 `json:"goodsIds"`
	HotelIds  []int64 `json:"hotelIds"`
	StartDate string  `json:"startDate"`
	EndDate   string  `json:"endDate"`
	PriceType int     `json:"priceType"`
}

type HotelGoodsPriceResponse struct {
	GoodsPrices []models.GoodsPrice `json:"goodsPrices"`
}

func (c *Client) HotelGoodsPrice(request *HotelGoodsPriceRequest) (*HotelGoodsPriceResponse, error) {
	e := envelope{
		Method: "hotel.realroom.info",
		Data:   request,
	}

	response := &HotelGoodsPriceResponse{}
	_, _, _, err := c.callV1(e, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// 查询产品实时房态 / hotel.goods.status

type HotelGoodsStatusRequest struct {
	HotelId      int64  `json:"hotelId"`
	CheckinDate  string `json:"checkinDate"`
	CheckoutDate string `json:"checkoutDate"`
	GoodsType    int    `json:"goodsType"`
}

type HotelGoodsStatusResponse struct {
	GoodsStatuses []models.GoodsStatus `json:"goodsStatuses"`
}

func (c *Client) HotelGoodsStatus(request *HotelGoodsStatusRequest) (*HotelGoodsStatusResponse, error) {
	e := envelope{
		Method: "hotel.realroom.info",
		Data:   request,
	}

	response := &HotelGoodsStatusResponse{}
	_, _, _, err := c.callV1(e, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// 目的地搜索 / hotel.city.suggest

type HotelCitySuggestRequest struct {
	Checkin   string `json:"checkin,omitempty"`
	Checkout  string `json:"checkout,omitempty"`
	CityCode  int64  `json:"cityCode,omitempty"`
	QueryText string `json:"queryText,omitempty"`
}

type HotelCitySuggestResponse struct {
	HotelCitySuggest []models.HotelCitySuggest `json:"hotelCitySuggest"`
}

func (c *Client) HotelCitySuggest(customerSessionId string, request *HotelCitySuggestRequest) (int, string, int, *HotelCitySuggestResponse, error) {
	e := envelope{
		Method:            "hotel.city.suggest",
		CustomerSessionId: customerSessionId,
		Data:              request,
	}

	response := &HotelCitySuggestResponse{}
	code, message, partnerId, err := c.callV2(e, response)
	if err != nil {
		return 0, "", 0, nil, err
	}

	return code, message, partnerId, response, nil
}

// 关键词搜索 / hotel.keyword.suggest

type HotelKeywordSuggestRequest struct {
	Checkin   string `json:"checkin,omitempty"`
	Checkout  string `json:"checkout,omitempty"`
	CityCode  int64  `json:"cityCode,omitempty"`
	QueryText string `json:"queryText,omitempty"`
}

type HotelKeywordSuggestResponse struct {
	HotelCitySuggest []models.HotelCitySuggest `json:"hotelCitySuggest"`
}

func (c *Client) HotelKeywordSuggest(request *HotelKeywordSuggestRequest) (*HotelKeywordSuggestResponse, error) {
	e := envelope{
		Method: "hotel.realroom.info",
		Data:   request,
	}

	response := &HotelKeywordSuggestResponse{}
	_, _, _, err := c.callV2(e, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// 列表页搜索 / hotel.poi.list.search

type HotelPoiListSearchRequest struct {
	BizType    int               `json:"bizType,omitempty"`
	Checkin    string            `json:"checkin,omitempty"`
	Checkout   string            `json:"checkout,omitempty"`
	CityCode   int64             `json:"cityCode,omitempty"`
	Coordinate models.Coordinate `json:"coordinate"`
	NextId     string            `json:"nextId,omitempty"`
	PageSize   int               `json:"pageSize,omitempty"`
	SortType   int               `json:"sortType,omitempty"`
	QueryText  string            `json:"queryText,omitempty"`
	Star       []int             `json:"star,omitempty"`
}

type HotelPoiListSearchResponse struct {
	TotalCount          int64                        `json:"totalCount,omitempty"`
	NextId              string                       `json:"nextId,omitempty"`
	HotelListSearchInfo []models.HotelListSearchInfo `json:"hotelListSearchInfo,omitempty"`
}

func (c *Client) HotelPoiListSearch(request *HotelPoiListSearchRequest) (*HotelPoiListSearchResponse, error) {
	e := envelope{
		Method: "hotel.realroom.info",
		Data:   request,
	}

	response := &HotelPoiListSearchResponse{}
	_, _, _, err := c.callV2(e, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// 批量拉取缓存酒店产品信息 / hotel.shopping.multi

type HotelShoppingMultiRequest struct {
	HotelIds     []int64             `json:"hotelIds,omitempty"`
	Checkin      string              `json:"checkin,omitempty"`
	Checkout     string              `json:"checkout,omitempty"`
	RoomCriteria models.RoomCriteria `json:"roomCriteria"`
	Currency     string              `json:"currency,omitempty"`
	Filter       models.Filter       `json:"filter"`
}

type HotelShoppingMultiResponse struct {
	HotelId      int64                 `json:"hotelId,omitempty"`
	AvailProduct []models.AvailProduct `json:"availProduct,omitempty"`
}

func (c *Client) HotelShoppingMulti(request *HotelShoppingMultiRequest) (*HotelShoppingMultiResponse, error) {
	e := envelope{
		Method: "hotel.realroom.info",
		Data:   request,
	}

	response := &HotelShoppingMultiResponse{}
	_, _, _, err := c.callV2(e, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
