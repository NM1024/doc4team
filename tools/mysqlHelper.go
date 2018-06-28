package tools

import (
	"os"

	//	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var Xdb *xorm.Engine

func init() {
	var (
		mysqlconnstr string
		err          error
	)

	mysqlconnstr = os.Getenv("DB_CONNSTR")
	//一般情况下如果只操作一个数据库，只需要创建一个engine即可。engine是GoRoutine安全的。
	//创建完成engine之后，并没有立即连接数据库，此时可以通过engine.Ping()来进行数据库的连接测试是否可以连接到数据库。另外对于某些数据库有连接超时设置的，可以通过起一个定期Ping的Go程来保持连接鲜活。
	Xdb, err = xorm.NewEngine("mysql", mysqlconnstr)
	// defer Xdb.Close()
	if err != nil {
		panic(err)
	}
	err = Xdb.Ping()
	if err != nil {
		panic(err)
	}
	//设置连接池的空闲数大小
	Xdb.SetMaxIdleConns(1)
	//设置最大打开连接数
	Xdb.SetMaxOpenConns(2)
	//会在控制台打印出生成的SQL语句
	Xdb.ShowSQL(true)
	//在控制台打印调试及以上的信息
	Xdb.Logger().SetLevel(core.LOG_DEBUG)

	// 启用缓存
	// cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	// MasterDB.SetDefaultCacher(cacher)

}

//func init() {
//	var (
//		mysqlconnstr string
//		err          error
//	)

//	mysqlconnstr = os.Getenv("DB_CONNSTR")
//	GormDB, err = gorm.Open("mysql", mysqlconnstr)
//	defer GormDB.Close()
//	// 全局禁用表名复数
//	GormDB.SingularTable(true)

//	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
//		return strings.Replace(defaultTableName, "_", "", -1)
//	}

//	err = GormDB.DB().Ping()
//	if err != nil {
//		panic(err)
//	}
//}

// // DB function
// func DB() *sql.DB {
// 	user := os.Getenv("DB_USER")
// 	password := os.Getenv("DB_PASSWORD")
// 	host := os.Getenv("DB_HOST")
// 	_db := os.Getenv("DB")

// 	db, _ := sql.Open("mysql", user+":"+password+"@tcp("+host+")/"+_db)
// 	err := db.Ping()

// 	if err != nil {
// 		panic(err)
// 	}
// 	return db
// }
