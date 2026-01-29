package enums

type ResponseCode int

const (
	// ServerError server error from 1000 to 1999
	ServerError   ResponseCode = 1000
	ServerTimeout ResponseCode = 1001
	ServerDown    ResponseCode = 1002

	// OK code success from 2000 to 2999
	OK       ResponseCode = 2000
	Created  ResponseCode = 2001
	Accepted ResponseCode = 2002

	// InvalidRequest client error from 3000 to ...
	InvalidRequest      ResponseCode = 3000
	InvalidCredentials  ResponseCode = 3001
	NotAuthorizedAccess ResponseCode = 3002
	NotFound            ResponseCode = 3003
	RequestTimeout      ResponseCode = 3004
	TooManyRequest      ResponseCode = 3005
)

func (e ResponseCode) Message() string {
	switch e {
	case ServerError:
		return "Server has error"
	case ServerTimeout:
		return "Server gateway is timed out"
	case ServerDown:
		return "Server is down or under maintenance"
	case OK:
		return "Success"
	case Created:
		return "Created"
	case Accepted:
		return "The request has been accepted"
	case InvalidRequest:
		return "Invalid Request"
	case InvalidCredentials:
		return "Security credentials is incorrect"
	case NotAuthorizedAccess:
		return "You are not authorized to access this resource"
	case NotFound:
		return "Resource not found"
	case RequestTimeout:
		return "Request timeout"
	case TooManyRequest:
		return "Too many request"
	default:
		return ""
	}
}
