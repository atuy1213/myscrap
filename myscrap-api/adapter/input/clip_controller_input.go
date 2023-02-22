package input

type ClipArticleInput struct {
	AccessKeyID     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
	URL             string `json:"url"`
	// 
	
	Title           *string `json:"title"`
}
