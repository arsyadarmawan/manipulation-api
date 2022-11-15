package domain

type Career struct {
	Id          string `json:"id"`
	Type        string `json:"type"`
	Company     string `json:"company"`
	Company_Url string `json:"company_url"`
	Location    string `json:"location"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
