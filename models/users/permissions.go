package users

import (
	"fmt"
	"github.com/ninefive/opms/models"
	"github.com/ninefive/opms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Permissions struct {
	Id         int64 `orm:"pk;column(userid);"`
	Permission string
	Model      string
	Modelc     string
}

func (this *Permissions) TableName() string {
	return models.TableName("permissions")
}

func init() {
	orm.RegisterModel(new(Permissions))
}

func GetPermissions(id int64) string {
	var err error
	var name string
	err = utils.GetCache("GetPermissions.id."+fmt.Sprintf("%d", id), &name)
	if err != nil {
		cache_expire, _ := beego.AppConfig.Int("cache_expire")
		var permission Permissions
		o := orm.NewOrm()
		o.QueryTable(models.TableName("permissions")).Filter("userid", id).One(&permission, "permission")
		name = permission.Permission
		utils.SetCache("GetPermissions.id."+fmt.Sprintf("%d", id), name, cache_expire)
	}
	return name
}

func GetPermissionsAll(id int64) (Permissions, error) {
	var per Permissions
	var err error
	o := orm.NewOrm()

	per = Permissions{Id: id}
	err = o.Read(&per)

	if err == orm.ErrNoRows {
		return per, nil
	}
	return per, err
}

func AddPermissions(updPer Permissions) error {
	o := orm.NewOrm()
	o.Using("default")
	per := new(Permissions)

	per.Id = updPer.Id
	per.Permission = updPer.Permission
	per.Model = updPer.Model
	per.Modelc = updPer.Modelc
	_, err := o.Insert(per)
	return err
}

func UpdatePermissions(id int64, updPer Permissions) error {
	var per Permissions
	o := orm.NewOrm()
	per = Permissions{Id: id}

	per.Permission = updPer.Permission
	per.Model = updPer.Model
	per.Modelc = updPer.Modelc
	_, err := o.Update(&per, "permission", "model", "modelc")
	return err
}
