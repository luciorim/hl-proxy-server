package dto

type ProxyRequest struct {
	Method  *string            `json:"method"`
	URL     *string            `json:"url"`
	Headers *map[string]string `json:"headers"`
}
