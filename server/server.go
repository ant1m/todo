package server

import (
	"encoding/json"
	"net/http"

	"github.com/ant1m/todo/tasks"
	"github.com/gorilla/mux"
)

type AppContext struct {
	tm *tasks.TaskManager
}

func RunServer() {
	context := &AppContext{tasks.NewTaskManager()}
	r := mux.NewRouter()
	r.HandleFunc("/tasks", tasksHandler(context))
	http.Handle("/tasks", r)
	go http.ListenAndServe(":8080", nil)
	println("started")
}

func tasksHandler(context *AppContext) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(context.tm.All())
	}
}
