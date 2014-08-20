package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
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
