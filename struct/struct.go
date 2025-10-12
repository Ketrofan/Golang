package main

import "fmt"

type User struct {
	Name    string
	Age     int
	Phone   string
	IsClose bool
	Rating  float64
}

//Констурктор с валидацией(инкапсулирование/ограничение для других)) они не смогут сделать занчения больше положенных
//Но при вводе напрямую user1.Age =200 - пройдет. НО ЕСТЬ возмодность ограничений к переменной и полю Age, но я пока не знаю
func NewUser(
	name string,
	age int,
	phone string,
	isClose bool,
	rating float64,
) User {
	fmt.Println("Валидирую имя")
	if name == "" {
		fmt.Println("Не прошел валидацию")
		return User{}
	}
	if age < 0 || age > 150 {
		return User{}
	}
	if phone == "" || len(phone) != 12 {
		return User{}
	}
	if rating > 10.0 || rating < 0.0 {
		return User{}
	}
	return User{
		Name:    name,
		Age:     age,
		Phone:   phone,
		IsClose: isClose,
		Rating:  rating,
	}
}
func (u *User) ChangeName(newName string) {
	if newName != "" {
		u.Name = newName
	}
}

func (u *User) ChangeAge(newAge int) {
	if newAge > 0 && newAge < 150 {
		u.Age = newAge
	}
}

func (u *User) ChangePhone(newPhone string) {
	if newPhone != "" && len(newPhone) == 12 {
		u.Phone = newPhone
	}
}

func (u *User) ChangeIsClose(newIsClose bool) {
	u.IsClose = newIsClose
}

func (u *User) ChangeRatingUp(ratingUp float64) {
	if u.Rating+ratingUp <= 10.0 {
		u.Rating += ratingUp
	}
}

func (u *User) ChangeRatingDown(ratingDown float64) {
	if u.Rating-ratingDown >= 0.0 {
		u.Rating -= ratingDown
	}
	fmt.Println("Все говно не пошло")
}
func main() {
	user1 := NewUser(
		"John Doe",
		30,
		"+12345678901",
		false,
		5.0,
	)
	fmt.Println("Before:", user1)
	user1.ChangeName("Andy")
	fmt.Println("After:", user1)

	user2 := NewUser(
		"asd",
		10,
		"+55555555555",
		true,
		8.2,
	)
	fmt.Println("Before:", user2)
	user1.ChangeName("ABikla")
	fmt.Println("After:", user2)
	fmt.Println("After:", user1)

}
