package main
import "net/http"

func main(){
	http.HandleFunc("/favicon.ico", getFavicon)
	http.HandleFunc("/", getMain)
	http.HandleFunc("/pravila", getPr)

	http.ListenAndServe(":8080", nil)
}

func getMain(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "index.html")
}

func getPr(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "pravila.html")
}

func getFavicon(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "favicon.ico")
}
