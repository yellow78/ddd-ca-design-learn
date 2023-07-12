package train

import (
	"valueObjectsSimple/internal/core"
	valueObject "valueObjectsSimple/internal/train/valueObject"
)

// 座位
type seat struct {
	Id     core.UUID // 物件追蹤ID
	SeatId int       // 位置ID
	Status string    // 狀態
	Price  int       // 價格
	Stand  valueObject.Stand
}

func CreatSeat(seatId int, stand valueObject.Stand) seat {
	return seat{
		Id:     core.NewUUID(),
		SeatId: seatId,
		Status: "create",
		Price:  0,
		Stand:  stand,
	}
}

func (s *seat) ChangeStand(price int, stand valueObject.Stand) {
	s.Stand = stand
	s.Price = price
}
