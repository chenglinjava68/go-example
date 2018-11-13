package common

import (
	"context"
	"github.com/labstack/echo"
	"github.com/rcrowley/go-metrics"
	"github.com/xiaomeng79/go-example/cinit"
	"github.com/xiaomeng79/go-example/internal/jwt"
	metrics2 "github.com/xiaomeng79/go-example/internal/metrics"
	"github.com/xiaomeng79/go-example/internal/trace"
	"github.com/xiaomeng79/go-log"
	"strconv"
	"strings"
	"time"
)

//opentracing中间件
func Opentracing(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, span, err := trace.TraceFromHeader(context.Background(), "api:"+c.Request().URL.Path, c.Request().Header)
		if err == nil {
			defer span.Finish()
			c.Set(cinit.TRACE_CONTEXT, ctx)
		} else {
			c.Set(cinit.TRACE_CONTEXT, context.Background())
		}
		return next(c)
	}
}

//验证jwt
func JWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//获取span
		ctx := c.Get(cinit.TRACE_CONTEXT).(context.Context)
		ctx, span, _ := trace.TraceIntoContext(ctx, "VerifyToken")
		defer span.Finish()
		//从请求头获取token信息
		jwtString := c.Request().Header.Get(cinit.JWT_NAME)
		//log.Debug(jwtString, ctx)
		//解析JWT
		auths := strings.Split(jwtString, " ")
		if strings.ToUpper(auths[0]) != "BEARER" || auths[1] == " " {
			return HandleError(c, ReqNoAllow, "token验证失败")
		}
		jwtmsg, err := jwt.Decode(strings.Trim(auths[1], " "))
		if err != nil {
			log.Info(err.Error(), ctx)
			return HandleError(c, ReqNoAllow, "token验证失败")
		}
		c.Set(cinit.JWT_MSG, jwtmsg)
		return next(c)
	}
}

//verifyparam
//trace
func VerifyParam(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//获取span
		ctx := c.Get(cinit.TRACE_CONTEXT).(context.Context)
		ctx, span, _ := trace.TraceIntoContext(ctx, "VerifyParam")
		defer span.Finish()
		//解析公共参数
		r := new(ReqParam)
		err := c.Bind(r)
		if err != nil {
			log.Info("解析参数错误:"+err.Error(), ctx)
			return HandleError(c, CommonParamConvertError)
		}
		//验证公共参数
		err = r.Validate()
		if err != nil {
			log.Info("验证参数错误:"+err.Error(), ctx)
			return HandleError(c, CommonParamConvertError, err.Error())
		}
		//请求appsecret

		//验证签名
		b, err := r.CompareSign()
		if !b {
			log.Info("获取appsecret"+err.Error(), ctx)
			return HandleError(c, CommonSignError)
		}
		//通过验证，绑定参数
		c.Set(cinit.REQ_PARAM, r)
		return next(c)
	}
}

//metrics
func MetricsFunc(m *metrics2.Metrics) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			res := c.Response()
			start := time.Now()
			if err := next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()

			latency := stop.Sub(start)
			status := res.Status
			//Total request count.
			meter := metrics.GetOrRegisterMeter("status."+strconv.Itoa(status), m.GetRegistry())
			meter.Mark(1)

			//Request size in bytes.
			//meter = metrics.GetOrRegisterMeter(m.WithPrefix("status."+strconv.Itoa(status)), m.GetRegistry())
			//meter.Mark(req.ContentLength)

			//Request duration in nanoseconds.
			h := metrics.GetOrRegisterHistogram("h_status."+strconv.Itoa(status), m.GetRegistry(),
				metrics.NewExpDecaySample(1028, 0.015))
			h.Update(latency.Nanoseconds())

			//Response size in bytes.
			//meter = metrics.GetOrRegisterMeter(m.WithPrefix("status."+strconv.Itoa(status)), m.GetRegistry())
			//meter.Mark(res.Size)

			return nil
		}
	}
}
