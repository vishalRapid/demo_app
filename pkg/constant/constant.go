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
	MISSING_SLUG            = "Missing Slug"
	UPLOAD_FAILED           = "Upload failed, Please try after sometime"
	INTERNAL_SERVER_ERR     = "Internal server error"
	FAILED_REQUEST          = "Failed request, Please try again later"
	BLOG_NOT_FOUND          = "Unable to find blog"
)

var (
	USERCOLLECTION = "users"
	BLOGCOLLECTION = "blogs"
)

var (
	SORTED_TAGS = "sorted_tags"
)

const (
	Draft = iota
	Published
	Deleted
)

const (
	Recommendation = "recommendation"
)
