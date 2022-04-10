package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	UserInfo User
	VIPLevel int
	Price    int
}

type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

type Student struct {
	Age   int32
	Score float64
}

func structEx() {
	var house House
	house.Address = "서울시 강동구"
	house.Size = 28
	house.Price = 90000
	house.Type = "아파트"

	fmt.Println("주소: ", house.Address)
	fmt.Printf("크기: %d평\n", house.Size)
	fmt.Printf("가격: %.2f억 원\n", house.Price)
	fmt.Println("타입: ", house.Type)

	user := User{"송하나", "hana", 23}
	vip := VIPUser{User{"화랑", "hwarang", 40}, 3, 250}

	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s 나이: %d VIP 레벨: %d VIP 가격: %d만\n",
		vip.UserInfo.Name,
		vip.UserInfo.ID,
		vip.UserInfo.Age,
		vip.VIPLevel,
		vip.Price,
	)

	student := Student{24, 88}
	fmt.Println(unsafe.Sizeof(student))
}
