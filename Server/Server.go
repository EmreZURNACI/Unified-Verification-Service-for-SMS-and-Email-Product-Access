package server

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

var wg sync.WaitGroup

func Server() {
	_mux := mux.NewRouter()
	wg.Add(2)
	go AuthServer(_mux, &wg)
	go ProductServer(_mux, &wg)
	http.ListenAndServe(":8080", _mux)
	//http.ListenAndServeTLS(":4443", "server.crt", "server.key", _mux)
	wg.Wait()
}

//routerları farklı dosyaya yönlendir +
//verify body ile gönder +
// 1 servis daha yaz servis ile code karsıla
// servise süre ekle dbden yap
//cookieleri uuid göre sil
//sıralama yap
//search ekle
//standart response model tanımla
//idli ürün yoksa patlatma update delete
