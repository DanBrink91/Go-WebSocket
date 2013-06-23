package main

import(
    "code.google.com/p/go.net/websocket"
    "flag"
    "log"
    "net/http"
    "html/template"
)

var addr = flag.String("addr", ":8080", "http service address")
var homeTempl = template.Must(template.ParseFiles("test.html"))

func homeHandler(w http.ResponseWriter, r *http.Request){
    homeTempl.Execute(w, r.Host)
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