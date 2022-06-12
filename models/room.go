package models

type RoomExtendInfo struct {
	RoomFacilities map[int]string `json:"roomFacilities,omitempty"`
}

type RoomBedInfo struct {
	RoomId   int    `json:"roomId,omitempty"`
	BedType  string `json:"bedType,omitempty"`
	BedDesc  string `json:"bedDesc,omitempty"`
	BedCount int    `json:"bedCount,omitempty"`
}

type RoomInfo struct {
	RoomBaseInfo     RoomBaseInfo    `json:"roomBaseInfo"`
	RoomExtendInfo   RoomExtendInfo  `json:"roomExtendInfo"`
	RoomBedInfoLists [][]RoomBedInfo `json:"roomBedInfoLists,omitempty"`
}

type RoomBaseInfo struct {
	RoomId        int    `json:"roomId,omitempty"`
	RealRoomId    int    `json:"realRoomId,omitempty"`
	HotelId       int64  `json:"hotelId,omitempty"`
	RoomType      int    `json:"roomType,omitempty"`
	RoomName      string `json:"roomName,omitempty"`
	RoomDesc      string `json:"roomDesc,omitempty"`
	UseableArea   string `json:"useableArea,omitempty"`
	Capacity      string `json:"capacity,omitempty"`
	Window        int    `json:"window,omitempty"`
	WindowView    string `json:"windowView,omitempty"`
	WindowBad     string `json:"windowBad,omitempty"`
	ExtraBed      int    `json:"extraBed,omitempty"`
	Floor         string `json:"floor,omitempty"`
	InternetWay   int    `json:"internetWay,omitempty"`
	WeekdayPrice  int    `json:"weekdayPrice,omitempty"`
	Status        int    `json:"status,omitempty"`
	ChildCapacity string `json:"childCapacity,omitempty"`
	WeekendPrice  int    `json:"weekendPrice,omitempty"`
}
