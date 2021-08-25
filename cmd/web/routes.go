package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable, noSurf, app.authenticate)

	mux := pat.New()
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Post("/todo/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.addTodo))
	mux.Get("/todo/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.addTodoForm))
	mux.Post("/todo/edit/:id", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.editTodo))
	mux.Get("/todo/edit/:id", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.editTodoForm))
	mux.Get("/todo/delete/:id", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.deleteTodo))
	mux.Get("/todo/:id", dynamicMiddleware.ThenFunc(app.showTodo))
	mux.Get("/todos", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.showTodos))

	mux.Get("/user/register", dynamicMiddleware.ThenFunc(app.registerUserForm))
	mux.Post("/user/register", dynamicMiddleware.ThenFunc(app.registerUser))
	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.logoutUser))

	mux.Get("/ping", http.HandlerFunc(ping))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
