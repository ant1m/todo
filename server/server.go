package server

import (
	"encoding/json"
	"net/http"

	"github.com/ant1m/todo/tasks"
)

type AppContext struct {
	tm *tasks.TaskManager
}

func RunServer() {
	context := &AppContext{tasks.NewTaskManager()}
	http.HandleFunc("/tasks", tasksHandler(context))
	go http.ListenAndServe(":8080", nil)
	println("started")
}

func tasksHandler(context *AppContext) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(context.tm.All())
	}
}
