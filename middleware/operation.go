package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"rebuildServer/global"
	"rebuildServer/model/system"
	"rebuildServer/service"
	"rebuildServer/utils"
	"strconv"
	"time"
)

var operationRecordService = service.ServuceGroupApp.SystemServiceGroup.OperationRecordService

func OperationRecord() gin.HandlerFunc {
	return func(context *gin.Context) {
		var body []byte
		var userId int
		if context.Request.Method != http.MethodPost {
			var err error
			body, err := ioutil.ReadAll(context.Request.Body)
			if err != nil {
				global.GVA_LOG.Error("read body from request error:", zap.Error(err))
			} else {
				context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			}
		}
		claims, _ := utils.GetClaims(context)
		if claims.ID != 0 {
			userId = int(claims.ID)
		} else {
			id, err := strconv.Atoi(context.Request.Header.Get("x-user-id"))
			if err != nil {
				userId = 0
			}
			userId = id
		}
		record := system.SysOperationRecord{
			Ip:     context.ClientIP(),
			Method: context.Request.Method,
			Path:   context.Request.URL.Path,
			Agent:  context.Request.UserAgent(),
			Body:   string(body),
			UserID: userId,
		}
		// 存在某些位置错误 TODO
		//values := context.Request.Header.Values("content-type")
		//if len(values) > 0 && strings.Contains(values[0], "boundary") {
		//	record.Body = "file"
		//}
		writer := responseBodyWriter{
			ResponseWriter: context.Writer,
			body:           &bytes.Buffer{},
		}
		context.Writer = writer
		now := time.Now()

		context.Next()

		latency := time.Since(now)
		record.ErrorMessage = context.Errors.ByType(gin.ErrorTypePrivate).String()
		record.Status = context.Writer.Status()
		record.Latency = latency
		record.Resp = writer.body.String()

		if err := operationRecordService.CreateSysOperationRecord(record); err != nil {
			global.GVA_LOG.Error("create operation record error: ", zap.Error(err))
		}
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
