package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"html"
	"strings"
	"time"
)

type Post struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	PostTitle string 	`gorm:"size:500;" json:"post_title"`
	PostBody  string    `gorm:"size:20000;not null;" json:"post_body"`
	Likes 	  uint		`gorm:"default:0" json:"likes"`
}

func (post *Post) SavePost(db *gorm.DB) (*Post, error) {

	var err error
	err = db.Debug().Create(&post).Error
	if err != nil {
		return &Post{}, err
	}
	return post, nil
}

func (posts *Post) FindAllPostsWithLimit(db *gorm.DB,  offset int, limit int) (*[]Post, error) {
	var err error
	var post []Post
	err = db.Debug().Model(&Post{}).Order("created_at desc").Offset(offset).Limit(limit).Find(&post).Error
	if err != nil {
		return &[]Post{}, err
	}
	return &post, err
}

func (post *Post) SetPostRate(db *gorm.DB, postId int, likes int) (*Post, error) {
	var err error
	err = db.Debug().Model(&Post{}).Where("id = ?", postId).Update("likes", likes).Error
	if err != nil {
		return &Post{}, err
	}
	return post, nil
}

func (p *Post) ToTable() {
	p.CreatedAt = time.Now()
	p.PostTitle = html.EscapeString(strings.TrimSpace(p.PostTitle))
	p.PostBody =  html.EscapeString(strings.TrimSpace(p.PostBody))
}
