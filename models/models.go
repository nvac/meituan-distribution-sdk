package models

type MeituanResponse struct {
	Code      int    `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
	PartnerId int    `json:"partnerId,omitempty"`
	Result    any    `json:"result,omitempty"`
}

type Facility struct {
	Id    int `json:"id"`
	Name  int `json:"name"`
	Value int `json:"value"`
}

type Image struct {
	Category int    `json:"category"`
	Title    int    `json:"title"`
	Links    []Link `json:"links"`
}

type Link struct {
	Size        string `json:"size"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

type RoomContent struct {
}

type RealRoomBaseInfo struct {
	RealRoomId     int    `json:"realRoomId,omitempty"`
	HotelId        int64  `json:"hotelId,omitempty"`
	Name           string `json:"name,omitempty"`
	NameEn         string `json:"nameEn,omitempty"`
	UseAbleArea    string `json:"useAbleArea,omitempty"`
	Floor          string `json:"floor,omitempty"`
	Status         int    `json:"status,omitempty"`
	AllowExtraBed  bool   `json:"allowExtraBed,omitempty"`
	InternetAccess int    `json:"internetAccess,omitempty"`
}

type RoomDescription struct {
	Overview string `json:"overview"`
}

type Occupancy struct {
	Total int `json:"total"`
}

type Window struct {
	HasWindow    int    `json:"hasWindow"`
	Description  string `json:"description"`
	Disadvantage string `json:"disadvantage"`
}

type HotelBaseInfo struct {
	HotelId        int64  `json:"hotelId,omitempty"`
	PointName      string `json:"pointName,omitempty"`
	Info           string `json:"info,omitempty"`
	Longitude      int    `json:"longitude,omitempty"`
	Latitude       int    `json:"latitude,omitempty"`
	Address        string `json:"address,omitempty"`
	CityName       string `json:"cityName,omitempty"`
	CityLocationId int    `json:"cityLocationId,omitempty"`
	LocationName   string `json:"locationName,omitempty"`
	LocationId     int    `json:"locationId,omitempty"`
	BareaName      string `json:"bareaName,omitempty"`
	AvgScore       int    `json:"avgScore,omitempty"`
	Phone          string `json:"phone,omitempty"`
	OpenInfo       string `json:"openInfo,omitempty"`
	CloseStatus    int    `json:"closeStatus,omitempty"`
	FrontImage     string `json:"frontImage,omitempty"`
	BrandId        int    `json:"brandId,omitempty"`
	BrandName      string `json:"brandName,omitempty"`
}

type HotelExtendInfo struct {
	HotelFacilities map[int]string `json:"hotelFacilities,omitempty"`
	HotelService    map[int]string `json:"hotelService,omitempty"`
	PoiExtInfo      PoiExtInfo     `json:"poiExtInfo"`
}

type PoiExtInfo struct {
	HotelId           int64   `json:"hotelId,omitempty"`
	OpenDate          string  `json:"openDate,omitempty"`
	DecorationDate    string  `json:"decorationDate,omitempty"`
	RoomNum           int     `json:"roomNum,omitempty"`
	FloorNum          int     `json:"floorNum,omitempty"`
	HotelStar         int     `json:"hotelStar,omitempty"`
	PoiType           string  `json:"poiType,omitempty"`
	ThemeTag          string  `json:"themeTag,omitempty"`
	CheckinTimeBegin  string  `json:"checkinTimeBegin,omitempty"`
	CheckinTimeEnd    string  `json:"checkinTimeEnd,omitempty"`
	CheckoutTime      string  `json:"checkoutTime,omitempty"`
	CheckoutTimeHours float64 `json:"checkoutTimeHours,omitempty"`
	ForeignPolicy     int     `json:"foreignPolicy,omitempty"`
	HotelRemind       string  `json:"hotelRemind,omitempty"`
}

type PoiImage struct {
	TypeId   int    `json:"typeId,omitempty"`
	TypeName string `json:"typeName,omitempty"`
	ImgDesc  string `json:"imgDesc,omitempty"`
	Url      string `json:"url,omitempty"`
	RoomId   int    `json:"roomId,omitempty"`
}

type RealRoomInfo struct {
	RealRoomBaseInfo  RealRoomBaseInfo `json:"realRoomBaseInfo"`
	GoodsIds          []int            `json:"goodsIds"`
	RoomBedInfoList   [][]RoomBedInfo  `json:"roomBedInfoList"`
	RealRoomImageList []RealRoomImage  `json:"realRoomImageList"`
	RoomFacilities    map[int]string   `json:"roomFacilities"`
}

type RealRoomImage struct {
	ImgDesc string `json:"imgDesc"`
	Url     string `json:"url"`
	IsFront int    `json:"isFront"`
}

type HotelGoods struct {
	HotelId int64       `json:"hotelId"`
	Goods   []GoodsInfo `json:"goods"`
}

type GoodsInfo struct {
	HotelId               int64            `json:"hotelId,omitempty"`
	GoodsId               int64            `json:"goodsId,omitempty"`
	GoodsName             string           `json:"goodsName,omitempty"`
	NeedRealTel           int              `json:"needRealTel,omitempty"`
	GoodsStatus           int              `json:"goodsStatus,omitempty"`
	GoodsType             int              `json:"goodsType,omitempty"`
	ConfirmType           int              `json:"confirmType,omitempty"`
	InvRemain             int              `json:"invRemain,omitempty"`
	AveragePrice          int              `json:"averagePrice,omitempty"`
	ThirdParty            int              `json:"thirdParty,omitempty"`
	Breakfast             []Breakfast      `json:"breakfast,omitempty"`
	RoomInfoList          []GoodsRoomInfo  `json:"roomInfoList,omitempty"`
	CancelRules           []CancelRule     `json:"cancelRules,omitempty"`
	Gifts                 []Gift           `json:"gifts,omitempty"`
	BookRules             []BookRule       `json:"bookRules,omitempty"`
	InvoiceInfo           InvoiceInfo      `json:"invoiceInfo"`
	PriceModels           []PriceModel     `json:"priceModels,omitempty"`
	GoodsActivityMap      map[string]int   `json:"goodsActivityMap,omitempty"`
	GoodsStatuses         []DayGoodsStatus `json:"goodsStatuses,omitempty"`
	HourRoomReceptionTime []ReceptionTime  `json:"hourRoomReceptionTime,omitempty"`
	TimeIntervalMin       int              `json:"timeIntervalMin,omitempty"`
	TimeLimitMin          int              `json:"timeLimitMin,omitempty"`
	GoodsLimitRule        string           `json:"goodsLimitRule,omitempty"`
	PaymentType           int              `json:"paymentType,omitempty"`
	RpGuarantees          []RpGuarantee    `json:"rpGuarantees,omitempty"`
	NeedIdentityCard      bool             `json:"needIdentityCard,omitempty"`
	OriginalPrice         int              `json:"originalPrice,omitempty"`
}

type Breakfast struct {
	BreakfastType int `json:"breakfastType,omitempty"`
	BreakfastNum  int `json:"breakfastNum,omitempty"`
	InStartDate   int `json:"inStartDate,omitempty"`
	InEndDate     int `json:"inEndDate,omitempty"`
}

type GoodsRoomInfo struct {
	RoomId     int    `json:"roomId,omitempty"`
	RoomName   string `json:"roomName,omitempty"`
	CityId     int    `json:"cityId,omitempty"`
	RealRoomId int    `json:"realRoomId,omitempty"`
}

type CancelRule struct {
	CancelType       int    `json:"cancelType,omitempty"`
	AheadCancelDays  int    `json:"aheadCancelDays,omitempty"`
	DeductType       int    `json:"deductType,omitempty"`
	AheadCancelHours string `json:"aheadCancelHours,omitempty"`
}

type Gift struct {
	GiveType          int        `json:"giveType,omitempty"`
	GiftName          string     `json:"giftName,omitempty"`
	GiftDesc          string     `json:"giftDesc,omitempty"`
	GiftPrice         int        `json:"giftPrice,omitempty"`
	AvailableDateType int        `json:"availableDateType,omitempty"`
	AvailableDate     []GiftDate `json:"availableDate,omitempty"`
	AdultMax          string     `json:"adultMax,omitempty"`
	AdultMin          string     `json:"adultMin,omitempty"`
	Child             string     `json:"child,omitempty"`
	EntryMode         string     `json:"entryMode,omitempty"`
	GiftTemplate      int        `json:"giftTemplate,omitempty"`
	MealTimeStartTime string     `json:"mealTimeStartTime,omitempty"`
	MealTimeEndTime   string     `json:"mealTimeEndTime,omitempty"`
	SpoiName          string     `json:"spoiName,omitempty"`
	Saddress          string     `json:"saddress,omitempty"`
}

type GiftDate struct {
	StartDate string `json:"startDate,omitempty"`
	EndDate   string `json:"endDate,omitempty"`
}

type BookRule struct {
	SerialCheckinMin     int    `json:"serialCheckinMin,omitempty"`
	SerialCheckinMax     int    `json:"serialCheckinMax,omitempty"`
	RoomCountMin         int    `json:"roomCountMin,omitempty"`
	RoomCountMax         int    `json:"roomCountMax,omitempty"`
	EarliestBookingDays  int    `json:"earliestBookingDays,omitempty"`
	EarliestBookingHours string `json:"earliestBookingHours,omitempty"`
	LatestBookingDays    int    `json:"latestBookingDays,omitempty"`
	LatestBookingHours   string `json:"latestBookingHours,omitempty"`
	IsDaybreakBooking    int    `json:"isDaybreakBooking,omitempty"`
	InStartDate          int    `json:"inStartDate,omitempty"`
	InEndDate            int    `json:"inEndDate,omitempty"`
}

type InvoiceInfo struct {
	InvoiceMode    int           `json:"invoiceMode,omitempty"`
	SupportTypes   []InvoiceType `json:"supportTypes,omitempty"`
	TmcInvoiceMode int           `json:"tmcInvoiceMode,omitempty"`
}

type InvoiceType struct {
	Type    int `json:"type"`
	Postage int `json:"postage"`
}

type PriceModel struct {
	Date        string `json:"date,omitempty"`
	SalePrice   int    `json:"salePrice,omitempty"`
	SubPrice    int    `json:"subPrice,omitempty"`
	SubRatio    int    `json:"subRatio,omitempty"`
	DayType     int    `json:"dayType,omitempty"`
	LowestPrice int    `json:"lowestPrice,omitempty"`
}

type DayGoodsStatus struct {
	Date   string `json:"date,omitempty"`
	Status int    `json:"status,omitempty"`
}

type ReceptionTime struct {
	StartTime string `json:"startTime,omitempty"`
	EndTime   string `json:"endTime,omitempty"`
	Type      int    `json:"type,omitempty"`
}

type RpGuarantee struct {
	GuaranteeType     int          `json:"guaranteeType,omitempty"`
	GuaranteeTimeType int          `json:"guaranteeTimeType,omitempty"`
	DatePeriods       []DatePeriod `json:"datePeriods,omitempty"`
	ArrivalHour       string       `json:"arrivalHour,omitempty"`
}

type DatePeriod struct {
	StartDate string `json:"startDate,omitempty"`
	EndDate   string `json:"endDate,omitempty"`
}

type GoodsPrice struct {
	HotelId           int64        `json:"hotelId,omitempty"`
	GoodsId           int64        `json:"goodsId,omitempty"`
	PriceModels       []PriceModel `json:"priceModels,omitempty"`
	LowestPriceModels []PriceModel `json:"lowestPriceModels,omitempty"`
}

type GoodsStatus struct {
	GoodsId       int64            `json:"goodsId,omitempty"`
	Status        int              `json:"status,omitempty"`
	GoodsStatuses []DayGoodsStatus `json:"goodsStatuses,omitempty"`
}

type HotelCitySuggest struct {
	CityCode    int64  `json:"cityCode,omitempty"`
	Type        string `json:"type,omitempty"`
	Keyword     string `json:"keyword,omitempty"`
	HotelId     int64  `json:"hotelId,omitempty"`
	Address     string `json:"address,omitempty"`
	LowestPrice int    `json:"lowestPrice,omitempty"`
}

type Coordinate struct {
	Longitude int    `json:"longitude,omitempty"`
	Latitude  int    `json:"latitude,omitempty"`
	Radius    int    `json:"radius,omitempty"`
	Provider  string `json:"provider,omitempty"`
}

type HotelListSearchInfo struct {
	HotelId     int64  `json:"hotelId,omitempty"`
	Name        string `json:"name,omitempty"`
	Address     string `json:"address,omitempty"`
	LowestPrice string `json:"lowestPrice,omitempty"`
	Longitude   int    `json:"longitude,omitempty"`
	Latitude    int    `json:"latitude,omitempty"`
	Star        int    `json:"star,omitempty"`
	Rating      int    `json:"rating,omitempty"`
	BrandName   string `json:"brandName,omitempty"`
	GroupName   string `json:"groupName,omitempty"`
	FrontImg    string `json:"frontImg,omitempty"`
}

type RoomCriteria struct {
	RoomCount  int   `json:"roomCount,omitempty"`
	AdultCount int   `json:"adultCount,omitempty"`
	ChildCount int   `json:"childCount,omitempty"`
	ChildAges  []int `json:"childAges,omitempty"`
}

type Filter struct {
	MinimumPrice int `json:"minimumPrice,omitempty"`
	MaximumPrice int `json:"maximumPrice,omitempty"`
	ProductType  int `json:"productType,omitempty"`
}

type AvailProduct struct {
	ProductId      int64            `json:"productId,omitempty"`
	ProductBase    []ProductBase    `json:"productBase,omitempty"`
	Room           []ProductRoom    `json:"room,omitempty"`
	Invoice        []Invoice        `json:"invoice,omitempty"`
	BookingRule    []BookingRule    `json:"bookingRule,omitempty"`
	Services       []Services       `json:"services,omitempty"`
	InclusivePrice []InclusivePrice `json:"inclusivePrice,omitempty"`
	Promotion      []Promotion      `json:"promotion,omitempty"`
	Guarantee      []RpGuarantee    `json:"guarantee,omitempty"`
	CancelPolicy   []CancelPolicy   `json:"cancelPolicy,omitempty"`
}

type CancelPolicy struct {
	CancelType       int    `json:"cancelType,omitempty"`
	MoveUpCancelDays int    `json:"moveUpCancelDays,omitempty"`
	MoveUpCancelHour string `json:"moveUpCancelHour,omitempty"`
	Amount           int    `json:"amount,omitempty"`
	Nights           int    `json:"nights,omitempty"`
	Percent          int    `json:"percent,omitempty"`
}

type Promotion struct {
	OfferType        string `json:"offerType,omitempty"`
	ActivityType     string `json:"activityType,omitempty"`
	AvgDiscount      int    `json:"avgDiscount,omitempty"`
	TotalDiscount    int    `json:"totalDiscount,omitempty"`
	DiscountCalendar []int  `json:"discountCalendar,omitempty"`
	Tag              string `json:"tag,omitempty"`
}

type InclusivePrice struct {
	Date      string `json:"date,omitempty"`
	SalePrice int    `json:"salePrice,omitempty"`
	SubPrice  int    `json:"subPrice,omitempty"`
}

type Services struct {
	HourRoom []HourRoom `json:"hourRoom,omitempty"`
	MealType []MealType `json:"mealType,omitempty"`
	Gifts    []Gift     `json:"gifts,omitempty"`
}

type HourRoom struct {
	HourRoomReceptionTime []HourRoomReceptionTime `json:"hourRoomReceptionTime,omitempty"`
	TimeIntervalMin       int                     `json:"timeIntervalMin,omitempty"`
	TimeLimitMin          int                     `json:"timeLimitMin,omitempty"`
}

type HourRoomReceptionTime struct {
	StartTime string `json:"startTime,omitempty"`
	EndTime   string `json:"endTime,omitempty"`
}

type MealType struct {
	Code      string `json:"code,omitempty"`
	Desc      string `json:"desc,omitempty"`
	Count     int    `json:"count,omitempty"`
	StartDate int    `json:"startDate,omitempty"`
	EndDate   int    `json:"endDate,omitempty"`
}

type BookingRule struct {
	SerialCheckinMin     int    `json:"serialCheckinMin,omitempty"`
	SerialCheckinMax     int    `json:"serialCheckinMax,omitempty"`
	RoomCountMin         int    `json:"roomCountMin,omitempty"`
	RoomCountMax         int    `json:"roomCountMax,omitempty"`
	EarliestBookingDays  int    `json:"earliestBookingDays,omitempty"`
	EarliestBookingHours string `json:"earliestBookingHours,omitempty"`
	LatestBookingDays    int    `json:"latestBookingDays,omitempty"`
	LatestBookingHours   string `json:"latestBookingHours,omitempty"`
	IsDaybreakBooking    int    `json:"isDaybreakBooking,omitempty"`
	InStartDate          int    `json:"inStartDate,omitempty"`
	InEndDate            int    `json:"inEndDate,omitempty"`
}

type Invoice struct {
	InvoiceMode    int `json:"invoiceMode,omitempty"`
	TmcInvoiceMode int `json:"tmcInvoiceMode,omitempty"`
}

type ProductRoom struct {
	RoomId    int64   `json:"roomId,omitempty"`
	RoomName  string  `json:"roomName,omitempty"`
	BedGroups [][]Bed `json:"bedGroups,omitempty"`
	Window    int     `json:"window,omitempty"`
	ExtraBed  int     `json:"extraBed,omitempty"`
}

type Bed struct {
	BedType  string `json:"bedType,omitempty"`
	BedDesc  string `json:"bedDesc,omitempty"`
	BedCount int    `json:"bedCount,omitempty"`
}

type ProductBase struct {
	ProductName      string `json:"productName,omitempty"`
	Inventory        int    `json:"inventory,omitempty"`
	ProductType      int    `json:"productType,omitempty"`
	ConfirmType      int    `json:"confirmType,omitempty"`
	ProductStatus    int    `json:"productStatus,omitempty"`
	PaymentType      int    `json:"paymentType,omitempty"`
	ProductLimitRule string `json:"productLimitRule,omitempty"`
	NeedCertificate  bool   `json:"needCertificate,omitempty"`
}
