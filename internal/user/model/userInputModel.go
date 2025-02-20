package userModel

// jsonタグを使用するには大文字にする必要があるが、それだと外部からのアクセスができてしまうので入力を受け付ける用のモデルを定義する。
type UserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Sex      string `json:"sex"`
}