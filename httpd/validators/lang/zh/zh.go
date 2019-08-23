package zh

import (
	"go-trading/httpd/validators/lang"
)

var fields = lang.Fields{
	"Mobile":   "手机号",
	"Password": "密码",
}

var rules = lang.Rules{
	"required": "{0}为必填项",
}

var Lang = lang.Lang{Fields: fields, Rules: rules}
