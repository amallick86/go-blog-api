package crud

import (
	"errors"
	"go-blog-api/models"
	"go-blog-api/utils/channels"
	"time"

	"github.com/jinzhu/gorm"
)

type repositoryPostCRUD struct {
	db *gorm.DB
}

func NewRepositoryPostsCRUD(db *gorm.DB) *repositoryPostCRUD {
	return &repositoryPostCRUD{db}
}

func (r *repositoryPostCRUD) Save(post models.Post) (models.Post, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&models.Post{}).Create(&post).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return post, nil
	}
	return models.Post{}, err
}

func (r *repositoryPostCRUD) FindAll() ([]models.Post, error) {
	var err error
	posts := []models.Post{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Post{}).Limit(100).Find(&posts).Error
		if err != nil {
			ch <- false
			return
		}
		if len(posts) > 0 {
			for i := range posts {
				err = r.db.Debug().Model(&models.User{}).Where("id = ?", posts[i].AuthorID).Take(&posts[i].Author).Error
				if err != nil {
					ch <- false
					return
				}
			}
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return posts, nil
	}
	return nil, err
}

func (r *repositoryPostCRUD) FindById(pid uint64) (models.Post, error) {
	var err error
	post := models.Post{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Post{}).Where("id = ?", pid).Find(&post).Error
		if err != nil {
			ch <- false
			return
		}
		if post.ID != 0 {
			err = r.db.Debug().Model(&models.User{}).Where("id = ?", post.AuthorID).Take(&post.Author).Error
			if err != nil {
				ch <- false
				return
			}
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return post, nil
	}
	return models.Post{}, err
}

func (r *repositoryPostCRUD) Update(pid uint64, post models.Post) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.Post{}).Where("id = ?", pid).Take(&models.Post{}).UpdateColumns(
			map[string]interface{}{
				"title":      post.Title,
				"content":    post.Content,
				"updated_at": time.Now(),
			},
		)
		ch <- true
	}(done)
	if channels.OK(done) {
		if rs.Error != nil {
			if gorm.IsRecordNotFoundError(rs.Error) {
				return 0, errors.New("post not found")
			}
			return 0, rs.Error
		}
		return rs.RowsAffected, nil
	}
	return 0, rs.Error
}

func (r *repositoryPostCRUD) Delete(pid uint64, uid uint32) (int64, error) {

	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.Post{}).Where("id = ? and author_id=?", pid, uid).Take(&models.Post{}).Delete(&models.Post{})
		ch <- true
	}(done)
	if channels.OK(done) {
		if rs.Error != nil {
			return 0, rs.Error
		}
		return rs.RowsAffected, nil
	}
	return 0, rs.Error
}
