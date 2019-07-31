package testmod

import (
	"fmt"
)

func Hi(name, lang string) string {
	return fmt.Sprintf("Hi, %s !~  %s~ ", name, lang)
} 
