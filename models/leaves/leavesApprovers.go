package leaves

import (
	"github.com/ninefive/opms/models"
	"time"

	"github.com/astaxie/beego/orm"
)

type LeavesApprover struct {
	Id      int64 `orm:"pk;column(approverid);"`
	Leaveid int64
	Userid  int64
	Summary string
	Status  int
	Created int64
	Changed int64
}
