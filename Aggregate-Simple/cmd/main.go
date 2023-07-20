package main

import (
	aggregate "aggregateSimple/internal/train/aggregate"
	train "aggregateSimple/internal/train/entity"
	valueObject "aggregateSimple/internal/train/valueObject"
	"fmt"
)

func main() {
	// 使用者
	user := train.CreateUser("jack")

	// 建立站別 value-object
	createStand := valueObject.CreatStand("台北", "台中")

	// 建立持久化 entity
	createSeat := train.CreatSeat(10, createStand)

	// 下單
	order := aggregate.CreateOrder(user, createSeat)
	fmt.Printf("%+v\n", order)

	// 同個座位, 站別改變價錢也改變
	createStand = valueObject.CreatStand("台中", "台北")
	createSeat.ChangeStand(15, createStand)

	// 改單
	order.ChangeOrder(createSeat)
	fmt.Printf("%+v\n", order)
}
