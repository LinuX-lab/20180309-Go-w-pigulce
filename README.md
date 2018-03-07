## Samodzielny serwer WWW "w pigułce"

Przykład jednoplikowego serwera WWW, w którym zawartość jest osadzona w pliku wykonywalnym

### Potrzebne pakiety (trzy kropki **są** częścią nazwy pakietu):

```
go get -t -v -u github.com/jteeuwen/go-bindata/...
go get -t -v -u github.com/elazarl/go-bindata-assetfs/...
go get -t -v -u github.com/gorilla/mux
go get -t -v -u github.com/gorilla/sessions
go get -t -v -u github.com/gorilla/websocket
```

Po zmianie zawartości katalogu `www` trzeba przegenerować plik `bindata.go` za pomocą polecenia `go generate`. Mechanizm *generate* opisany jest na [oficjalnym blogu Go](https://blog.golang.org/generate).

