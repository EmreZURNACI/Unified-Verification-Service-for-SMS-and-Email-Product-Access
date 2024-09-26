package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	d "Datadb.go"
	helpers "Helpers.go"
	m "Models.go"
	"github.com/gorilla/mux"
)

func ProductServer(_mux *mux.Router) {
	dataRouter := _mux.PathPrefix("/product").Subrouter()
	dataRouter.Use(IsSetTokenMiddleware)
	dataRouter.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		if http.MethodGet == r.Method {
			var bdy m.BodyProductRes
			if r.Body != http.NoBody && r.Body != nil {
				bs, err := ioutil.ReadAll(r.Body)
				if err != nil {
					panic(err)
				}
				err = json.Unmarshal(bs, &bdy)
				if err != nil {
					panic(err)
				}
				fmt.Fprintf(w, "%+v", helpers.Response(true, 200, "Veriler başarıyla getirildi.", d.Products(bdy)))
			} else {
				fmt.Fprintf(w, "%+v", helpers.Response(true, 200, "Veriler başarıyla getirildi.", d.Products(bdy)))
			}
		} else {
			fmt.Fprintf(w, "%+v", helpers.Response(false, 405, "Bu URL için kullanılan HTTP yöntemi desteklenmiyor", nil))
		}
	})
	dataRouter.HandleFunc("/product-{id}", func(w http.ResponseWriter, r *http.Request) {
		if http.MethodGet == r.Method {
			vars := mux.Vars(r)
			id := vars["id"]
			number_id, err := strconv.Atoi(id)
			if err != nil {
				fmt.Fprintln(w, helpers.Response(false, 400, "Girilen değer int'e çevrilememektedir.", nil))
			} else {
				if d.GetData(number_id)[0].Id != 0 {
					fmt.Fprintln(w, helpers.Response(true, 200, "Veri başarıyla getirildi.", d.GetData(number_id)))
				} else {
					fmt.Fprintln(w, helpers.Response(false, 400, "Bu ID değerine sahip kayıt yoktur.", nil))
				}
			}
		} else {
			fmt.Fprintf(w, "%+v", helpers.Response(false, 405, "Bu URL için kullanılan HTTP yöntemi desteklenmiyor", nil))
		}
	})
	dataRouter.HandleFunc("/createproduct", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			if r.Body != http.NoBody && r.Body != nil {
				var p m.Product
				bs, err := ioutil.ReadAll(r.Body)
				if err != nil {
					panic(err)
				}
				err = json.Unmarshal(bs, &p)
				if err != nil {
					panic(err)
				}
				statu, message := d.CreateProduct(p.Marka, p.Model, p.IsletimSistemi)
				if statu {
					fmt.Fprintln(w, helpers.Response(true, 200, message, nil))
				} else {
					fmt.Fprintln(w, helpers.Response(false, 400, message, nil))
				}
			} else {
				fmt.Fprintf(w, "%+v", helpers.Response(false, 400, "Request body boş olamaz.", nil))
			}
		} else {
			fmt.Fprintf(w, "%+v", helpers.Response(false, 405, "Bu URL için kullanılan HTTP yöntemi desteklenmiyor", nil))
		}
	})
	dataRouter.HandleFunc("/deleteproduct", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			if r.Body != http.NoBody && r.Body != nil {
				var p m.Product
				bs, err := ioutil.ReadAll(r.Body)
				if err != nil {
					panic(err)
				}
				err = json.Unmarshal(bs, &p)
				if err != nil {
					panic(err)
				}
				statu, message := d.DeleteProduct(p.Id)
				if statu {
					fmt.Fprintln(w, helpers.Response(true, 200, message, nil))
				} else {
					fmt.Fprintln(w, helpers.Response(false, 400, message, nil))
				}
			} else {
				fmt.Fprintf(w, "%+v", helpers.Response(false, 400, "Request body boş olamaz.", nil))
			}
		} else {
			fmt.Fprintf(w, "%+v", helpers.Response(false, 405, "Bu URL için kullanılan HTTP yöntemi desteklenmiyor", nil))
		}
	})
	dataRouter.HandleFunc("/updateproduct", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			if r.Body != http.NoBody && r.Body != nil {
				var p m.Product
				bs, err := ioutil.ReadAll(r.Body)
				if err != nil {
					panic(err)
				}
				err = json.Unmarshal(bs, &p)
				if err != nil {
					panic(err)
				}
				statu, message := d.UpdateProduct(p.Id, p.Marka, p.Model, p.IsletimSistemi)
				if statu {
					fmt.Fprintln(w, helpers.Response(true, 200, message, nil))
				} else {
					fmt.Fprintln(w, helpers.Response(false, 400, message, nil))
				}
			} else {
				fmt.Fprintf(w, "%+v", helpers.Response(false, 400, "Request body boş olamaz.", nil))
			}
		} else {
			fmt.Fprintf(w, "%+v", helpers.Response(false, 405, "Bu URL için kullanılan HTTP yöntemi desteklenmiyor", nil))
		}
	})
}

//middleware ==> Cookie kontrolu için
