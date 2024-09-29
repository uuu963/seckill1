package model

type user struct {
	userName string
	passWord string
}

func newUser(userName string, passWord string) *user {
	return &user{
		userName: userName,
		passWord: passWord,
	}
}

// 增删改查
