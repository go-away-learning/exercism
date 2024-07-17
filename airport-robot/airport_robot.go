// Package airportrobot provides functionality for greeting visitors in different languages.
package airportrobot

import "fmt"

// Greeter is an interface that defines methods for obtaining the name of a language
// and greeting a visitor in that language.
type Greeter interface {
	LanguageName() string
	Greet(visitor string) string
}

// German represents a greeter that speaks German.
type German struct {
	languageName string
}

// Italian represents a greeter that speaks Italian.
type Italian struct {
	languageName string
}

// Portuguese represents a greeter that speaks Portuguese.
type Portuguese struct {
	languageName string
}

// LanguageName returns the name of the language spoken by the German greeter.
func (ger German) LanguageName() string {
	return ger.languageName
}

// LanguageName returns the name of the language spoken by the Italian greeter.
func (ger Italian) LanguageName() string {
	return ger.languageName
}

// LanguageName returns the name of the language spoken by the Portuguese greeter.
func (ger Portuguese) LanguageName() string {
	return ger.languageName
}

// Greet returns a greeting message in German for the given visitor.
func (ger German) Greet(visitor string) string {
	return fmt.Sprintf("I can speak German: Hallo %s!", visitor)
}

// Greet returns a greeting message in Italian for the given visitor.
func (ger Italian) Greet(visitor string) string {
	return fmt.Sprintf("I can speak Italian: Ciao %s!", visitor)
}

// Greet returns a greeting message in Portuguese for the given visitor.
func (ger Portuguese) Greet(visitor string) string {
	return fmt.Sprintf("I can speak Portuguese: Olá %s!", visitor)
}

// SayHello uses the provided greeter to greet the visitor.
func SayHello(visitor string, greeter Greeter) string {
	return greeter.Greet(visitor)
}
