package service

import (
	"github.com/liuhongdi/digv09/cache"
	"github.com/liuhongdi/digv09/dao"
	"github.com/liuhongdi/digv09/model"
)
//得到一篇文章的详情
//得到一篇文章的详情
func GetOneArticle(articleId uint64) (*model.Article, error) {
	//get from bigcache
	article,err := cache.GetOneArticleCache(articleId);
	if ( err != nil) {
		//get from mysql
		article,errSel := dao.SelectOneArticle(articleId);
		if (errSel != nil) {
			return nil,errSel
		} else {
			//set bigcache
			errSet := cache.SetOneArticleCache(articleId,article)
			if (errSet != nil){
				return nil,errSet
			} else {
				return article,errSel
			}
		}
	} else {
		return article,err
	}
}

func GetArticleSum() (int, error) {
	return dao.SelectcountAll()
}

//得到多篇文章，按分页返回
func GetArticleList(page int ,pageSize int) ([]*model.Article,error) {
	articles, err := dao.SelectAllArticle(page,pageSize)
	if err != nil {
		return nil,err
	} else {
		return articles,nil
	}
}