package db

import (
	"context"
	"fmt"

	"github.com/it512/sqlt"
	"github.com/sixi-store/users/pb"
)

// List 查询
func List(queryName string, params map[string]interface{}) (w *pb.UserListReplay, err error) {
	smr := sqlt.NewSliceMapRowsHandler(func(i string) string { return i })
	err = DbOp.QueryContext(context.Background(), queryName, params, smr)
	if err != nil {
		return
	}
	for i := 0; i < smr.Count(); i++ {
		c := smr.ResuleSet(i)
		for _, r := range c {
			fmt.Println(r)
			//item := make(map[string]string)
			//w.Data = append(w.Data, item)
		}
	}
	return
}
