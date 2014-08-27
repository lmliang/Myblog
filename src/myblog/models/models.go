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
	Category        string
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

// 评论
type Comment struct {
	Id      int64
	Tid     int64
	Name    string
	Content string    `orm:size(1000)`
	Created time.Time `orm:index`
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	orm.RegisterModel(new(Category), new(Topic), new(Comment))
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DR_Sqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

func AddCategory(name string) error {
	o := orm.NewOrm()
	cate := &Category{Title: name}

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

func AddTopic(title, content, category string) error {
	o := orm.NewOrm()

	topic := &Topic{Title: title, Content: content, Category: category, Created: time.Now(), Updated: time.Now()}

	_, err := o.Insert(topic)
	if err != nil {
		return err
	}

	cate := &Category{Title: category}

	qs := o.QueryTable("category")

	err = qs.Filter("Title", category).One(cate)
	if err != nil {
		return err
	}

	cate.TopicCount++
	cate.TopicTime = topic.Created

	_, err = o.Update(cate)
	return err
}

func GetAllTopics(category string, isDesc bool) ([]*Topic, error) {
	topics := make([]*Topic, 0)

	o := orm.NewOrm()

	qs := o.QueryTable("topic")

	var err error
	if isDesc {
		if len(category) > 0 {
			qs = qs.Filter("category", category)
		}
		_, err = qs.OrderBy("-id").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}

	return topics, err
}

func GetTopic(id string) (*Topic, error) {
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	t := new(Topic)

	o := orm.NewOrm()
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tid).One(t)
	if err != nil {
		return nil, err
	}

	t.Views++
	_, err = o.Update(t)

	return t, err
}

func ModifyTopic(id, title, content, category string) error {
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	t := &Topic{Id: tid}

	o := orm.NewOrm()
	if o.Read(t) == nil {
		t.Title = title
		t.Content = content
		t.Category = category
		t.Updated = time.Now()
		_, err = o.Update(t)
	}

	return err
}

func DeleteTopic(id string) error {
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	topic := &Topic{Id: tid}

	err = o.Read(topic)
	if err != nil {
		return err
	}

	category := topic.Category

	_, err = o.Delete(topic)

	cate := &Category{Title: category}

	qs := o.QueryTable("category")

	err = qs.Filter("Title", category).One(cate)
	if err != nil {
		return err
	}

	if cate.TopicCount > 0 {
		cate.TopicCount--
	}

	_, err = o.Update(cate)

	return err
}

func AddReply(tid, nickname, content string) error {
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	reply := &Comment{Tid: id, Name: nickname, Content: content, Created: time.Now()}

	o := orm.NewOrm()
	_, err = o.Insert(reply)
	return err
}

func GetTopicComments(tid string) ([]*Comment, error) {
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	replys := make([]*Comment, 0)

	o := orm.NewOrm()

	qs := o.QueryTable("comment")

	_, err = qs.Filter("tid", id).All(&replys)

	return replys, err
}

func DeleteReply(rid string) error {
	id, err := strconv.ParseInt(rid, 10, 64)
	if err != nil {
		return err
	}

	r := &Comment{Id: id}

	o := orm.NewOrm()

	_, err = o.Delete(r)

	return err
}
