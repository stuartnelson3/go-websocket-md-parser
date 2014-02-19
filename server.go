package main

import (
    "github.com/codegangsta/martini"
    "github.com/codegangsta/martini-contrib/render"
    "github.com/russross/blackfriday"
    "github.com/gorilla/websocket"
    "net/http"
)

func main() {
    m := martini.Classic()
    m.Use(render.Renderer(render.Options{
        Layout:     "layout",
        Extensions: []string{".html"}}))

    m.Get("/", func(r render.Render) {
        r.HTML(200, "index", nil)
    })

    m.Get("/markdown_preview", func(w http.ResponseWriter, r *http.Request) {
        ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
        if _, ok := err.(websocket.HandshakeError); ok {
            http.Error(w, "Not a websocket handshake", 400)
            return
        } else if err != nil {
            return
        }

        for {
            messageType, message, err := ws.ReadMessage()
            if err != nil {
                return
            }
            message = ParseMarkdown(message)
            if err := ws.WriteMessage(messageType, message); err != nil {
                return
            }
        }
    })

    m.Run()
}

func ParseMarkdown(markdown []byte) []byte {
    return blackfriday.MarkdownCommon(markdown)
}
