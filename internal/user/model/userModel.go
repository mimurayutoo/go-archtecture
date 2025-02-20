package userModel

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Sex      string `json:"sex"`
}

func (u *User) NewUser(username, email, password, sex string, age int) *User {
	// ここでバリデーションを行う
	// バリデーションが通らない場合はエラーを返す
	return &User{
		Username: username,
		Email:    email,
		Password: password,
		Age:      age,
		Sex:      sex,
	}
}

// ゲッター関数
func (u *User) GetUsername() string {
	return u.Username
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) GetAge() int {
	return u.Age
}

func (u *User) GetSex() string {
	return u.Sex
}

// セッター関数
func (u *User) SetUsername(username string) {
	u.Username = username
}

func (u *User) SetEmail(email string) {
	u.Email = email
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

func (u *User) SetAge(age int) {
	u.Age = age
}

func (u *User) SetSex(sex string) {
	u.Sex = sex
}
