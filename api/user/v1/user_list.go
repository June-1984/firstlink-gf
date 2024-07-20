package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf-demo-user/v2/internal/model/entity"
)

type ListReq struct {
	g.Meta `path:"/user/list" method:"get" tags:"UserService" summary:"User list"`
}
type ListRes struct {
	*entity.User
}