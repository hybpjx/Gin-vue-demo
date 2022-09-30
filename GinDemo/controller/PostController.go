package controller

import (
	"GinDemo/common"
	"GinDemo/model"
	"GinDemo/response"
	"GinDemo/vo"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type IPostController interface {
	RestController
	PageList(ctx *gin.Context)
}

type PostController struct {
	DB *gorm.DB
}

func (p PostController) PageList(ctx *gin.Context) {
	// 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	// 分页
	var posts []model.Post
	p.DB.Order("created desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts)

	// 记录的总条数
	var total int64
	p.DB.Model(model.Post{}).Count(&total)

	response.Success(ctx, gin.H{
		"data":  posts,
		"total": total,
	}, "查询成功")

}

func NewPostController() IPostController {
	db := common.GetDb()
	_ = db.AutoMigrate(model.Post{})
	return PostController{DB: db}
}

func (p PostController) Create(ctx *gin.Context) {

	var requestPost vo.CreatePostRequest

	if err := ctx.ShouldBind(&requestPost); err != nil {
		log.Println(err)
		response.Fail(ctx, gin.H{"数据": requestPost}, "数据验证错误")
		return
	}
	//
	//I8dSWsgFP0
	// 获取登录用户
	user, _ := ctx.Get("userInfo")

	fmt.Println(user.(model.User).ID)
	// 创建文章
	post := model.Post{
		UserID:     user.(model.User).ID,
		CategoryID: requestPost.CategoryID,
		Title:      requestPost.Title,
		HeadImg:    requestPost.HeadImg,
		Content:    requestPost.Content,
	}
	if err := p.DB.Create(&post).Error; err != nil {
		panic(err)
		return
	}
	response.Success(ctx, gin.H{
		"文章": post,
	}, "创建成功")
}

func (p PostController) Delete(ctx *gin.Context) {
	// 获取path的ID
	postID := ctx.Params.ByName("id")

	var post model.Post

	if err := p.DB.First(&post, postID).Error; err != nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	// 判断当前用户是否是文章的作者
	//  获取登录用户
	user, _ := ctx.Get("userInfo")
	userID := user.(model.User).ID
	if userID != post.UserID {
		response.Fail(ctx, nil, "您不是作者,请不要非法操作")
		return
	}

	p.DB.Delete(&post)
	response.Success(ctx, gin.H{"post": post}, "删除成功")
	return

}

func (p PostController) Put(ctx *gin.Context) {
	var requestPost vo.CreatePostRequest

	if err := ctx.ShouldBind(&requestPost); err != nil {
		log.Println("1111")
		response.Fail(ctx, nil, "数据验证错误")
		return
	}

	// 获取path的ID
	postID := ctx.Params.ByName("id")

	var post model.Post

	if err := p.DB.First(&post, postID).Error; err != nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}
	// 判断当前用户是否是文章的作者
	//  获取登录用户
	user, _ := ctx.Get("userInfo")
	userID := user.(model.User).ID
	if userID != post.UserID {
		response.Fail(ctx, nil, "您不是作者,请不要非法操作")
		return
	}

	// 更新文章
	if err := p.DB.Model(&post).Updates(requestPost).Error; err != nil {
		response.Fail(ctx, nil, "更新失败")
		return
	}
	response.Success(ctx, gin.H{"post": post}, "更新成功")
	return
}

func (p PostController) Select(ctx *gin.Context) {
	// 获取path的ID
	postID := ctx.Params.ByName("id")

	var post model.Post

	if err := p.DB.Preload("Category").First(&post, postID).Error; err != nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}
	response.Success(ctx, gin.H{"post": post}, "查询成功")
	return
}
