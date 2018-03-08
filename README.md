# Samodzielny serwer WWW "w pigułce"

Przykład jednoplikowego serwera WWW, w którym zawartość jest osadzona w pliku wykonywalnym

## Potrzebne pakiety GO (trzy kropki **są** częścią nazwy pakietu):

Opcje:
* `-t` - instaluje też ewentualne pakiety zależne używane tylko do uruchamiania testów
* `-u` - jeżeli pakiet jest już zainstalowany, to go aktualizauje albo reinstaluje
* `-v` - tryb gadatliwy, wyświetka wykonywane czynności (np instalacja pakietów zależnych)

```
go get -t -v -u github.com/jteeuwen/go-bindata/...
go get -t -v -u github.com/elazarl/go-bindata-assetfs/...
go get -t -v -u github.com/gorilla/mux
go get -t -v -u github.com/gorilla/sessions
go get -t -v -u github.com/gorilla/websocket
```

Po zmianie zawartości katalogu `www` trzeba przegenerować plik `bindata.go` za pomocą polecenia `go generate`. Mechanizm *generate* opisany jest na [oficjalnym blogu Go](https://blog.golang.org/generate).

## Kompilacja

Uruchomić `go build`.
