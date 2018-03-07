package main

import (
	"log"
	"net/http"
	"os"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

const (
	_key = "UI_SESSION_KEY"
)

var (
	sessionStore *sessions.FilesystemStore
)

func startServer(sessionDir string, enpoint string) error {

	// Tworzenie katalogu na sesje
	if err := os.MkdirAll(sessionDir, 0700); err != nil {
		return err
	}

	// Tworzenie składnicy sesji
	sessionStore = sessions.NewFilesystemStore(sessionDir, []byte(_key))

	// Konfiguracja źródła plików z bindata.go
	assetFs := &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo}

	// Stworzenie nowego routera zapytań HTTP
	r := mux.NewRouter()

	// Podpięcie obsługi WebSocketów
	r.Path("/ws").HandlerFunc(webSocketHandler)

	// Podpięcie statycznych plików pod router
	r.PathPrefix("/").Handler(http.FileServer(assetFs))

	// Podpięcie mechanizmu logującego
	r.Use(loggingMiddleware)

	log.Println("Starting UI server at", enpoint)

	// Uruchomienie serwera HTTP wo osobnym nie-do-końca-wątku
	go func() {
		e := http.ListenAndServe(enpoint, r)
		log.Panicln(e)
	}()
	return nil
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		session, err := sessionStore.Get(r, "session-name")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
