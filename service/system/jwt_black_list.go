package system

import (
	"context"
	"go.uber.org/zap"
	"rebuildServer/global"
	"rebuildServer/model/system"
	"time"
)

type JwtService struct {
}

//
//  JsonInBlacklist: 拉黑jwt
func (jwtService *JwtService) JsonInBlacklist(jwtList system.JwtBlacklist) (err error) {
	err = global.GVA_DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	global.BlackCache.SetDefault(jwtList.Jwt, struct {
	}{})
	return
}

//
//  IsBlacklist: 判断JWT是否在黑名单内部
func (jwtService *JwtService) IsBlacklist(jwt string) bool {
	_, ok := global.BlackCache.Get(jwt)
	return ok
}

//
//  GetRedisJWT: 从redis取jwt
func (jwtService *JwtService) GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = global.GVA_REDIS.Get(context.Background(), userName).Result()
	return err, redisJWT
}

//
//  SetRedisJWT: jwt存入redis并设置过期时间
func (jwtService *JwtService) SetRedisJWT(jwt string, username string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.GVA_CONFIG.JWT.ExpiresTime) * time.Second
	err = global.GVA_REDIS.Set(context.Background(), username, jwt, timer).Err()
	return err
}

func LoadAll() {
	var data []string
	err := global.GVA_DB.Model(&system.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		global.GVA_LOG.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct {
		}{})
	} // jwt黑名单 加入 BlackCache重
}
