package api

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"litemall-api/user_web/forms"
	"litemall-api/user_web/global"
	"litemall-api/user_web/global/response"
	"litemall-api/user_web/middlewares"
	"litemall-api/user_web/models"
	"litemall-api/user_web/proto"

	//ut "github.com/go-playground/universal-translator"
	//"github.com/go-playground/validator/v10"
	//en_translations "github.com/go-playground/validator/v10/translations/en"
	//zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func InitSrcConn() {

}

//grpc code 转成http的状态码

func HandleGrpcErrorToHttp(c *gin.Context, err error) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg:": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "用户服务不可用",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": e.Code(),
				})
			}
			return
		}
	}
}

/************************** Trans **************************/
//var trans ut.Translator

func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

//func InitTrans(locale string) (err error) {
//	//修改gin框架中的validator引擎属性, 实现定制
//}

/************************** Trans **************************/

func GetUserList(ctx *gin.Context) {
	///***********************从consul获取用户服务信息*************************/
	//zap.S().Infof("从consul获取用户服务信息")
	//cfg := api.DefaultConfig()
	//consulInfo := global.ServerConfig.ConsulInfo
	//cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)
	//
	//userSrvHost := ""
	//userSrvPort := 0
	//client, err := api.NewClient(cfg)
	//if err != nil {
	//	panic(any(err))
	//}
	//
	//data, err := client.Agent().ServicesWithFilter(fmt.Sprintf("Service == \"%s\"", global.ServerConfig.UserSrvInfo.Name))
	//if err != nil {
	//	panic(any(err))
	//}
	//for _, value := range data {
	//	zap.S().Infof("user info value from consul :%v", value)
	//	userSrvHost = value.Address
	//	userSrvPort = value.Port
	//	break
	//}
	//
	//if userSrvHost == "" {
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		"msg": "user-srv不能获取",
	//	})
	//	return
	//}
	///***********************从consul获取用户服务信息*************************/
	//
	///***********************连接grpc*************************/
	//zap.S().Infof("连接grpc")
	////userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.UserSrvInfo.Host, global.ServerConfig.UserSrvInfo.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	//userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	zap.S().Errorw("[GetUserList]连接失败：", err.Error())
	//}
	//claims, _ := ctx.Get("claims")
	//currentClaims := claims.(*models.CustomClaims)
	//zap.S().Infof("访问用户：%d", currentClaims.ID)
	//userSrvClient := proto.NewUserClient(userConn)
	///***********************连接grpc*************************/
	claims, _ := ctx.Get("claims")
	currentClaims := claims.(*models.CustomClaims)
	zap.S().Infof("访问用户：%d", currentClaims.ID)

	pn := ctx.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	psize := ctx.DefaultQuery("psize", "10")
	psizeInt, _ := strconv.Atoi(psize)
	rsp, err := global.UserSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    uint32(pnInt),
		PSize: uint32(psizeInt),
	})
	if err != nil {
		zap.S().Errorw("[GetUserList]失败：", err.Error())
		HandleGrpcErrorToHttp(ctx, err)
		return
	}
	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		user := response.UserResponse{
			Id:       value.Id,
			NickName: value.NickName,
			//Birthday: time.Time(time.Unix(int64(value.BirthDay), 0)).Format("2006-01-02"),
			Birthday: response.JsonTime(time.Unix(int64(value.BirthDay), 0)),
			Gender:   int32(value.Gender),
			Mobile:   value.Mobile,
		}
		//data := make(map[string]interface{})
		//data["id"] = value.Id
		//data["name"] = value.NickName
		//data["birthday"] = value.BirthDay
		//data["gender"] = value.Gender
		//data["mobile"] = value.Mobile
		result = append(result, user)
	}
	ctx.JSON(http.StatusOK, result)
	zap.S().Debug("用户列表")
}

func HandleValidatorError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"error": removeTopStruct(errs.Translate(global.Trans)),
	})
	return
}

func PassWordLogin(c *gin.Context) {
	//表单验证
	passwordLoginForm := forms.PassWordLoginForm{}
	if err := c.ShouldBind(&passwordLoginForm); err != nil {
		//errs, ok := err.(validator.ValidationErrors)
		//if !ok {
		//	c.JSON(http.StatusOK, gin.H{
		//		"msg": err.Error(),
		//	})
		//}
		//
		//c.JSON(http.StatusBadRequest, gin.H{
		//	"error": removeTopStruct(errs.Translate(global.Trans)),
		//})
		HandleValidatorError(c, err)
		return
	}

	if !store.Verify(passwordLoginForm.CaptchaId, passwordLoginForm.Captcha, false) {
		c.JSON(http.StatusBadRequest, gin.H{
			"captcha": "图形验证码错误",
		})
		return
	}
	///***********************连接grpc*************************/
	//userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.UserSrvInfo.Host, global.ServerConfig.UserSrvInfo.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	zap.S().Errorw("[GetUserList]连接失败：", err.Error())
	//}
	//
	//userSrvClient := proto.NewUserClient(userConn)
	///***********************连接grpc*************************/

	//登录的逻辑
	if rsp, err := global.UserSrvClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
		Mobile: passwordLoginForm.Mobile,
	}); err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusBadRequest, map[string]string{
					"mobile": "用户不存在",
				})
			default:
				c.JSON(http.StatusInternalServerError, map[string]string{
					"mobile": "登录失败",
				})
			}
			return
		}
	} else {
		//只是查询到用户了而已，并没有检查密码
		if passRsp, pasErr := global.UserSrvClient.CheckPassWord(context.Background(), &proto.PasswordCheckInfo{
			Password:          passwordLoginForm.PassWord,
			EncryptedPassword: rsp.PassWord,
		}); pasErr != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"password": "登录失败",
			})
		} else {
			if passRsp.Success {
				/********************生成token********************/
				j := middlewares.NewJWT()
				claims := models.CustomClaims{
					ID:          uint(rsp.Id),
					NickName:    rsp.NickName,
					AuthorityId: uint(rsp.Role),
					StandardClaims: jwt.StandardClaims{ //jwt内置自带
						NotBefore: time.Now().Unix(),               //签名的生效时间
						ExpiresAt: time.Now().Unix() + 60*60*24*30, //30天过期
						Issuer:    "sam",
					},
				}
				token, err := j.CreateToken(claims)
				/********************生成token********************/

				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"msg": "生成token失败",
					})
					return
				}

				c.JSON(http.StatusOK, gin.H{
					"id":         rsp.Id,
					"nick_name":  rsp.NickName,
					"token":      token,
					"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000, //毫秒级别
				})
			} else {
				c.JSON(http.StatusBadRequest, map[string]string{
					"msg": "登录失败",
				})
			}
		}
	}
}
