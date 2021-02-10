package http

type status struct {
	Code int
	Message string
	Series Series
}

type Series string

func (status status) isInformational() bool {
	return status.Series == INFORMATIONAL
}

func (status status) isSuccess() bool {
	return status.Series == SUCCESS
}

func (status status) isRedirection() bool {
	return status.Series == REDIRECTION
}

func (status status) isClientError() bool {
	return status.Series == CLIENT
}

func (status status) isServerError() bool {
	return status.Series == SERVER
}

func (status status) isError() bool {
	return status.isClientError() || status.isServerError()
}

// SERIES DECLARATION
//goland:noinspection ALL
const (
	INFORMATIONAL Series = "Informational Response"
	SUCCESS Series = "Success"
	REDIRECTION Series = "Redirection"
	CLIENT Series = "Client Error"
	SERVER Series = "Server Error"
)

// HTTP STATUS DECLARATION
//goland:noinspection ALL
var (
	CONTINUE = status{Code: 100, Message: "Continue", Series: INFORMATIONAL}
	SWITCHING_PROTOCOLS = status{Code: 101, Message: "Switching Protocols", Series: INFORMATIONAL}

	OK = status{Code: 200, Message: "OK", Series: SUCCESS}
	CREATED = status{Code: 201, Message: "Created", Series: SUCCESS}
	ACCEPTED = status{Code: 202, Message: "Accepted", Series: SUCCESS}
	NON_AUTHORITATIVE_INFORMATION = status{Code: 203, Message: "Non-Authoritative Information", Series: SUCCESS}
	NO_CONTENT = status{Code: 204, Message: "No Content", Series: SUCCESS}
	RESET_CONTENT = status{Code: 205, Message: "Reset Content", Series: SUCCESS}

	MULTIPLE_CHOICES = status{Code: 300, Message: "Multiple Choices", Series: REDIRECTION}
	MOVED_PERMANENTLY = status{Code: 301, Message: "Moved Permanently", Series: REDIRECTION}
	FOUND = status{Code: 302, Message: "Found", Series: REDIRECTION}
	SEE_OTHER = status{Code: 303, Message: "See Other", Series: REDIRECTION}
	NOT_MODIFIED = status{Code: 304, Message: "Not Modified", Series: REDIRECTION}
	USE_PROXY = status{Code: 305, Message: "Use Proxy", Series: REDIRECTION}
	TEMPORARY_REDIRECT = status{Code: 307, Message: "Temporary Redirect", Series: REDIRECTION}
	PERMANENT_REDIRECT = status{Code: 308, Message: "Permanent Redirect", Series: REDIRECTION}

	BAD_REQUEST = status{Code: 400, Message: "Bad Request", Series: CLIENT}
	UNAUTHORIZED = status{Code: 401, Message: "Unauthorized", Series: CLIENT}
	PAYMENT_REQUIRED = status{Code: 402, Message: "Payment Required", Series: CLIENT}
	FORBIDDEN = status{Code: 403, Message: "Forbidden", Series: CLIENT}
	NOT_FOUND = status{Code: 404, Message: "Not Found", Series: CLIENT}
	METHOD_NOT_ALLOWED = status{Code: 405, Message: "Method Not Allowed", Series: CLIENT}
	NOT_ACCEPTABLE = status{Code: 406, Message: "Not Acceptable", Series: CLIENT}
	PROXY_AUTHENTICATION_REQUIRED = status{Code: 407, Message: "Proxy Authentication Required", Series: CLIENT}
	REQUEST_TIMEOUT = status{Code: 408, Message: "Request Timeout", Series: CLIENT}
	CONFLICT = status{Code: 409, Message: "Conflict", Series: CLIENT}
	GONE = status{Code: 410, Message: "Gone", Series: CLIENT}
	BLAZE_IT = status{Code: 420, Message: "Blaze It", Series: CLIENT}

	INTERNAL_SERVER_ERROR = status{Code: 500, Message: "Internal Server Error", Series: SERVER}
	NOT_IMPLEMENTED = status{Code: 501, Message: "Not Implemented", Series: SERVER}
	BAD_GATEWAY = status{Code: 502, Message: "Bad Gateway", Series: SERVER}
	SERVICE_UNAVAILABLE = status{Code: 503, Message: "Service Unavailable", Series: SERVER}
	GATEWAY_TIMEOUT = status{Code: 504, Message: "Gateway Timeout", Series: SERVER};
)
