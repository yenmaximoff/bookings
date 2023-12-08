package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

/*
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}

*/

// добавляет защиту CSRF ко всем запросам POST
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// загружает и сохраняет сессию при каждом запросе
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
