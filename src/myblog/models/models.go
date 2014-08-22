package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"strconv"
	"time"
)

const (
	_DB_NAME        = "data/beelog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

// 分类
type Category struct {
	Id              int64
	Title           string    // 标题
	Created         time.Time `orm:"index"` // 创建时间
	Views           int64     `orm:"index"` // 浏览次数
	TopicTime       time.Time `orm:"index"` // 发表时间
	TopicCount      int64     // 文章数目
	TopicLastUserId int64     // 最后操作者
}

// 文章
type Topic struct {
	Id              int64
	UserId          int64     // 作者
	Title           string    // 标题
	Content         string    `orm:"size(5000)"` // 内容
	Attachment      string    // 附件
	Created         time.Time `orm:"index"` // 创建时间
	Updated         time.Time `orm:"index"` // 更新时间
	Views           int64     `orm:"index"` // 浏览次数
	Author          string    // 作者
	ReplyTime       time.Time `orm:"index"` // 回复时间
	ReplyCount      int64     // 回复数目
	ReplyLastUserId int64
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DR_Sqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

func AddCategory(name string) error {
	o := orm.NewOrm()
	cate := &Category{Title: name, Created: time.Now(), TopicTime: time.Now()}

	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)

	if err == nil {
		return err
	}

	_, err = o.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}

func GetAllCategorys() ([]*Category, error) {
	o := orm.NewOrm()

	cates := make([]*Category, 0)

	qs := o.QueryTable("category")

	_, err := qs.All(&cates)

	return cates, err
}

func DeleteCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	cate := &Category{Id: cid}

	_, err = o.Delete(cate)
	return err
}
