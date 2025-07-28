package model

type User struct {
	ID       int    // Unique identifier for the user
	Username string // Username of the user
	Email    string // Email address of the user
	Password string // Password for the user account
	Role     string // Role of the user (e.g., admin, user)
	Notes    []Note // List of notes created by the user
}

func NewUser(id int, username, email, password, role string) *User {
	return &User{
		ID:       id,
		Username: username,
		Email:    email,
		Password: password,
		Role:     role,
		Notes:    []Note{},
	}
}

func (u *User) GetNotes() []Note {
	return u.Notes
}

func (u *User) AddNote(note Note) {
	u.Notes = append(u.Notes, note)
}

func (u *User) RemoveNote(noteID int) {
	for i, note := range u.Notes {
		if note.ID == noteID {
			u.Notes = append(u.Notes[:i], u.Notes[i+1:]...)
			break
		}
	}
}

func (u *User) UpdateNote(noteID int, updatedNote Note) {
	for i, note := range u.Notes {
		if note.ID == noteID {
			u.Notes[i] = updatedNote
			break
		}
	}
}

func (u *User) GetUsername() string {
	return u.Username
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role
}

func (u *User) GetID() int {
	return u.ID
}
