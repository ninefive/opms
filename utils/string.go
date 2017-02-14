package utils

import (
	"crypto/rand"
	"math/big"

	"net/smtp"
	"regexp"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/zheng-ji/goSnowFlake"
)

//字符串截取
func SubString(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func GetFileSuffix(s string) string {
	re, _ := regexp.Compile(".(jpg|jpeg|png|gif|exe|doc|docx|ppt|pptx|xls|xlsx)")
	suffix := re.ReplaceAllString(s, "")
	return suffix
}

func RandInt64(min, max int64) int64 {
	maxBigInt := big.NewInt(max)
	i, _ := rand.Int(rand.Reader, maxBigInt)
	if i.int64() < min {
		RandInt64(min, max)
	}
	return i.Int64()
}

func Strim(str string) string {
	str = strings.Replace(str, "\t", "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\r", "", -1)
	return str
}

func Unicode(rs string) string {
	json := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			json += string(r)
		} else {
			json += "\\u" + strconv.FormatInt(int64(rint), 16)
		}
	}
	return json
}

func HTMLEncode(rs string) string {
	html := ""
	for _, r := range rs {
		html += "&#" + strconv.Itoa(int(r)) + ";"
	}
	return html
}

/**
* to: example@example.com;example1@example.com;...
* subject: The subject of mail
* body: The content of mail
 */
func SendMail(to, subject, body string) error {
	user := beego.AppConfig.String("mailfrom")
	password := beego.AppConfig.String("mailpassword")
	host := beego.AppConfig.String("mailhost")

	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	content_type = "Content-type:text/html;charset=utf-8"

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strins.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func SnowFlakeId() int64 {
	iw, _ := goSnowFlake.NewIdWorker(1)
	if id, err := iw.NextId(); err != nil {
		return 0
	} else {
		return id
	}
}
