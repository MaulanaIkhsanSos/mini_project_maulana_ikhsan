package model

// User adalah entitas yang merepresentasikan pengguna dalam aplikasi.
type User struct {
	UserID       int
	NamaPengguna string
	Email        string
	KataSandi    string
}

// NewUser adalah fungsi untuk membuat pengguna baru.
func NewUser(userID int, namaPengguna, email, kataSandi string) *User {
	return &User{
		UserID:       userID,
		NamaPengguna: namaPengguna,
		Email:        email,
		KataSandi:    kataSandi,
	}
}
