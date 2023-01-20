package constant

var (
	MISSING_REQUIRED_FIELDS = "Missing required fields"
	PARSING_ERROR           = "Error while parsing request"
	MISSING_NAME            = "Name  required"
	MISSING_EMAIL           = "Email  required"
	MISSING_PASSWORD        = "Password is required"
	INVALID_EMAIL           = "Please check email "
	USER_EXIST              = "User exist with this email, Try login"
	SUCCESS                 = "Success"
	USER_NOT_EXIST          = "Email is not register, try creating an account"
	INVALID_CREDENTIAL      = "Invalid credentials, Please try again"
	MISSING_TOKEN           = "Token is required"
	INVALID_REQUEST         = "Invalid Request"
	TOKEN_MALFORMED         = "Token is malformed or expired"
	UNAUTH_REQUEST          = "UnAuthorized Request"
	MISSING_TITLE           = "Please provide title for blog"
)

var (
	USERCOLLECTION = "users"
	BLOGCOLLECTION = "blogs"
)

const (
	Draft = iota
	Published
	Deleted
)
