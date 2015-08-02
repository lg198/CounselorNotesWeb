package main

import (
	"log"
	"os"
	"github.com/lg198/CounselorNotesWeb/database"
	"net/http"
	"bytes"
	"io/ioutil"
	"strings"
	"mime"
)

var logger *log.Logger


func echoHandler(w http.ResponseWriter, r *http.Request) {
	buff := new(bytes.Buffer)
	buff.ReadFrom(r.Body)
	buff.WriteTo(w)
}

func ResourceHandler(w http.ResponseWriter, r *http.Request) {
	realPath := strings.TrimPrefix(r.URL.Path, "/web/")
	contents, err := ioutil.ReadFile("../web/" + realPath)
	extension := realPath[strings.LastIndex(realPath, "."):]
	logger.Println(realPath, " ", extension, " ", mime.TypeByExtension(extension))
	if err != nil {
		logger.Fatalln(err)
	}
	w.Header().Set("Content-Type", mime.TypeByExtension(extension))
	w.Write(contents)
}

func main() {
	logger = log.New(os.Stdout, "[CounselorNotesWeb]:", 0)

	database.InitDB(logger)
	defer database.Datab.Close()
	
	preped, err := database.Datab.Prepare("INSERT OR REPLACE INTO student VALUES (?, ?, ?)")
	if err != nil {
		logger.Fatalln(err.Error())
	}
	preped.Exec("1", "Luke", "Seymour")
	preped.Exec("2", "Jason", "Barnes")
	
	http.HandleFunc("/handle", HandleWebRequest)
	http.HandleFunc("/test", echoHandler)
	http.HandleFunc("/web/", ResourceHandler)
    http.ListenAndServe(":25678", nil)
}
