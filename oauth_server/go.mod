module github.com/sterenczak-marek/notes-app-oauth-stack/oauth_server

require (
	github.com/antonlindstrom/pgstore v0.0.0-20170604072116-a407030ba6d0
	github.com/gofrs/uuid v3.2.0+incompatible // indirect
	github.com/google/uuid v1.1.1
	github.com/gorilla/mux v1.6.2
	github.com/gorilla/sessions v1.2.0 // indirect
	github.com/jackc/pgx v3.6.0+incompatible
	github.com/jinzhu/gorm v1.9.11
	github.com/spf13/cobra v0.0.5
	github.com/urfave/negroni v1.0.0
	github.com/vgarvardt/go-oauth2-pg v0.0.0-20190915192239-5a35396402f1
	github.com/vgarvardt/go-pg-adapter v0.3.0
	golang.org/x/crypto v0.0.0-20190911031432-227b76d455e7
	gopkg.in/oauth2.v3 v3.12.0
)

go 1.13
