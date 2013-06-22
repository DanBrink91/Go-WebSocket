package main

import(
    "fmt"
    "code.google.com/p/go.net/websocket"
    "flag"
    "log"
    "net/http"
)
var addr = flag.String("addr", ":8080", "http service address")

func homeHandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func main(){
    flag.Parse()
    go h.run()
    http.HandleFunc("/", homeHandler)
    http.Handle("/ws", websocket.Handler(wsHandler))
    if err := http.ListenAndServe(*addr, nil); err != nil {
            log.Fatal("ListenAndServe:", err)
    }
}