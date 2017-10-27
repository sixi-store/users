package app

import (
	"context"
	"fmt"

	"github.com/it512/sqlt"
	"github.com/sixi-store/users/db"
	"github.com/sixi-store/users/pb"
)

// type context = gin.Context

type pager struct {
	Total    uint64                   `json:"total"`
	LastPage uint64                   `json:"last_page"`
	Data     []map[string]interface{} `json:"data"`
}

// GetLastPage 计算最后一页
func GetLastPage(total, pageSize int64) (lastPage int64) {
	lastPage = total / pageSize
	if total%pageSize > 0 {
		lastPage++
	}
	return
}

// UserServer 用户相关服务
type UserServer struct{}

// Info implements UserServer.GetUser
func (s *UserServer) Info(ctx context.Context, r *pb.UserInfoRequest) (w *pb.UserInfoReply, err error) {
	param := make(map[string]interface{})
	param["name"] = r.GetName()
	w = &pb.UserInfoReply{}
	// 查询详情
	smr := sqlt.NewSliceMapRowsHandler(func(i string) string { return i })
	err = db.DbOp.QueryContext(context.Background(), "user.info", param, smr)
	if err != nil {
		return
	}
	for i := 0; i < smr.Count(); i++ {
		c := smr.ResuleSet(i)
		for _, r := range c {
			PullUserInfoField(w, r)
		}
	}
	return
}

// List implements UserServer.GreeterServer
func (s *UserServer) List(ctx context.Context, r *pb.UserListRequest) (w *pb.UserListReplay, err error) {
	param := make(map[string]interface{})
	page := r.GetPage()
	pageSize := r.GetPageSize()
	w = &pb.UserListReplay{}
	w.Total = db.Field("user.count", param)
	if w.Total == 0 {
		return
	}
	param["pageSize"] = pageSize
	param["offset"] = pageSize * (page - 1)
	// 查询列表
	smr := sqlt.NewSliceMapRowsHandler(func(i string) string { return i })
	err = db.DbOp.QueryContext(context.Background(), "user.list", param, smr)
	if err != nil {
		return
	}
	for i := 0; i < smr.Count(); i++ {
		c := smr.ResuleSet(i)
		for _, r := range c {
			item := &pb.UserInfoReply{}
			PullUserInfoField(item, r)
			w.Data = append(w.Data, item)
		}
	}
	return
}

//PullUserInfoField 填充用户字段
func PullUserInfoField(w *pb.UserInfoReply, r map[string]interface{}) {
	w.Id, _ = r["id"].(int64)
	w.Name = fmt.Sprintf("%s", r["name"])
	w.Email = fmt.Sprintf("%s", r["email"])
	w.Mobile = fmt.Sprintf("%s", r["mobile"])
	w.CreatedAt = fmt.Sprintf("%s", r["created_at"])
	w.UpdatedAt = fmt.Sprintf("%s", r["updated_at"])
	w.Password = fmt.Sprintf("%s", r["password"])
	w.Salt = fmt.Sprintf("%s", r["salt"])
	//w.Sex = r["sex"].(int64)
	w.Nickname = fmt.Sprintf("%s", r["nickname"])
	if w.Nickname == "%!s(<nil>)" {
		w.Nickname = ""
	}
}
