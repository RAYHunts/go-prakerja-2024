package models

type SocialMedia struct {
	Id             uint   `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name           string `json:"name" gorm:"column:name"`
	SocialMediaUrl string `json:"social_media_url" gorm:"column:social_media_url"`
	UserId         uint   `json:"user_id" gorm:"column:user_id"`
}
