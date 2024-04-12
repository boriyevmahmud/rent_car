package config

const (
	ERR_INFORMATION     = "The server has received the request and is continuing the process"
	SUCCESS             = "The request was successful"
	ERR_REDIRECTION     = "You have been redirected and the completion of the request requires further action"
	ERR_BADREQUEST      = "Bad request"
	ERR_INTERNAL_SERVER = "While the request appears to be valid, the server could not complete the request"
	CUSTOMER_ROLE       = "customer"
	ADMIN_ROLE          = "admin"
	STATUS_NEW          = "new"
	STATUS_IN_PROCESS   = "in-process"
	STATUS_FINISHED     = "finished"
	STATUS_CANCELED     = "canceled"
	SmtpServer          = "smtp.gmail.com"
	SmtpPort            = "587"
	SmtpUsername        = "boriyevmahmud@gmail.com"
	SmtpPassword        = "vlcv iaxj duak jcww"
)

var SignedKey = []byte("MGJd@Ro]yKoCc)mVY1^c:upz~4rn9Pt!hYd]>c8dt#+%")
var ORDER_STATUS = []string{
	"new", "in-process", "finished", "canceled",
}
