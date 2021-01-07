package util

const (
	Success     = 200
	QueryError  = 401
	ServerError = 500
)

var codeMessage = map[int]string {
	Success:     "OK",
	QueryError:  "QUERY PARAM ERROR",
	ServerError: "SERVER ERROR",
}

func GetCodeMessage(code int) string {
	return codeMessage[code]
}