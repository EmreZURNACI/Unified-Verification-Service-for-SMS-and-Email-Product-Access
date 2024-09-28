package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	a "Auth.go"
	h "Helpers.go"
	m "Models.go"
	"github.com/gorilla/mux"
)

func AuthServer(_mux *mux.Router, _wg *sync.WaitGroup) {
	authRouter := _mux.PathPrefix("/auth").Subrouter()
	authRouter.Use(IsDatabaseConnected)
	authRouter.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		if http.MethodPost == r.Method {
			if r.Body != http.NoBody && r.Body != nil {
				var u m.BodyAuthRes
				bs, err := ioutil.ReadAll(r.Body)
				if err != nil {
					panic(err)
				}
				err = json.Unmarshal(bs, &u)
				if err != nil {
					panic(err)
				}
				statu, message := a.Signin(u.Email, u.Password)
				if statu {
					c, err := r.Cookie("Token")
					if err != nil {
						c = &http.Cookie{
							Name:     "Token",
							Value:    "Token",
							Path:     "/",
							HttpOnly: false,
							Secure:   false,
							MaxAge:   3600,
							Expires:  time.Now().Add(3600 * time.Second),
						}
						http.SetCookie(w, c)
					}

					fmt.Fprintln(w, h.Response(true, 200, message, nil))
				} else {
					fmt.Fprintln(w, h.Response(false, 400, message, nil))
				}
			} else {
				fmt.Fprintf(w, "%+v", h.Response(false, 400, "Request body boş olamaz.", nil))
			}
		} else {
			fmt.Fprintf(w, "%+v", h.Response(false, 405, "Bu URL için kullanılan HTTP yöntemi desteklenmiyor", nil))
		}
	})
	authRouter.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		if http.MethodPost == r.Method {
			if r.Body != http.NoBody && r.Body != nil {
				var u m.BodyAuthRes
				bs, err := ioutil.ReadAll(r.Body)
				if err != nil {
					panic(err)
				}
				err = json.Unmarshal(bs, &u)
				if err != nil {
					panic(err)
				}
				statu, message := a.Signup(u.Email, u.Name, u.Lastname, u.Nickname, u.Password, u.Tel)
				if statu {
					fmt.Fprintln(w, h.Response(true, 200, message, nil))
				} else {
					fmt.Fprintln(w, h.Response(false, 400, message, nil))
				}
			} else {
				fmt.Fprintf(w, "%+v", h.Response(false, 400, "Request body boş olamaz.", nil))
			}
		} else {
			fmt.Fprintf(w, "%+v", h.Response(false, 405, "Bu URL için kullanılan HTTP yöntemi desteklenmiyor", nil))
		}
	})
	authRouter.HandleFunc("/verification", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			if r.Body != http.NoBody && r.Body != nil {
				var u m.BodyAuthRes
				bs, err := ioutil.ReadAll(r.Body)
				if err != nil {
					panic(err)
				}
				err = json.Unmarshal(bs, &u)
				if err != nil {
					panic(err)
				}
				statu, message := a.IsAccountVerified(u.Tel, u.Email)
				if statu {
					fmt.Fprintf(w, "%+v", h.Response(true, 200, message, nil))
				} else {
					statu, message := a.Verifications(u.Email, u.Tel, u.Verifytype)
					if statu {
						fmt.Fprintf(w, "%+v", h.Response(true, 200, message, nil))
					} else {
						fmt.Fprintf(w, "%+v", h.Response(false, 400, message, nil))
					}
				}
			} else {
				fmt.Fprintf(w, "%+v", h.Response(false, 400, "Request body boş olamaz.", nil))
			}
		} else {
			fmt.Fprintf(w, "%+v", h.Response(false, 405, "Bu URL için kullanılan HTTP yöntemi desteklenmiyor", nil))
		}

	})
	authRouter.HandleFunc("/verifycode", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			if r.Body != http.NoBody && r.Body != nil {
				var c m.VerifyCode
				bs, err := ioutil.ReadAll(r.Body)
				if err != nil {
					panic(err)
				}
				err = json.Unmarshal(bs, &c)
				if err != nil {
					panic(err)
				}
				statu, message := a.VerifyCode(c.Code)
				if statu {
					fmt.Fprintf(w, "%+v", h.Response(false, 400, message, nil))
				} else {
					fmt.Fprintf(w, "%+v", h.Response(true, 200, message, nil))
				}

			} else {
				fmt.Fprintf(w, "%+v", h.Response(false, 400, "Request body boş olamaz.", nil))
			}
		} else {
			fmt.Fprintf(w, "%+v", h.Response(false, 405, "Bu URL için kullanılan HTTP yöntemi desteklenmiyor", nil))
		}
	})
	authRouter.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("Token")
		if err == nil {
			c := &http.Cookie{
				Name:     "Token",
				Value:    "Token",
				Path:     "/",
				HttpOnly: true,
				Secure:   false,
				Expires:  time.Now().Add(-10000000),
				MaxAge:   -1,
			}
			http.SetCookie(w, c)
		}
		fmt.Fprintln(w, "Cookie'ler temizlendi.")
	})
	_wg.Done()
}
