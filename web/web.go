package web

import (
	"encoding/json"
	"fmt"
	"github.com/roambin/litemail/utils/logutil"
	"io/ioutil"
	"net/http"
	"strings"
)

func PostMail(prefix string, n int) {
	httpUrl := fmt.Sprintf("https://apit.hebijinxiang.com/vw3/verification_code")
	payload := fmt.Sprintf(`{"account":"%s%d@litemail.top"}`, prefix, n)
	resp, err := http.Post(httpUrl, "application/json;charset=UTF-8", strings.NewReader(payload))
	if err := logutil.LogWrapIfError(nil, err); err != nil {
		return
	}
	s, _ := ioutil.ReadAll(resp.Body)
	logutil.Info(nil, string(s))
	defer resp.Body.Close()
}

func DealMail(addr string, data string) {
	code := GetCodeFromMail(data)
	resp := Register(addr, code)
	Invitation(31254700, resp.Token)
}

type RegisterResp struct {
	Msg           string `json:"msg"`
	Success       bool   `json:"success"`
	Token         string `json:"token"`
}

func GetCodeFromMail(s string) string {
	startIndex := strings.Index(s, "<code>") + len("<code>")
	endIndex := strings.Index(s, "</code>")
	if startIndex < 0 || endIndex < 0 || startIndex > endIndex {
		return ""
	}
	codeStr := s[startIndex: endIndex]
	return codeStr
}

func Register(addr string, code string) (registerResp RegisterResp) {
	httpUrl := fmt.Sprintf("https://apit.hebijinxiang.com/vw3/register")
	payload := fmt.Sprintf(`{"account":"%s","password":"111111","verification_code":"%s","platform":"pwa"}`, addr, code)
	resp, err := http.Post(httpUrl, "application/json;charset=UTF-8", strings.NewReader(payload))
	if err := logutil.LogWrapIfError(nil, err); err != nil {
		return
	}
	s, _ := ioutil.ReadAll(resp.Body)
	logutil.Info(nil, string(s))
	defer resp.Body.Close()

	err = json.Unmarshal(s, &registerResp)
	logutil.LogIfError(nil, err)
	return
}

func Invitation(invitationCode int, token string) {
	payload := fmt.Sprintf(`{"invitation_code": "%d"}`, invitationCode)

	req, err := http.NewRequest("POST", "https://apit.hebijinxiang.com/vw3/202206/invitation", strings.NewReader(payload))
	if err := logutil.LogWrapIfError(nil, err); err != nil {
		return
	}
	req.Header.Set("Access-Token", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	resp, err := http.DefaultClient.Do(req)
	if err := logutil.LogWrapIfError(nil, err); err != nil {
		return
	}
	defer resp.Body.Close()
	s, _ := ioutil.ReadAll(resp.Body)
	logutil.Info(nil, string(s))
}