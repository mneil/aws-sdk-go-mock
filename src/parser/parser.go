package parser

import (
	"sync"

	"github.com/alecthomas/participle/v2"
)

var instance *participle.Parser[JqLite]
var once sync.Once

func Parser() *participle.Parser[JqLite] {
	once.Do(func() {
		instance = participle.MustBuild[JqLite]()
	})
	return instance
}
