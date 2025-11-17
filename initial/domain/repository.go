package domain

type InitialRepository interface {
	GetUserByID(id UserID) (*User, error)
}
