package easyvk

import (
	"encoding/json"
	"fmt"
	"strings"
)

// A Board describes a set of methods to work with topics.
// https://vk.com/dev/board
type Users struct {
	vk *VK
}

type UserInfoResponse struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// AddTopic creates a new topic on a community's discussion board.
// https://vk.com/dev/board.addTopic
func (b *Users) Get(ids []int, fields []string, nameCase string) ([]UserInfoResponse, error) {
	params := map[string]string{
		"user_ids":  strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ids)), ","), "[]"),
		"fields":    strings.Join(fields, ","),
		"name_case": nameCase,
	}
	resp, err := b.vk.Request("users.get", params)
	if err != nil {
		return nil, err
	}

	var info []UserInfoResponse
	err = json.Unmarshal(resp, &info)
	if err != nil {
		return nil, err
	}
	return info, nil
}
