package web

import (
	"fmt"
	"testing"
)

func TestPostMail(t *testing.T) {
	PostMail("t20220630t", 4)
}

func TestDealMail(t *testing.T) {
	DealMail("t20220630t3@litemail.top", GetCodeFromMailData)
}

func TestGetCodeFromMail(t *testing.T) {
	code := GetCodeFromMail(GetCodeFromMailData)
	fmt.Println(code)
}

func TestRegister(t *testing.T) {
	r := Register("t7@litemail.top", "569371")
	fmt.Println(r)
}

func TestInvitation(t *testing.T) {
	Invitation(26607745, Token)
}

const (
	Token =
		`eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzaWQiOjI3MTE2ODU0LCJkZXZpY2VfaWQiOiJ0NEBsaXRlbWFpbC50b3AiLCJlbWFpbCI6InQ0QGxpdGVtYWlsLnRvcCIsInBhc3N3b3JkIjoiMTExMTExIiwiZXhwIjoxNjM3MDA4MDQwLCJzeXN0ZW1fdGltZSI6MTYzNjU3NjU3MCwiZG93bmxvYWRfY291bnRzIjowLCJzaGFyZWRfY291bnRzIjowLCJ0b2RheV9zaGFyZWRfY291bnRzIjowLCJzaGFyZWRfbGFzdF90aW1lIjowLCJjaGVja2luX2xhc3RfdGltZSI6MCwiZmF2b3JpdGVfbGltaXQiOjEwMCwidGltZXN0b3BfY291bnRzIjowLCJ0aW1lc3RvcF91c2VkX2NvdW50cyI6MCwidHVybnRhYmxlX2xhc3RfdGltZSI6MCwidHVybnRhYmxlX2NvdW50cyI6MCwiYm91Z2h0X3ZpcF9iZWZvcmUiOjAsInZpcF9sZXZlbCI6MCwidmlwX3RpbGwiOjAsIm1lbWJlcl9wYWdlIjpmYWxzZSwiaGVhZCI6IiIsInN0YXR1cyI6dHJ1ZSwicG9pbnQiOjAsImRyaW5rIjowLCJpbXBvcnRfbGltaXQiOjMwLCJmb2xkZXJfbGltaXQiOjUsInN1YnNjcmliZV9saW1pdCI6MywiaXNfc2lnbiI6ZmFsc2UsIm1lbWJlcl9pZCI6MjcxMTY4NTQsInBsYXRmb3JtIjoiaW9zIn0.asame_1zFu1s_Wk4TMbpqa3eyb9gdRg6tf3uFLPNYac`
	GetCodeFromMailData =
`X-AliDM-RcptTo: dDIwMjIwNjMwdDFAbGl0ZW1haWwudG9w
Feedback-ID: default:no-reply@bjsgqx.com:trigger:193286
X-EnvId: 244582394131
Received: from 172.18.0.9(mailfrom:no-reply@bjsgqx.com fp:SMTPD_----1XmDmX4)
          by smtpdm.aliyun.com(127.0.0.1);
          Thu, 30 Jun 2022 22:26:05 +0800
Content-Type: multipart/mixed; boundary="===============9211883660444297913=="
MIME-Version: 1.0
Subject: =?utf-8?b?54ix5aiB5aW25Lya5ZGY6aqM6K+B56CB?=
From: no-reply@bjsgqx.com
To: t20220630t1@litemail.top
Date: Thu, 30 Jun 2022 14:26:05 +0000
Message-ID:
 <165659916338.1.7660546221149212193@jp-avnight-manager-worker-mail-sender>

--===============9211883660444297913==
Content-Type: multipart/alternative;
 boundary="===============8272474566714024138=="
MIME-Version: 1.0

--===============8272474566714024138==
Content-Type: text/plain; charset="utf-8"
MIME-Version: 1.0
Content-Transfer-Encoding: 7bit


--===============8272474566714024138==
Content-Type: text/html; charset="utf-8"
MIME-Version: 1.0
Content-Transfer-Encoding: 8bit

<html>

<head></head>

<body>
    <h1>爱威奶注册邮箱验证码</h1>
    <p>亲爱的用户您好，

        您的验证码是 <code>708055</code>
        
        赶紧完成注册去看片吧～
    </p>
</body>

</html>
--===============8272474566714024138==--

--===============9211883660444297913==--`
)
