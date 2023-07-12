package main

import (
	"fmt"
	train "valueObjectsSimple/internal/train/entity"
	valueObject "valueObjectsSimple/internal/train/valueObject"
)

func main() {
	// 建立站別 value-object
	createStand := valueObject.CreatStand("台北", "台中")
	createStand = valueObject.CreatStand("台中", "台北")

	// 建立持久化 entity
	createSeat := train.CreatSeat(10, createStand)

	fmt.Println(createSeat)

	// 同個座位, 站別改變價錢也改變
	createSeat.ChangeStand(15, createStand)
	fmt.Println(createSeat)
}
