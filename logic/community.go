package logic

import (
	"WebApp/dao/mysql"
	"WebApp/models"
)

// func GetCommunityList(input *models.Community)([]models.Community,error) {
// 	return mysql.GetCommunityList(input)
// }
func GetCommunityList()([]models.Community,error) {
	return mysql.GetCommunityList()
}
func GetCommunitDetail(id int64)(*models.CommunityDetail,error){
	return mysql.GetCommunityDetailByID(id)
}