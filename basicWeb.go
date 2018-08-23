package main
import
(
  "fmt"
  "log"
  "net/http"
)

func darMensaje(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Laboratory #1")
}
func darSaludo(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hola Juan")
}

func main() {
  http.HandleFunc("/olvera", darMensaje)
  http.HandleFunc("/juan", darSaludo)
  err := http.ListenAndServe("localhost"+":"+"8080", nil)
  if err != nil {
    log.Fatal("error en el servidor : ", err)
    return
  }


}
