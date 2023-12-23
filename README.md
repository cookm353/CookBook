# Cookbook

(Or some better title to be added later)

## Tech Stack

- Go
  - Gin
    - https://gin-gonic.com/
    - https://pkg.go.dev/github.com/gin-gonic/gin
  - Templ - https://templ.guide/
- SQLite
  - https://sqlite.org/cli.html
  - https://www.tutorialspoint.com/sqlite/sqlite_commands.htm
- HTMX
- Tailwind

## Libraries, Packages, and Tools Needed

- Bcrypt
- Auth (JWT? OAuth 2.0? Phantom tokens?)
- Gin
- Gorilla/CSRF
  - https://github.com/gorilla/csrf

## Useful Websites While Working On Project

- https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/05.3.html
- https://www.allhandsontech.com/programming/golang/how-to-use-sqlite-with-go/
- https://cheatsheetseries.owasp.org/cheatsheets/Authentication_Cheat_Sheet.html
- https://www.stackhawk.com/blog/golang-csrf-protection-guide-examples-and-how-to-enable-it/
- https://pkg.go.dev/golang.org/x/net/xsrftoken
- https://templ.guide/
- https://www.reddit.com/r/htmx/comments/13ah9z8/a_todo_htmx_application_that_uses_go_and_the/
  - https://github.com/stackus/todos

## Goals

- Create an init script that creates the database and tables
- Add a rudimentary auth system
- Protect against CSRF, XSS, and SQL injection attacks
