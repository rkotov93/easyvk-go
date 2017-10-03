package easyvk

import (
	"fmt"
	"encoding/json"
)

// A Likes describes a set of methods to work with likes.
// https://vk.com/dev/likes
type Likes struct {
	vk *VK
}

type likeType string

const (
	// PostLikeType - post on user or community wall
	PostLikeType likeType = "post"
	// CommentLikeType - comment on a wall post
	CommentLikeType likeType = "comment"
	// PhotoLikeType - photo
	PhotoLikeType likeType = "photo"
	// AudioLikeType - audio
	AudioLikeType likeType = "audio"
	// VideoLikeType - video
	VideoLikeType likeType = "video"
	// NoteLikeType - note
	NoteLikeType likeType = "note"
	// MarketLikeType - market
	MarketLikeType likeType = "market"
	// PhotoCommentLikeType - comment on the photo
	PhotoCommentLikeType likeType = "photo_comment"
	// VideoCommentLikeType - comment on the video
	VideoCommentLikeType likeType = "video_comment"
	// TopicCommentLikeType - comment in the discussion
	TopicCommentLikeType likeType = "topic_comment"
	// MarketCommentLikeType - comment on the market
	MarketCommentLikeType likeType = "market_comment"
)

// Add adds the specified object to the Likes list of the current user.
// https://vk.com/dev/likes.add
func (l *Likes) Add(t likeType, ownerID int, itemID uint, accessKey string) (int, error) {
	params := map[string]string{
		"type":       string(t),
		"owner_id":   fmt.Sprint(ownerID),
		"item_id":    fmt.Sprint(itemID),
		"access_key": accessKey,
	}
	resp, err := l.vk.Request("likes.add", params)
	if err != nil {
		return 0, err
	}
	var likes struct {
		Likes int `json:"likes"`
	}
	err = json.Unmarshal(resp, &likes)
	if err != nil {
		return 0, err
	}
	return likes.Likes, nil
}

// Delete deletes the specified object from the Likes list of the current user.
// https://vk.com/dev/likes.delete
func (l *Likes) Delete(t likeType, ownerID int, itemID uint) (int, error) {
	params := map[string]string{
		"type":     string(t),
		"owner_id": fmt.Sprint(ownerID),
		"item_id":  fmt.Sprint(itemID),
	}
	resp, err := l.vk.Request("likes.delete", params)
	if err != nil {
		return 0, err
	}
	var likes struct {
		Likes int `json:"likes"`
	}
	err = json.Unmarshal(resp, &likes)
	if err != nil {
		return 0, err
	}
	return likes.Likes, nil
}

// LikesIsLikedResponse describes info about like and repost.
// https://vk.com/dev/likes.isLiked
type LikesIsLikedResponse struct {
	Liked  bool
	Copied bool
}

// IsLiked checks for the object in the Likes list of the specified user.
// https://vk.com/dev/likes.isLiked
func (l *Likes) IsLiked(userID uint, t likeType, ownerID int, itemID uint) (LikesIsLikedResponse, error) {
	params := map[string]string{
		"type":     string(t),
		"owner_id": fmt.Sprint(ownerID),
		"item_id":  fmt.Sprint(itemID),
		"user_id":  fmt.Sprint(userID),
	}
	resp, err := l.vk.Request("likes.isLiked", params)
	if err != nil {
		return LikesIsLikedResponse{}, err
	}
	var r struct {
		Liked  int `json:"liked"`
		Copied int `json:"copied"`
	}
	err = json.Unmarshal(resp, &r)
	if err != nil {
		return LikesIsLikedResponse{}, err
	}
	response := LikesIsLikedResponse{
		r.Liked == 1,
		r.Copied == 1,
	}
	return response, nil
}

// LikesGetListParams provides struct for getList parameters.
// https://vk.com/dev/likes.getList
type LikesGetListParams struct {
	Type        likeType
	OwnerID     int
	ItemID      int
	PageURL     string
	Filter      string
	FriendsOnly bool
	Offset      uint
	Count       uint
	SkipOwner   bool
}

// LikesGetListResponse provides structure for getList response.
// https://vk.com/dev/likes.getList
type LikesGetListResponse struct {
	Count int `json:"count"`
	Items []UserObject
}

// GetList returns a list of IDs of users who added the specified object to their Likes list.
// https://vk.com/dev/likes.getList
func (l *Likes) GetList(params LikesGetListParams) (LikesGetListResponse, error) {
	p := map[string]string{
		"type":         string(params.Type),
		"owner_id":     fmt.Sprint(params.OwnerID),
		"item_id":      fmt.Sprint(params.ItemID),
		"page_url":     params.PageURL,
		"filter":       params.Filter,
		"friends_only": boolConverter(params.FriendsOnly),
		"extended":     "1",
		"offset":       fmt.Sprint(params.Offset),
		"count":        fmt.Sprint(params.Count),
		"skip_own":     boolConverter(params.SkipOwner),
	}
	resp, err := l.vk.Request("likes.getList", p)
	if err != nil {
		return LikesGetListResponse{}, err
	}
	var response LikesGetListResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return LikesGetListResponse{}, err
	}
	return response, nil
}
