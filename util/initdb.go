package util

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func InitTables(db  *gorm.DB,tname string,fslice  []string) {

	if db.HasTable(tname) == false{
		// 不存在表就新建表
		CreateSql := "CREATE TABLE "+ tname+ " ( `id` bigint(25)  unsigned NOT NULL AUTO_INCREMENT, imei varchar(32) COLLATE utf8_bin DEFAULT NULL,creative datetime DEFAULT CURRENT_TIMESTAMP, "
		End_sql :="PRIMARY KEY (id)) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin"
		for  _,v  :=range fslice{
			CreateSql = CreateSql + v + " decimal(10,2) DEFAULT NULL,"
		}
		CreateSql =  CreateSql + End_sql

		fmt.Println("Create_Sql:",CreateSql)
		err  :=db.Exec(CreateSql).Error
		if err  !=nil{
			fmt.Println("创建数据出现了问题")
		}
	}else{
		//当表存在的时候
		for   _,v  := range  fslice{
			var  addItem = "ALTER TABLE  "+ tname +" ADD COLUMN  "+v+" DECIMAL(10) NULL"
			db.Exec(addItem)
		}
	}
}

