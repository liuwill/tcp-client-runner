package abstract

type Client interface {
	IsConnect() bool
	IsLogin() bool
	Login(loginStatus bool)
	SetUid(uid string)
	GetUid() string
	SetUsername(username string)
	GetUsername() string
	SetProtocol(protocol string)
	GetProtocol() string
	SetTempProtocol(protocol string)
	GetTempProtocol() string
	Connect()
	Close()
	SendBytes(message []byte)
	SendMessage(message interface{})
}
