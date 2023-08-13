package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/bartosz-bartosz/go-web/pkg/render"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.html")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.html")
}

func Test(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "test.html")
}

func Divide(w http.ResponseWriter, r *http.Request) {
	result, err := divideValues(2.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("Error: %d", err))
	}

	fmt.Fprintf(w, fmt.Sprintf("Result is %f", result))
}

// addValues adds two integers and returns the sum
func addValues(x, y int) int {
	return x + y
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Cannot divide by zero")
		return 0, err
	}
	return x / y, nil
}
