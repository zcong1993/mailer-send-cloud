package sender

import (
	"errors"
	"github.com/tidwall/gjson"
	"github.com/zcong1993/mailer/common"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const API_ADDRESS = "http://api.sendcloud.net/apiv2/mail/send"

type SendCloud struct {
	ApiUser    string
	ApiKey     string
	Client     http.Client
	ApiAddress string
}

func (sc *SendCloud) Send(msg common.MailMsg) (error, bool) {
	form := sc.toForm(msg)

	resp, err := sc.Client.PostForm(sc.ApiAddress, form)
	if err != nil {
		return err, true
	}
	defer resp.Body.Close()

	bt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, true
	}

	code := gjson.GetBytes(bt, "statusCode").Int()
	if code == 200 {
		return nil, false
	}
	message := gjson.GetBytes(bt, "message").String()

	return errors.New(message), true
}

func (sc *SendCloud) toForm(msg common.MailMsg) url.Values {
	u := url.Values{}

	u.Add("apiUser", sc.ApiUser)
	u.Add("apiKey", sc.ApiKey)
	u.Add("from", msg.From)
	u.Add("to", strings.Join(msg.To, ";"))
	u.Add("html", msg.Body)
	u.Add("subject", msg.Subject)

	return u
}
