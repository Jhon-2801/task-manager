package routers

import "net/http"

func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("¡Hola, mundo!"))
}
