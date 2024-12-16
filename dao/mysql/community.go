package mysql

import (
	"WebApp/models"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// func GetCommunityList(input *models.Community)([]models.Community,error){
// 	var community []models.Community
// 	err := db.Where("community_id = ? AND community_name = ?", input.ID, input.Name).First(&community).Error
//     if err != nil {
//         if errors.Is(err, gorm.ErrRecordNotFound) {
// 			zap.L().Warn("there is no community")
//             return nil, errors.New("没有查询到community") // 如果没有找到记录，返回nil和nil错误
//         }
//         return nil, err // 返回查询过程中出现的错误
//     }

//	    return community, nil // 返回查询结果和nil错误
//	}
func GetCommunityList() ([]models.Community, error) {
	var community []models.Community
	err := db.Find(&community).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zap.L().Warn("there is no community")
			return nil, errors.New("没有查询到community") // 如果没有找到记录，返回nil和err错误
		}
		return nil, err // 返回查询过程中出现的错误
	}

	return community, nil // 返回查询结果和nil错误
}
func GetCommunityDetailByID(id int64) (*models.CommunityDetail, error) {
	var community models.CommunityDetail
	if err := db.First(&community, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("没有查询到community")
		} else {
			return nil, err
		}
	}
	return &community,nil
}
