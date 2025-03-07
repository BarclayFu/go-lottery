package comm

import (
	"crypto/md5"
	"fmt"
	"go-lottery/conf"
	"go-lottery/models"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
)

func ClientIP(request *http.Request) string {
	host, _, _ := net.SplitHostPort(request.RemoteAddr)
	return host
}

func Redirect(writer http.ResponseWriter, url string) {
	writer.Header().Add("Location", url)
	writer.WriteHeader(http.StatusFound)
}

func GetLoginUser(request *http.Request) *models.ObjLoginuser {
	c, err := request.Cookie("lottery_loginuser")
	if err != nil {
		return nil
	}
	params, err := url.ParseQuery(c.Value)
	if err != nil {
		return nil
	}
	uid, err := strconv.Atoi(params.Get("uid"))
	if err != nil || uid < 1 {
		return nil
	}
	now, err := strconv.Atoi(params.Get("now"))
	if err != nil || NowUnix()-now > 86400*30 {
		return nil
	}
	loginuser := &models.ObjLoginuser{}
	loginuser.Uid = uid
	loginuser.Username = params.Get("username")
	loginuser.Now = now
	loginuser.Ip = ClientIP(request)
	loginuser.Sign = params.Get("sign")
	sign := createLoginuserSign(loginuser)
	if sign != loginuser.Sign {
		log.Println("func_web GetLoginuser createloginusersign not signed", sign, loginuser.Sign)
		return nil
	}
	return loginuser
}

func SetLoginuser(writer http.ResponseWriter, loginuser *models.ObjLoginuser) {
	if loginuser == nil || loginuser.Uid < 1 {
		c := &http.Cookie{
			Name:   "lottery_loginuser",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		}
		http.SetCookie(writer, c)
	}
	if loginuser.Sign == "" {
		loginuser.Sign = createLoginuserSign(loginuser)
	}
	params := url.Values{}
	params.Add("uid", strconv.Itoa(loginuser.Uid))
	params.Add("username", loginuser.Username)
	params.Add("now", strconv.Itoa(loginuser.Now))
	params.Add("sign", loginuser.Sign)
	params.Add("ip", loginuser.Ip)

	c := &http.Cookie{
		Name:  "lottery_loginuser",
		Value: params.Encode(),
		Path:  "/",
	}
	http.SetCookie(writer, c)
}

func createLoginuserSign(loginuser *models.ObjLoginuser) string {
	str := fmt.Sprintf("uid=%d&username=%s&secret=%s&now=%d",
		loginuser.Uid, loginuser.Username, conf.CookieSecret, loginuser.Now)
	sign := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return sign
}
