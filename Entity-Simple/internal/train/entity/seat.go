package entity

import "entitySimple/internal/core"

// 座位
type seat struct {
	id     core.UUID // 物件追蹤ID
	SeatId int       // 位置ID
	Stand  stand     // 旅程
	Status string    // 狀態
	Price  int       // 價格
}

// 旅程
type stand struct {
	Start string
	End   string
}

func CreatSeat(seatId int) seat {
	return seat{
		id:     core.NewUUID(),
		SeatId: seatId,
		Stand: stand{
			Start: "",
			End:   "",
		},
		Status: "create",
		Price:  0,
	}
}
