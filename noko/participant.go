package noko

type Participant struct {
	ID              int    `json:"id"`
	Email           string `json:"email"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	ProfileImageURL string `json:"profile_image_url"`
	URL             string `json:"url"`
}
