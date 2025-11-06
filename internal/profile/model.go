package profile

type Profile struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type Profiles struct {
	Active  string     `json:"active"`
	Entries []Profile  `json:"entries"`
}
