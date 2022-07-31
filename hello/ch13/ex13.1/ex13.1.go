package main

import "fmt"

type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

func main() {
	var house House
	house.Address = "경기도 성남시 분당구"
	house.Size = 20
	house.Price = 7.5
	house.Type = "APT"

	fmt.Println("주소: ", house.Address)
	fmt.Printf("크기: %d평\n", house.Size)
	fmt.Printf("가격: %.2f억 원\n", house.Price)
	fmt.Println("타입: ", house.Type)
}
