package easyvk

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Users struct that provides an acess to users api mthods
type Users struct {
	vk *VK
}

// User struct stores information of user
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Nickname  string `json:"nickname"`
	Sex       int    `json:"sex"`
}

// Get loads the information of users with provided ids
func (u *Users) Get(ids []string, fields []string, nameCase string) ([]User, error) {

	params := map[string]string{}

	if ids != nil && len(ids) > 0 {
		params["user_ids"] = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ids)), ","), "[]")
	}

	if fields != nil && len(fields) > 0 {
		params["fields"] = strings.Join(fields, ",")
	}

	if len(nameCase) > 0 {
		params["name_case"] = nameCase
	}

	resp, err := u.vk.Request("users.get", params)
	if err != nil {
		return nil, err
	}

	var info []User
	err = json.Unmarshal(resp, &info)
	if err != nil {
		return nil, err
	}
	return info, nil
}
