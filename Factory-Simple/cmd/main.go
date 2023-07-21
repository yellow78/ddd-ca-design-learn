package main

import (
	aggregate "factorySimple/internal/train/aggregate"
	"fmt"
)

func main() {
	// 創建車次
	car := aggregate.CreateCard(1, 30, "台中", "台北")
	fmt.Printf("%+v\n", car)
}
