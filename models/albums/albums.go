package albums

import (
	"github.com/ninefive/opms/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Albums struct {
	Id       int64 `orm:"pk;column(albumid);"`
	Userid   int64
	Title    string
	Picture  string
	Keywords string
	Summary  string
	Created  int64
	Viewnum  int
	Comtnum  int
	Laudum   int
	Status   int
}

func (this *Albums) TableName() string {
	return models.TableName("albums")
}

func init() {
	orm.RegisterModel(new(Albums))
}

/*
* 获取相册详情
 */
