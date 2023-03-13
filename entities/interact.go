package entities

import "time"

type Like struct {
	ID        int        `json:"id,omitempty" gorm:"primaryKey"`
	Liker     *User      `json:"liker,omitempty"`
	Likee     *User      `json:"likee,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

type Match struct {
	ID        int        `json:"id,omitempty" gorm:"primaryKey"`
	Matchers  []*User    `json:"matchers,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}
