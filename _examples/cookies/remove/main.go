package main

import (
	"github.com/kataras/iris/v12"
)

func newApp() *iris.Application {
	app := iris.New()

	// Set A Cookie.
	app.Get("/cookies/add/{name}/{value}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		value := ctx.Params().Get("value")

		// ctx.SetCookieKV(name, value) // <--
		// Alternatively: ctx.SetCookie(&http.Cookie{...})
		//
		// If you want to set custom the path:
		ctx.SetCookieKV(name, value, iris.CookiePath("/cookies"))
		//
		// If you want to be visible only to current request path:
		// (note that client should be responsible for that if server sent an empty cookie's path, all browsers are compatible)
		// ctx.SetCookieKV(name, value, iris.CookieCleanPath /* or iris.CookiePath("") */)
		// More:
		//                              iris.CookieExpires(time.Duration)
		//                              iris.CookieHTTPOnly(false)

		ctx.Writef("cookie added: %s = %s", name, value)
	})

	// Retrieve A Cookie.
	app.Get("/cookies/remove/{name}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")

		ctx.RemoveCookie(name, iris.CookieCleanPath) // <--
		// If you want more than the value then:
		// cookie, err := ctx.Request().Cookie(name)
		// if err != nil {
		//  handle error.
		// }

		ctx.WriteString("cookie removed: " + name)
	})

	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("hello world")
	})

	// Delete A Cookie.
	app.Delete("/cookies/{name}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")

		ctx.RemoveCookie(name) // <--
		// If you want to set custom the path:
		// ctx.SetCookieKV(name, value, iris.CookiePath("/custom/path/cookie/will/be/stored"))

		ctx.Writef("cookie %s removed", name)
	})

	return app
}

func main() {
	app := newApp()

	// GET:    http://localhost:8080/cookies/my_name/my_value
	// GET:    http://localhost:8080/cookies/my_name
	// DELETE: http://localhost:8080/cookies/my_name
	app.Listen(":8080")
}
