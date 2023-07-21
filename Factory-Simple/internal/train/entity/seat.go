package train

import (
	"factorySimple/internal/core"
	valueObject "factorySimple/internal/train/valueObject"
)

// 座位
type Seat struct {
	Id     core.UUID // 物件追蹤ID
	SeatId int       // 位置ID
	Status string    // 狀態
	Price  int       // 價格
	Stand  valueObject.Stand
}

func CreatSeat(seatId int, stand valueObject.Stand) Seat {
	return Seat{
		Id:     core.NewUUID(),
		SeatId: seatId,
		Status: "create",
		Price:  0,
		Stand:  stand,
	}
}

func (s *Seat) ChangeStand(price int, stand valueObject.Stand) {
	s.Stand = stand
	s.Price = price
}
