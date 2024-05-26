package test

import (
	"fmt"
	"godevops/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMysqlByDefault(host, port, user, pass, dbname string) (*gorm.DB, error) {
	// user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, dbname)
	return gorm.Open(mysql.Open(dns), &gorm.Config{})
}

func GormAutoMigrate(host, port, use, pass, database string) error {
	// ConnectMysqlByDefault 代码参见上面的: 默认连接
	mysqlByDefault, err := ConnectMysqlByDefault(host, port, use, pass, database)
	if err != nil {
		return err
	}
	// 指定引擎和表备注
	err = mysqlByDefault.Set("gorm:table_options", "ENGINE=InnoDB COMMENT='用户表'").AutoMigrate(&models.UserBasic{})
	if err != nil {
		return err
	}
	return nil
}

// 测试迁移
func TestAutoMigrate() {
	host := "162.14.100.125"
	use, pass, port, database := "root", "zhaoCHENG11.", "13306", "zc"
	err := GormAutoMigrate(host, port, use, pass, database)
	if err != nil {
		//println("连接失败", err)
		panic("连接失败,error=" + err.Error())
	}
	fmt.Println("创建表结构完成!")
}

func TestDdl() {
	host := "162.14.100.125"
	use, pass, port, database := "root", "zhaoCHENG11.", "13306", "zc"
	db, err := ConnectMysqlByDefault(host, port, use, pass, database)
	// 单个插入创建
	if err != nil {
		//println("连接失败", err)
		panic("连接失败,error=" + err.Error())
	}
	db.Create(&models.UserBasic{Name: "zhangsan", Password: "thisispassword"})

	users := []models.UserBasic{
		{Name: "lisi", Password: "333"},
		{Name: "wangwu", Password: "444"},
	}

	// 多个插入并演示返回值
	result := db.Create(users)
	if result != nil {
		fmt.Println(result.RowsAffected)
		fmt.Println(result.Error)
	}
	//根据主键ID查询主键值为1的记录并返回数据
	var name = models.UserBasic{}
	db.First(&name, 2)
	fmt.Println(name)
}
