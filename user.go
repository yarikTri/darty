package darty

// User defines current session's user's data
type User struct {
	ID         int      `json:"-"`
	Name       string   `json:"name" binding:"required"`
	Username   string   `json:"username" binding:"required"`
	Password   string   `json:"password" binding:"required"`
	Categories []string `json:"categories"`
}
