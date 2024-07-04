package dto

type ProxyResponse struct {
	ID      string              `json:"id"`
	URL     string              `json:"url"`
	Status  int                 `json:"status"`
	Headers map[string][]string `json:"headers"`
	Length  int                 `json:"length"`
}
