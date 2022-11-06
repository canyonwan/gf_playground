// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"time"
)

// User is the golang structure for table user.
type User struct {
	Id       uint      `json:"id"       description:"User ID"`       // User ID
	Passport string    `json:"passport" description:"User Passport"` // User Passport
	Password string    `json:"password" description:"User Password"` // User Password
	Nickname string    `json:"nickname" description:"User Nickname"` // User Nickname
	CreateAt time.Time `json:"createAt" description:"Created Time"`  // Created Time
	UpdateAt time.Time `json:"updateAt" description:"Updated Time"`  // Updated Time
}