package greetings

var GreetingsString = "Hello World"

func PrintGreetings(name string) string {
	return GreetingsString + "-" + name
}
