package user

import "fmt"

func New() *User {
	return new(User)
}

type User struct {
	Username string
	Password string
	APIToken string
	Name     string `json:"name"`
	Email    string `json:"email"`
	Initials string `json:"initials"`
	Timezone struct {
		Kind      string `json:"kind"`
		Offset    string `json:"offset"`
		OlsonName string `json:"olson_name"`
	} `json:"time_zone"`
}

func (u *User) Login(name, pass string) {
	u.Username = name
	u.Password = pass
}

func (u *User) Print() {
	fmt.Printf("uname: %s\n", u.Username)
	fmt.Printf("name: %s\n", u.Name)
	fmt.Printf("email: %s\n", u.Email)
	fmt.Printf("initials: %s\n", u.Initials)
	fmt.Printf("time_zone: %v\n", u.Timezone)
}
