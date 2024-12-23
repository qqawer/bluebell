package mysql

import (
	"WebApp/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// CreatePost 创建帖子
func CreatePost(p *models.Post) error {
	if err := db.Create(&p).Error; err != nil {
		return err
	}
	return nil
}

// GetPostById 根据id查询单个帖子数据
func GetPostById(pid int64) (*models.Post, error) {
	var post models.Post
	if err := db.Where("post_id=?", pid).First(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("cannot find data with this pid")
		} else {
			return nil, err
		}
	}
	return &post, nil
}

// GetPostList 查询帖子列表函数
func GetPostList(page, size int) (data []*models.Post, err error) {
	data = make([]*models.Post, 0, size)
	offset := (page - 1) * size
	if err := db.Offset(offset).Limit(size).Find(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("没有找到")
		} else {
			return nil, err
		}
	}
	return data, nil

}

// GetPostListByIDs 根据给定的id列表查询帖子
/*
func GetPostListByIDs(ids []string) ([]*models.Post, error) {
	var posts []*models.Post
	if err := db.Where("post_id IN ?", ids).
		Order("post_id ASC"). // 按照post_id升序排序，使用DESC则为降序
		Find(&posts).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("没有找到相关帖子")
		}
		return nil, err
	}
	return posts, nil
}
*/
func GetPostListByIDs(ids []string) ([]*models.Post, error) {
	var posts []*models.Post

	// 构建CASE语句以按照ids的顺序排序
	caseStatement := "CASE"
	for i, id := range ids {
		caseStatement += fmt.Sprintf(" WHEN post_id = '%s' THEN %d", id, i)
	}
	caseStatement += " END"

	if err := db.Where("post_id IN ?", ids).
		Order(caseStatement). // 按照ids的顺序排序
		Find(&posts).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("没有找到相关帖子")
		}
		return nil, err
	}
	return posts, nil
}

// GetPostListByIDs 根据给定的id列表查询帖子数据
/*func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	if len(ids) == 0 {
		return nil, errors.New("ids cannot be empty")
	}

	// 使用 GORM 自定义 SQL 表达式实现 FIND_IN_SET
	// err = db.Where("post_id IN ?", ids).
	//     Order(gorm.Expr("FIND_IN_SET(post_id, ?)", strings.Join(ids, ","))). // 按给定的 ids 顺序排序
	//     Find(&postList).Error
	// if err != nil {
	//     return nil, err
	// }
	orderClause := "CASE"
	for i, id := range ids {
		orderClause += fmt.Sprintf(" WHEN post_id = '%s' THEN %d", id, i)
	}
	orderClause += " END"

	err = db.Where("post_id IN ?", ids).
		Order(gorm.Expr(orderClause)).
		Find(&postList).Error

	return postList, nil
}
*/
