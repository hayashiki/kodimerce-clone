package entities

// import

const (
	ENTITY_USER         = "user"
	ENTITY_USER_SESSION = "user_session"
)

type User struct {
	Email           string `json:"email" datastore:"-"`
	PasswordHash    string `json:"password_hash" datastore:"password_hash,noindex"`
	UserType        string `json:"user_type" datastore:"user_type"`
	LastVisitedPath string `json:"last_visited_path" datastore:"last_visited_path"`
}

func NewUser(email string) *User {
	return &User{
		Email:    email,
		UserType: "regular",
	}
}
