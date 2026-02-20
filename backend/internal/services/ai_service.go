package services
type AIService interface{Summarize(text string)(string,error);ExtractKeywords(text string)([]string,error)}
func NewAIService()AIService{return &noopAI{}}
type noopAI struct{}
func(n *noopAI)Summarize(t string)(string,error){return "",nil}
func(n *noopAI)ExtractKeywords(t string)([]string,error){return nil,nil}