package web

import (
	"encoding/json"
	"fmt"
	"github.com/roambin/litemail/utils/logutil"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func PostMail(prefix string, n int) {
	resp, err := http.Get(fmt.Sprintf("https://apiw.zhidaiguanjia.com/vw2/verification_code/%s%d@litemail.top", prefix, n))
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
	Invitation(26607745, resp.Data.Token)
}

type RegisterResp struct {
	Data struct {
		Email         string `json:"email"`
		FavoriteLimit int    `json:"favorite_limit"`
		MemberID      int    `json:"member_id"`
		Msg           string `json:"msg"`
		Password      string `json:"password"`
		Success       bool   `json:"success"`
		Token         string `json:"token"`
	} `json:"data"`
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
	url := fmt.Sprintf("https://apiw.zhidaiguanjia.com/vw2/register/%s", addr)
	method := "POST"
	payload := strings.NewReader("password=111111&code=" + code)
	client := &http.Client {}
	req, err := http.NewRequest(method, url, payload)
	if err := logutil.LogWrapIfError(nil, err); err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err := logutil.LogWrapIfError(nil, err); err != nil {
		return
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	err = json.Unmarshal(body, &registerResp)
	return
}

func Invitation(invitationCode int, token string) {
	params := url.Values{}
	params.Add("invitation_code", strconv.Itoa(invitationCode))
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", "https://apiw.zhidaiguanjia.com/vw2/invitation", body)
	if err := logutil.LogWrapIfError(nil, err); err != nil {
		return
	}
	req.Header.Set("Access-Token", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err := logutil.LogWrapIfError(nil, err); err != nil {
		return
	}
	s, _ := ioutil.ReadAll(resp.Body)
	logutil.Info(nil, string(s))
	defer resp.Body.Close()
}