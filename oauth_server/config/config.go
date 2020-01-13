package config

import (
	"database/sql"
	"encoding/gob"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/antonlindstrom/pgstore"
	"github.com/jackc/pgx"
	"github.com/jinzhu/gorm"
	pg "github.com/vgarvardt/go-oauth2-pg"
	"github.com/vgarvardt/go-pg-adapter/pgxadapter"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/server"
)

const SessionCookieKey = "oauth-session"

var (
	OAuthServer      *server.Server
	OAuthManager     *manage.Manager
	OAuthClientStore *pg.ClientStore

	SessionStore *pgstore.PGStore

	GormDB *gorm.DB

	// helper variables
	db      *sql.DB
	pgxConn *pgx.Conn
)

func init() {
	initDB()
	initOAuth()
	initSessionStore()
}

func initDB() {
	dbPath := os.Getenv("OAUTH_DB")
	if dbPath == "" {
		log.Fatalf("provide `OAUTH_DB` env variable")
	}
	pgxConnConfig, err := pgx.ParseConnectionString(dbPath)
	if err != nil {
		log.Fatalf("unable to parse `OAUTH_DB` env value: %s", err)
	}
	pgxConn, err = pgx.Connect(pgxConnConfig)
	if err != nil {
		log.Fatalf("unable to connect to database: %s", err)
	}

	GormDB, err = gorm.Open("postgres", dbPath)
	if err != nil {
		log.Fatalf("unable to connect to database: %s", err)
	}

	db = GormDB.DB()
	if err = db.Ping(); err != nil {
		log.Fatalf("unable to connect to database: %s", err)
	}
}

func initOAuth() {
	OAuthManager = manage.NewDefaultManager()

	// use PostgreSQL token store with pgx.Connection adapter
	adapter := pgxadapter.NewConn(pgxConn)
	tokenStore, _ := pg.NewTokenStore(adapter, pg.WithTokenStoreGCInterval(time.Minute))
	defer tokenStore.Close()

	OAuthClientStore, _ = pg.NewClientStore(adapter)

	OAuthManager.MapTokenStorage(tokenStore)
	OAuthManager.MapClientStorage(OAuthClientStore)

	OAuthServer = server.NewDefaultServer(OAuthManager)
	OAuthServer.SetClientInfoHandler(server.ClientFormHandler)

	OAuthServer.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})
}

func initSessionStore() {
	sessionSecret := os.Getenv("OAUTH_SESSION_KEY")
	if sessionSecret == "" {
		log.Fatalf("provide `OAUTH_SESSION_KEY` env variable")
	}
	store, err := pgstore.NewPGStoreFromPool(db, []byte(sessionSecret))
	if err != nil {
		log.Fatalf(err.Error())
	}
	// enable storing request Form values in session store
	gob.Register(url.Values{})
	SessionStore = store
}
