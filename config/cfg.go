package config

// Cfg is singleton
var Cfg = &cfg{}

type cfg struct {
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
