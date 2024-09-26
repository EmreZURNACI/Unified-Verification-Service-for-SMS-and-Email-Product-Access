package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Server() {
	_mux := mux.NewRouter()
	AuthServer(_mux)
	ProductServer(_mux)
	http.ListenAndServe(":8080", _mux)
	//http.ListenAndServeTLS(":4443", "server.crt", "server.key", _mux)
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
