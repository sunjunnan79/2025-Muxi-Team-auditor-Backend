package request

import "muxi_auditor/repository/model"

type LoginReq struct {
	Code string `json:"code"`
}
type RegisterReq struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
type UpdateUserReq struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
}
type GetUserReq struct {
	Role        int `json:"role"`
	Projecet_id int `json:"project_id"`
}
type UpdateUserRoleReq struct {
	Role   int `json:"role"`
	UserId int `json:"user_id"`
	ProjectPermit []model.ProjectPermit `json:"project_permit"`
}
