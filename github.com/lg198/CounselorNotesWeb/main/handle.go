package main

import (
	"github.com/lg198/CounselorNotesWeb/database"
	"encoding/json"
	"net/http"
	"reflect"
	"bytes"
	"io"
)

var handlerMap = map[string]reflect.Type{
	"searchStudents": reflect.TypeOf(SearchStudents{}),
}

func readerToSlice(r io.Reader) []byte {
	buff := new(bytes.Buffer)
	_, err := buff.ReadFrom(r)
	if err != nil {
		logger.Fatalln("Whoops! READ PROBLEM: ", err)
	}
	logger.Println("strrep: ", buff.String())
	return buff.Bytes()
}

type Handler interface {
	Handle() interface{}
}

func HandleWebRequest(w http.ResponseWriter, r *http.Request) {
	var functionRec struct{
		Function string
	}
	umjson := readerToSlice(r.Body)
	json.Unmarshal(umjson, &functionRec)
	logger.Println("FUNCTION REC: ", functionRec.Function)
	handler := reflect.New(handlerMap[functionRec.Function]).Elem().Interface()
	hh := handler.(Handler)
	json.Unmarshal(umjson, &hh)
	respInterface := hh.Handle()
	respBytes, _ := json.Marshal(respInterface)
	w.Write(respBytes)
}

type SearchStudents struct {
	Query string
}

func (s SearchStudents) Handle() interface{} {
	rows, err := database.Datab.Query("SELECT * FROM student WHERE UPPER(firstname || \" \" || lastname) LIKE UPPER(?)", "%"+s.Query+"%")
	if err != nil {
		logger.Panicln(err)
	}
	names := make([]struct{Name, Id string}, 0)
	
	id, firstname, lastname := "", "", ""
	for rows.Next() {
		err := rows.Scan(&id, &firstname, &lastname)
		if err != nil {
			return struct {
				Error string
				Underlying string
			}{
				"Could not search students!",
				err.Error(),
			}
		} else {
			names = append(names, struct{Name, Id string}{firstname + " " + lastname, id})
		}
	}
	return names
}
