package models

import (
	"fmt"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	dbu := beego.AppConfig.String("DB::user")
	dbp := beego.AppConfig.String("DB::pwd")
	dbn := beego.AppConfig.String("DB::name")

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8", dbu, dbp, dbn), 10, 10)
	orm.RegisterModel(new(Task))
	_ = orm.RunSyncdb("default", false, true)

}
