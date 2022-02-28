module main

go 1.17

replace register => ./routes/register

replace home => ./routes/home

replace mungo => ./mungo

replace templates => ./rendertemplates

require (
	edit v0.0.0-00010101000000-000000000000
	github.com/carlmjohnson/gateway v1.20.7
	home v0.0.0-00010101000000-000000000000
	logout v0.0.0-00010101000000-000000000000
	post v0.0.0-00010101000000-000000000000
	register v0.0.0-00010101000000-000000000000
	showpost v0.0.0-00010101000000-000000000000
)

require (
	github.com/aws/aws-lambda-go v1.11.1 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/gorilla/sessions v1.2.1 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.0.2 // indirect
	github.com/xdg-go/stringprep v1.0.2 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	go.mongodb.org/mongo-driver v1.7.4 // indirect
	golang.org/x/crypto v0.0.0-20200302210943-78000ba7a073 // indirect
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e // indirect
	golang.org/x/text v0.3.5 // indirect
	mungo v0.0.0-00010101000000-000000000000 // indirect
	templates v0.0.0-00010101000000-000000000000 // indirect
)

replace showpost => ./routes/showpost

replace edit => ./routes/edit

replace logout => ./routes/logout

replace post => ./routes/post
