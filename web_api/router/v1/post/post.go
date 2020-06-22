package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	"github.com/CapsLock-Studio/neo-komica/model"
)

type indexRequest struct {
	TopicCode    string `form:"topic" binding:"required"`
	Size         uint   `form:"size,default=20"`
	FromPostUUID string `form:"from,defualt=''"`
}

// Index - list all posts with topic ID
// @Tags Post
// @Summary list all posts with topic ID
// @Accept json
// @Produce json
// @Param topic query string true "post's topic"
// @Param size query uint false "specify list size"
// @Param from query string false "specify list from"
// @Success 200 {array} model.Post
// @Router /post [get]
func Index(ctx *gin.Context) {
	var req indexRequest
	if err := ctx.BindQuery(&req); err != nil {
		return
	}

	var topic model.Topic
	model.SharedDB.Select("id").Where(&model.Topic{Code: req.TopicCode}).First(&topic)

	var fromPost model.Post
	model.SharedDB.Select("id").Where("uuid = ?", uuid.FromStringOrNil(req.FromPostUUID)).
		Assign(model.Post{PublicModel: model.PublicModel{ID: 0}}).
		First(&fromPost)

	var posts []model.Post
	model.SharedDB.Where("topic_id = ? AND parent_id IS NULL AND id > ?", topic.ID, fromPost.ID).
		Find(&posts).
		Limit(req.Size)

	ctx.JSON(http.StatusOK, posts)
}

type moreRequest struct {
	Size         uint   `form:"size,default=20"`
	FromPostUUID string `form:"from,defualt=''"`
}

// More - get reply posts of the post
// @Tags Post
// @Summary get reply posts
// @Accept json
// @Produce json
// @Param parent path string false "get reply posts"
// @Param size query uint false "specify list size"
// @Param from query string false "specify list from"
// @Success 200 {array} model.Post
// @Router /post/{parent}/more [get]
func More(ctx *gin.Context) {
	parentUUID := ctx.Param("parent")
	var req moreRequest
	if err := ctx.BindQuery(&req); err != nil {
		return
	}

	var fromPost model.Post
	model.SharedDB.Select("id").Where("uuid = ?", uuid.FromStringOrNil(req.FromPostUUID)).
		Assign(model.Post{PublicModel: model.PublicModel{ID: 0}}).
		First(&fromPost).
		Limit(req.Size)

	var parent model.Post
	model.SharedDB.Select("id").Where("uuid = ?", parentUUID).
		Find(&parent)

	var posts []model.Post
	model.SharedDB.
		Where("id > ? AND parent_id = ?", fromPost.ID, parent.ID).
		Find(&posts)

	ctx.JSON(http.StatusOK, posts)
}
