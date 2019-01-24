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

type City struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type Country struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// User struct stores information of user
type User struct {
	ID        uint64  `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Sex       uint8   `json:"sex"`
	Nickname  string  `json:"nickname"`
	Domain    string  `json:"domain"`
	City      City    `json:"city"`
	Country   Country `json:"country"`
	Hometown  string  `json:"home_town"`
	Status    string  `json:"status"`
	Bdate     string  `json:"bdate"`
	Interests string  `json:"interests"`
	Relation  uint8   `json:"relation"`
}

type UsersSearchResults struct {
	Count int    `json:"count"`
	Items []User `json:"items"`
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

	var users []User
	err = json.Unmarshal(resp, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Search returns search results
func (u *Users) Search(q string, params map[string]string) (*UsersSearchResults, error) {
	params["q"] = q

	resp, err := u.vk.Request("users.search", params)
	if err != nil {
		return nil, err
	}

	var searchResults UsersSearchResults
	err = json.Unmarshal(resp, &searchResults)
	if err != nil {
		return nil, err
	}

	return &searchResults, nil
}
