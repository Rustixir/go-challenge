package localization

import (
	_ "embed"
	"sync"
)

//go:embed catalog.json
var file []byte

var instance *schema

var instanceOnce sync.Once

func init() {
	instanceOnce.Do(func() {
		instance = newSchema()
		instance.load(file)
	})
}

func Get(lang string, value string) string {
	catalog, ok := instance.dict[lang]
	if !ok {
		return value
	}
	val, ok := catalog[value]
	if !ok {
		return value
	}
	return val
}
