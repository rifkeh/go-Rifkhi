package main

import "fmt"

type Car struct {
	typeCar string
	fuelIn  int
}

func (c Car) maxDistance() (distance float32) {
	distance = float32(c.fuelIn) / 1.5
	return distance
}

func main() {
	lambo := Car{"mahal", 300}
	fmt.Println(lambo.maxDistance())

}
