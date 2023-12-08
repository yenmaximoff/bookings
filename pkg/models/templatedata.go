package models

// TemplateData содержит данные, отправленные от обработчиков в шаблон.
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	FlashMsg  string
	ErrorMsg  string
	Warning   string
}
