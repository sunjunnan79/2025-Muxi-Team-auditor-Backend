package errors

import (
	"muxi_auditor/pkg/errorx"
	"net/http"
)

// 400
const (
	UNAUTHORIED_ERROR_CODE = 40001
	BAD_ENTITY_ERROR_CODE  = 40002
)

// 500
const (
	ERROR_TYPE_ERROR_CODE    = 50001
	OAUTH_GETINFO_ERROR_CODE = 50002
	LOGIN_ERROR_CODE         = 50003
)

// Auth
var (
	OAUTH_GETINFO_ERROR = func(err error) error {
		return errorx.New(http.StatusInternalServerError, OAUTH_GETINFO_ERROR_CODE, "从通行证获取用户信息失败!", "Auth", err)
	}

	LOGIN_ERROR = func(err error) error {
		return errorx.New(http.StatusInternalServerError, LOGIN_ERROR_CODE, "系统发生内部错误,登陆失败!", "Auth", err)
	}
)

// Common
var (
	BAD_ENTITY_ERROR = func(err error) error {
		return errorx.New(http.StatusUnprocessableEntity, BAD_ENTITY_ERROR_CODE, "请求参数错误", "Common", err)
	}

	UNAUTHORIED_ERROR = func(err error) error {
		return errorx.New(http.StatusUnauthorized, UNAUTHORIED_ERROR_CODE, "Authorization错误", "Common", err)
	}
)
