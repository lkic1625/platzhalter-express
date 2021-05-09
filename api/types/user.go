package types

type User struct {
	ID         string
	CreatedAt  int
	updatedAt  int
	NoteBookID []string
	Languages  []Language
}
