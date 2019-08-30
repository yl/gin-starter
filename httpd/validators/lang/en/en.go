package en

import (
	"github.com/yangliulnn/gin-starter/httpd/validators/lang"
)

var fields = lang.Fields{
	"Mobile":   "mobile",
	"Password": "password",
}

var rules = lang.Rules{
	"required": "{0} is required",
}

var Lang = lang.Lang{Fields: fields, Rules: rules}
