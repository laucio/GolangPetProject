package Entity

type Project struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	EmptyRepo   bool   `json:"empty_repo"`
}
