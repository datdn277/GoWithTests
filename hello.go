package main

var helloPrefixTable = map[string]string{
	"English":    "Hello, ",
	"Spanish":    "Hola, ",
	"Vietnamese": "Chào, ",
	"French":     "Bonjour, ",
}

const defaultHelloPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix, exists := helloPrefixTable[language]
	if exists {
		return prefix + name
	}
	return defaultHelloPrefix + name
}

func main() {

}
