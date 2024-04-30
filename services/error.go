package services

import "github.com/pkg/errors"

// 定义所有错误
var ErrCacheSavedFailed = errors.New("用户信息无法缓存,请联系管理员")

var ErrUserLoginFailedError = errors.New("用户名或密码错误")

var ErrFailedExtractToken = errors.New("提取凭证失败")

var ErrUserInfoNotExist = errors.New("用户信息不存在,请先注册用户")

var ErrUserNotExist = errors.New("用户不存在,请联系管理员创建")

var ErrSearchAccountFailed = errors.New("查询人员账户失败")

var ErrTooLongUserName = errors.New("用户名或姓名过长,请简短后重试")
