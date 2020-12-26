package core

import (
	"github.com/gorilla/mux"
	"github.com/wuwenbao/gcors"
	"log"

	"momoEngine/utils"
	"net/http"
)

func RunServer() {
	utils.Info("run momoEngine api service")
	r := mux.NewRouter()

	r.HandleFunc("/hello", Hello)
	r.HandleFunc("/cookies", HandleCookies)
	r.HandleFunc("/pages", HandlePages)
	r.HandleFunc("/page/{id}", HandlePage)
	r.HandleFunc("/cookies/cache", HandleCookiesCacheCount)
	r.HandleFunc("/login", HandleLogin)

	r.Use(loggingMiddleware)
	// r.Use(authMiddleware)
	// allow CORS
	cors := gcors.New(
		r,
		gcors.WithOrigin("*"),
		gcors.WithMethods("POST, GET, PUT, DELETE, OPTIONS"),
		gcors.WithHeaders("Authorization"),
	)
	utils.Info("listened on " + utils.GetConfig().APIAddr)
	log.Fatal(http.ListenAndServe(utils.GetConfig().APIAddr, cors))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.Info("[ROUTE] " + r.RemoteAddr + " " + r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != "/login" {
			ret := MessageBody{
				Total:   1,
				Message: "not allowed",
				Code:    401,
			}
			// no use bearer
			c, err := r.Cookie("token")
			if err != nil {
				ret.ResponseCommon(w, http.StatusUnauthorized)
				return
			}
			if CheckToken(c.Value) {
				next.ServeHTTP(w, r)
			} else {
				ret.ResponseCommon(w, http.StatusUnauthorized)
				return
			}
		}

	})
}

func Hello(w http.ResponseWriter, r *http.Request) {
	utils.Fail(CreateToken("3gpp"))
}

func HandlePage(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	cl := GetPageLink(v["id"])
	ret := MessageBody{
		Total:   len(cl),
		Data:    cl,
		Message: "ok",
	}
	ret.ResponseCommon(w, http.StatusOK)
}

func HandlePages(w http.ResponseWriter, r *http.Request) {
	cl := GetPages()
	ret := MessageBody{
		Total:   len(cl),
		Data:    cl,
		Message: "ok",
	}
	ret.ResponseCommon(w, http.StatusOK)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var ui UserInfo
	var token string

	ret := MessageBody{
		Total:   1,
		Message: "ok",
		Code:    401,
	}
	if err := utils.JsonBind(&ui, r); err != nil {
		utils.Warn("login params error")
	} else {
		token, err = LoginUser(ui)
		if err != nil {
			utils.Warn("login params error")
		} else {
			ret.Data = token
			ret.Code = 201
		}
	}
	ret.ResponseCommon(w, http.StatusOK)
}

func HandleCookies(w http.ResponseWriter, r *http.Request) {
	cl := CookiesList()
	ret := MessageBody{
		Total:   len(cl),
		Data:    cl,
		Message: "ok",
	}
	ret.ResponseCommon(w, http.StatusOK)
}

func HandleCookiesCacheCount(w http.ResponseWriter, r *http.Request) {
	ret := MessageBody{
		Total:   1,
		Data:    CookiesCount(),
		Message: "ok",
	}
	ret.ResponseCommon(w, http.StatusOK)
}
