package entities

type User struct {
	ID      int           `json:"id,omitempty" gorm:"primaryKey"`
	Name    string        `json:"name,omitempty"`
	Age     int           `json:"age,omitempty"`
	Status  int           `json:"status,omitempty"`
	Bio     string        `json:"bio,omitempty"`
	Avatars []*UserAvatar `json:"avatars,omitempty"`
}

func (*User) TableName() string {
	return "users"
}

type UserAvatar struct {
	Image    string `json:"image,omitempty"`
	Position int    `json:"position,omitempty"`
}

func (*UserAvatar) TableName() string {
	return "user_avatars"
}
