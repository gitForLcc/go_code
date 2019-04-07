package message

const (
	LoginMesType = "LoginMes"
	LoginResMes  = "LoginResMes"
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息的类型
}

//定义两个消息..后面需要在增加

type LoginMes struct {
	UserId   int    `json:"userId"`   //用户id
	UserPws  string `json:"userPwd"`  //用户密码
	UserName string `json:"userName"` //用户名
}

// type LoginResMes struct {
// 	Code  int    `json:"code"`  //返回状态码 500 表示用户为注册 200 表示登录成功
// 	Error string `json:"error"` //返回错误信息
// }
