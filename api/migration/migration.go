package migration

import (
	"awesomeProject/dao"
	"github.com/jinzhu/gorm"
	"log"
)

var posts =
	dao.Post{
		PostBody: "Hello world 1",
		PostTitle: "Tbl Pidr",
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&dao.Post{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&dao.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&dao.Post{}).Create(&posts).Error
	if err != nil {
		log.Fatalf("cannot seed posts table: %v", err)
	}

}
