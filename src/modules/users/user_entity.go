package users

type Users struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Tabler interface {
	TableName() string
}

func (Users) TableName() string {
	return "user"
}
