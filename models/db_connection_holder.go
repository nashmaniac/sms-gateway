package models

import "fmt"

type DBConnectionHolder struct {
	Host string
	Password string
	Name string
	Username string
	Port int
	Timezone string
	Sslmode string
}

func (d DBConnectionHolder) GetDSNString() string {
	return fmt.Sprintf("host=%v database=%v user=%v password=%v port=%v sslmode=%v TimeZone=%v",
		d.Host, d.Name, d.Username, d.Password, d.Port, d.Sslmode, d.Timezone)
}
