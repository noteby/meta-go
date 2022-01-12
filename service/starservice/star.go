package starservice

import (
	"errors"
	"log"
	"meta-go/db"
	"meta-go/model"

	"gorm.io/gorm"
)

// 获取star列表
func GetStarList(authorID uint, public bool, offset, limit int) ([]model.Star, int64, error) {
	var stars []model.Star
	var count int64
	star := model.Star{AuthorID: authorID, Public: public}
	if err := db.Conn().Model(&model.Star{}).Where(
		&star,
	).Count(&count).Error; err != nil {
		return stars, count, err
	}
	err := db.Conn().Offset(offset).Limit(limit).Order("created_at desc").Where(
		&star,
	).Find(&stars).Error
	return stars, count, err
}

// 获取star详情
func GetStarDetail(authorID, starId uint) (model.Star, error) {
	var star model.Star
	err := db.Conn().Where(&model.Star{AuthorID: authorID}).First(&star, starId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("记录不存在")
	}
	return star, err
}

// 新增star
func AddStar(authorID uint, public bool, title string, content string) error {
	star := model.Star{
		AuthorID: authorID,
		Title:    title,
		Public:   public,
		Content:  content,
	}
	err := db.Conn().Create(&star).Error
	return err
}

// 更新star
func UpdateStar(authorID, starId uint, public bool, title, content string) error {
	return db.Conn().Transaction(func(tx *gorm.DB) error {
		// 获取star
		var star model.Star
		err := tx.Where(&model.Star{AuthorID: authorID}).First(&star, starId).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(err)
			return errors.New("记录不存在")
		}
		star.Title = title
		star.Public = public
		star.Content = content
		err = tx.Save(&star).Error
		return err
	})
}
