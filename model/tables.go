package model

type User struct {
	Username    string `gorm:"size:64;primary_key;"       json:"username"     form:"username"`
	Email       string `gorm:"size:128;unique;not null;"  json:"email"        form:"email"`
	Password    string `gorm:"size:512;not null;"         json:"password"     form:"password"`
	Nickname    string `gorm:"size:16;"                   json:"nickname"     form:"nickname"`
	Avatar      string `gorm:"type:text"                  json:"avatar"       form:"avatar"`
	HomePage    string `gorm:"size:256"                   json:"home_page"    form:"home_page"`
	Description string `gorm:"type:text"                  json:"description"  form:"description"`
	Status      string `gorm:"size:32;not null;"          json:"status"       form:"status"`
	Role        string `gorm:"size:32;not null;"          json:"role"         form:"role"`
	CreatedAt   int64  `gorm:"not null"                   json:"created_at"   form:"created_at"`
}

type Token struct {
	Token     string `gorm:"size:64;primary_key;"       json:"token"       form:"token"`
	Username  string `gorm:"size:64;"                   json:"username"    form:"username"`
	Status    string `gorm:"size:32"                    json:"status"      form:"status"`
	CreatedAt int64  `gorm:"not null"                   json:"created_at"  form:"created_at"`
	ExpiredAt int64  `gorm:"not null"                   json:"expired_at"  form:"expired_at"`
}

type Config struct {
	ID      uint   `gorm:"primary_key"`
	Name    string `gorm:"size:32;"        json:"name"     form:"name"`
	Value   string `gorm:"size:64;"        json:"value"    form:"value"`
	Section string `gorm:"size:64;"        json:"section"  form:"section"`
}

type RuningInfo struct {
	ID      uint   `gorm:"primary_key"`
	Name    string `gorm:"size:32;"        json:"name"     form:"name"`
	Value   string `gorm:"size:64;"        json:"value"    form:"value"`
	Section string `gorm:"size:64;"        json:"section"  form:"section"`
}

type Article struct {
	ID        uint `gorm:"primary_key"`
	Title     string
	Content   string `gorm:"type:text"`
	Author    string `gorm:"size:64;"                   json:"author"    form:"author"`
	Status    string `gorm:"size:32"                    json:"status"      form:"status"`
	CreatedAt int64  `gorm:"not null"                   json:"created_at"  form:"created_at"`
	ExpiredAt int64  `gorm:"not null"                   json:"expired_at"  form:"expired_at"`
}

type Comment struct {
	// TODO: 评论 or 回复, 父子关系
	ID        uint   `gorm:"primary_key"`
	Author    string `gorm:"size:64;"                   json:"author"    form:"author"`
	ArticleID uint
	Email     string `gorm:"size:128;not null;"         json:"email"        form:"email"`
	Status    string `gorm:"size:32"                    json:"status"      form:"status"`
	CreatedAt int64  `gorm:"not null"                   json:"created_at"  form:"created_at"`
	ExpiredAt int64  `gorm:"not null"                   json:"expired_at"  form:"expired_at"`
	Content   string `gorm:"type:text"`
}
