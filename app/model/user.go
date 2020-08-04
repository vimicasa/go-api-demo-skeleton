package model

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// TODO: In memory user
var users = []User{
	{
		ID:       "1",
		Username: "admin",
		Password: "admin",
		Role:     "admin",
	},
	{
		ID:       "2",
		Username: "basic",
		Password: "basic",
		Role:     "basic",
	},
	{
		ID:       "3",
		Username: "read",
		Password: "read",
		Role:     "read",
	},
}

func ValidUsernameAndPassword(username string, pass string) (*User, bool) {
	user := &User{}
	for k, v := range users {
		if v.Username == username && v.Password == pass {
			return &users[k], true
		}
	}
	return user, false
}

func GetUser(userID string) (*User, bool) {
	user := &User{}
	for k, v := range users {
		if v.ID == userID {
			return &users[k], true
		}
	}
	return user, false
}
