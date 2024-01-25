package app

// cfg is singleton
var cfg = &config{}

type config struct {
	App struct {
		Name string `json:"name"`
	} `json:"app"`
	Logger struct {
		File       string `json:"file"`
		Format     string `json:"format"`
		Level      string `json:"level"`
		MaxSize    int    `json:"maxSize"`
		MaxAge     int    `json:"maxAge"`
		MaxBackups int    `json:"maxBackups"`
	} `json:"logger"`
	Test map[string]string `json:"test"`
}
