package Entity

type Project struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	EmptyRepo   bool   `json:"empty_repo"`
	ReadmeUrl   string `json:"readme_url"`
	StartDate   string `json:"created_at"`
	EndDate     string `json:"last_activity_at"`
}
