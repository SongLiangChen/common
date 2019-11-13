package exception

type Exception struct {
	Code    int
	Message string
}

var (
	exps = make(map[string]*Exception)
)

func RegisterException(key string, code int, message string) {
	if _, ok := exps[key]; ok {
		panic("register fail: exception " + key + " already exist")
	}

	exps[key] = &Exception{
		Code:    code,
		Message: message,
	}
}

func Exp(key string) *Exception {
	if _, ok := exps[key]; ok {
		panic("exception " + key + " not exist")
	}
	return exps[key]
}
