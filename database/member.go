package database

import (
	"fmt"

	"github.com/fatih/color"
)

type Member struct {
	Email  string `json:"email"`
	Passwd string `json:"passwd"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Tel    string `json:"tel"`
}

func InitMember() *Member {
	member := &Member{}

	return member
}

func (m *Member) PrintMember() {
	whilte := color.New(color.FgWhite)
	boldWhite := whilte.Add(color.BgGreen)
	boldWhite.Printf("email : %s", m.Email)
	fmt.Println()
	boldWhite.Printf("passwd : %s", m.Passwd)
	fmt.Println()
}
