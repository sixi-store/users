package db

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/it512/sqlt"
	"github.com/sixi-store/users/config"
)

// DbOp 数据库操作应用对象
var DbOp *sqlt.DbOp

//数据库连接
func init() {
	Content := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", config.MysqlUsername, config.MysqlPassword, config.MysqlHost, config.MysqlPort, config.MysqlDatabase, config.MysqlCharset)
	DbOp = sqlt.NewSqltDefault("mysql", Content, "template/*.tpl")

}

// Query 数据库查询
func Query(queryName string, params map[string]interface{}, fields []string) (result []map[string]string) {
	//smr := sqlt.NewSliceMapRowsHandler(funcs.Camal)
	smr := sqlt.NewSliceMapRowsHandler(func(i string) string {
		return i
	})
	e := DbOp.QueryContext(context.Background(), queryName, params, smr)
	if e != nil {
		log.Fatal(e)
	}
	result = make([]map[string]string, 0, 1)
	for i := 0; i < smr.Count(); i++ {
		c := smr.ResuleSet(i)
		for _, r := range c {
			item := make(map[string]string)
			if len(fields) > 0 {
				for _, f := range fields {
					switch r[f].(type) {
					case []byte:
						item[f] = fmt.Sprintf("%s", r[f])
					default:
						item[f] = fmt.Sprintf("%v", r[f])
					}
					if item[f] == "<nil>" {
						item[f] = ""
					}
				}
			} else {
				for k, f := range r {
					switch f.(type) {
					case []byte:
						item[k] = fmt.Sprintf("%s", f)
					default:
						item[k] = fmt.Sprintf("%v", f)
					}
					if item[k] == "<nil>" {
						item[k] = ""
					}
				}
			}
			result = append(result, item)
		}
	}
	return
}

// QueryRow 数据库查询单条记录
func QueryRow(queryName string, params map[string]interface{}, fields []string) (result map[string]string) {
	//smr := sqlt.NewSliceMapRowsHandler(funcs.Camal)
	smr := sqlt.NewSliceMapRowsHandler(func(i string) string {
		return i
	})
	e := DbOp.QueryContext(context.Background(), queryName, params, smr)
	if e != nil {
		log.Fatal(e)
	}
	result = make(map[string]string)
	for i := 0; i < smr.Count(); i++ {
		c := smr.ResuleSet(i)
		for _, r := range c {
			if len(fields) > 0 {
				for _, f := range fields {
					switch r[f].(type) {
					case []byte:
						result[f] = fmt.Sprintf("%s", r[f])
					default:
						result[f] = fmt.Sprintf("%v", r[f])
					}
					if result[f] == "<nil>" {
						result[f] = ""
					}
				}
			} else {
				for k, f := range r {
					switch f.(type) {
					case []byte:
						result[k] = fmt.Sprintf("%s", f)
					default:
						result[k] = fmt.Sprintf("%v", f)
					}
					if result[k] == "<nil>" {
						result[k] = ""
					}
				}
			}
		}
	}
	return
}

// Field 数据库查询单个字段
func Field(queryName string, params map[string]interface{}) (result int64) {
	//smr := sqlt.NewSliceMapRowsHandler(funcs.Camal)
	smr := sqlt.NewSliceMapRowsHandler(func(i string) string {
		return i
	})
	e := DbOp.QueryContext(context.Background(), queryName, params, smr)
	if e != nil {
		log.Fatal(e)
	}
	for i := 0; i < smr.Count(); i++ {
		c := smr.ResuleSet(i)
		for _, r := range c {
			for _, f := range r {
				var total string
				switch f.(type) {
				case []byte:
					total = fmt.Sprintf("%s", f)
				default:
					total = fmt.Sprintf("%d", f)
				}
				result, _ = strconv.ParseInt(total, 10, 64)
				return result
			}
		}
	}
	return
}
