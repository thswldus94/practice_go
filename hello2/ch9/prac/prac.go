package main

import "fmt"

func main() {
	calcAge()
	calcWeather()
}

func calcAge() {
	age := 29

	if age < 10 {
		fmt.Println("아직 애 입니돠")
	} else if age >= 20 && age < 30 {
		fmt.Println("서른 얼마안남음 ㅠ")
	} else {
		fmt.Println("나머지는 대충 아무거나")
	}
}

func calcWeather() {
	temp := 30
	rain := 40

	if temp >= 25 {
		if rain >= 80 {
			fmt.Println("덥고 비가 옵니다")
		} else if rain >= 20 {
			fmt.Println("덥고 습합니다")
		} else {
			fmt.Println("야외활동 하기 좋습니다")
		}
	} else if temp < 10 && rain >= 80 {
		fmt.Println("야외활동 하기 좋지 않습니다")
	} else {
		fmt.Println("좋은 날씨입니다")
	}
}
