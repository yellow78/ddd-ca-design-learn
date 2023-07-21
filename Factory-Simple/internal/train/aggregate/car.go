package train

import (
	"factorySimple/internal/core"
	entity "factorySimple/internal/train/entity"
	valueObject "factorySimple/internal/train/valueObject"
)

type car struct {
	Id     core.UUID // 物件追蹤ID
	CardId int
	Seat   []entity.Seat
}

// 創建車次
func CreateCard(cardId int, seatNum int, start string, end string) car {
	var seat []entity.Seat
	initStand := valueObject.CreatStand(start, end)

	for i := 0; i < seatNum; i++ {
		newSeat := entity.CreatSeat(i, initStand)
		seat = append(seat, newSeat)
	}

	return car{
		Id:     core.NewUUID(),
		CardId: cardId,
		Seat:   seat,
	}
}
