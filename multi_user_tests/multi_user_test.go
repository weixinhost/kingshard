package tests

import (
	"testing"
	"time"

	"github.com/weixinhost/litedb"
)

const (
	KingShardHost string = "127.0.0.1"
	KingShardPort uint32 = 9696
)

var userList []map[string]string = []map[string]string{

	map[string]string{
		"user":     "user_1",
		"password": "pass_1",
		"database": "test_1",
		"table":    "test_1",
	},

	map[string]string{
		"user":     "user_2",
		"password": "pass_2",
		"database": "test_2",
		"table":    "test_2",
	},

	map[string]string{
		"user":     "user_3",
		"password": "pass_3",
		"database": "test_3",
		"table":    "test_3",
	},

	map[string]string{
		"user":     "user_4",
		"password": "pass_4",
		"database": "test_4",
		"table":    "test_4",
	},
}

func initDB(user, pass, database string) *litedb.Client {

	client := litedb.NewTcpClient(KingShardHost, KingShardPort, user, pass, database)
	return client
}

//该示例用于测试用户隔离是否正常
func Test_Show_Database(t *testing.T) {

	for _, u := range userList {
		db := initDB(u["user"], u["password"], u["database"])
		go (func(db *litedb.Client, u map[string]string) {
			sql := "SHOW DATABASES;"
			for i := 0; i < 10000; i++ {
				ret := db.Query(sql)
				if ret.Err != nil {
					t.Fatal(u, ret.Err)
					break
				}

				store, err := ret.ToMap()
				if err != nil {
					t.Fatal(u, err)
					break
				}

				if len(store) < 1 {
					t.Fatal(u, "no data found")
					break
				}

				hasDB := false
				for _, item := range store {
					if item["Database"] == u["database"] {
						hasDB = true
						break
					}
				}

				if !hasDB {
					t.Fatal(u, "database not found")
				}

			}
		})(db, u)
	}
}

//该示例测试数据插入
func Test_Insert(t *testing.T) {

	for _, u := range userList {
		db := initDB(u["user"], u["password"], u["database"])
		go (func(db *litedb.Client, u map[string]string) {
			sql := "INSERT INTO `" + (u["table"]) + "` (`f1`,`f2`,`f3`) VALUES (1,'2','3')"
			for i := 0; i < 10000; i++ {
				ret := db.Exec(sql)
				if ret.Err != nil {
					t.Fatal(u, ret.Err)
					break
				}
			}
		})(db, u)
	}

}

//该示例测试数据更新
func Test_Update(t *testing.T) {
	for _, u := range userList {
		db := initDB(u["user"], u["password"], u["database"])
		go (func(db *litedb.Client, u map[string]string) {
			sql := "UPDATE `" + (u["table"]) + "` SET `f1` = 4 where 1;"
			for i := 0; i < 10000; i++ {
				ret := db.Exec(sql)
				if ret.Err != nil {
					t.Fatal(u, ret.Err)
					break
				}
			}
		})(db, u)
	}
}

//该示例测试数据更新
func Test_Delete(t *testing.T) {

	for _, u := range userList {
		db := initDB(u["user"], u["password"], u["database"])
		go (func(db *litedb.Client, u map[string]string) {
			sql := "DELETE FROM`" + (u["table"]) + "` WHERE `id` = 1"
			for i := 0; i < 10000; i++ {
				ret := db.Exec(sql)
				if ret.Err != nil {
					t.Fatal(u, ret.Err)
					break
				}
			}
		})(db, u)
	}
}

//该示例测试数据读取
func Test_Query(t *testing.T) {

	for _, u := range userList {
		db := initDB(u["user"], u["password"], u["database"])
		go (func(db *litedb.Client, u map[string]string) {
			sql := "SELECT * FROM`" + (u["table"]) + "` WHERE `id` = 1"
			for i := 0; i < 10000; i++ {
				ret := db.Exec(sql)
				if ret.Err != nil {
					t.Fatal(u, ret.Err)
					break
				}
			}
		})(db, u)
	}
}

//该示例混合随机测试，模拟真实请求时序
func Test_MultiRandom(t *testing.T) {

	go Test_Show_Database(t)

	go Test_Insert(t)

	go Test_Update(t)

	go Test_Delete(t)

	go Test_Query(t)

	time.Sleep(60 * time.Second)

}

// =================================================

func init() {

}
