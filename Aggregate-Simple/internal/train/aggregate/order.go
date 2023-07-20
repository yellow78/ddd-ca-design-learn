package train

import (
	"aggregateSimple/internal/core"
	entity "aggregateSimple/internal/train/entity"
)

type order struct {
	Id   core.UUID // 物件追蹤ID
	User entity.User
	Seat entity.Seat
}

// user 下單
func CreateOrder(user entity.User, seat entity.Seat) order {
	return order{
		Id:   core.NewUUID(),
		User: user,
		Seat: seat,
	}
}

// user 改單
func (o *order) ChangeOrder(seat entity.Seat) {
	o.Seat = seat
}
