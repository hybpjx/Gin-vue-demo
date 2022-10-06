package controller

import (
	"GinDemo/common"
	"GinDemo/model"
	"GinDemo/repository"
	"GinDemo/response"
	"GinDemo/vo"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ICategoryController interface {
	RestController
	SelectAll(ctx *gin.Context)
}

type CategoryController struct {
	//DB gorm.DB
	Repository repository.CategoryRepository
}

func NewCategoryController() ICategoryController {

	_repository := repository.NewCateGoryRepository()

	_ = _repository.DB.AutoMigrate(model.Category{})

	return CategoryController{
		Repository: _repository,
	}
}

// Create 创建某个分类
func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory vo.CreateCategoryRequest

	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, nil, "数据验证错误,分类名称必填项")
		return
	}

	category, err := c.Repository.Create(requestCategory.Name)
	if err != nil {
		//response.Fail(ctx, nil, "创建失败")
		// 写拦截 专门针对这个 同名的错误
		panic(err)
		return
	}

	response.Success(ctx, gin.H{"category": category}, "返回成功")
	return
}

// Delete 删除某个分类
func (c CategoryController) Delete(ctx *gin.Context) {
	// 获取path中的参数
	categoryID, _ := strconv.Atoi(ctx.Params.ByName("id"))

	category, err := c.Repository.SelectByID(categoryID)
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}
	err = c.Repository.DeleteByID(categoryID)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, nil, "删除失败")
		return
	}

	response.Success(ctx, gin.H{
		"category": category,
	}, "删除成功")
	return
}

// Put 更新某个分类
func (c CategoryController) Put(ctx *gin.Context) {
	// 绑定body中的参数
	var requestCategory vo.CreateCategoryRequest

	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, nil, "数据验证错误,分类名称必填项")
		return
	}
	// 获取path中的参数
	categoryID, _ := strconv.Atoi(ctx.Params.ByName("id"))

	updateCategory, err := c.Repository.SelectByID(categoryID)
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	category, err := c.Repository.Update(*updateCategory, requestCategory.Name)
	if err != nil {
		response.Fail(ctx, nil, "修改失败，有同名分类")
		return
	}
	response.Success(ctx, gin.H{"category": category}, "修改成功")
	return
}

// Select 查询单个ID
func (c CategoryController) Select(ctx *gin.Context) {
	// 获取path中的参数
	categoryID, _ := strconv.Atoi(ctx.Params.ByName("id"))

	SelectCategory, err := c.Repository.SelectByID(categoryID)
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	response.Success(ctx, gin.H{"SelectCategory": SelectCategory}, "查询成功")
	return
}

func (c CategoryController)SelectAll(ctx *gin.Context){
	db := common.DB

	// 种类列表
	var categories []model.Category

	db.Order("created_at desc").Offset(0).Limit(15).Find(&categories)
	fmt.Println(categories)

	// 记录种类列表的总条数
	var total int64
	db.Model(model.Category{}).Count(&total)

	response.Success(ctx, gin.H{
		"data":  categories,
		"total": total,
	}, "查询种类成功")

}