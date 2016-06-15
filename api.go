package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mies/todo-api/todo"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type ApiService struct {
	Port       string
	TodoHost   string
	TodoPort   string
	Connection grpc.ClientConn
}

func NewApiService(port string, todoHost string, todoPort string) (*ApiService, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", todoHost, todoPort), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &ApiService{port, todoHost, todoPort, *conn}, nil
}

func (s ApiService) ListTodosHandler(w http.ResponseWriter, r *http.Request) {
	/*
		conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		defer conn.Close()
	*/
	client := todo.NewDoSomethingClient(&s.Connection)
	result, err := client.ListTodos(context.Background(), &todo.Empty{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	b, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s ApiService) NewTodoHandler(w http.ResponseWriter, r *http.Request) {

	/*
	   conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	   if err != nil {
	       http.Error(w, err.Error(), http.StatusInternalServerError)
	   }
	   defer conn.Close()
	*/
	client := todo.NewDoSomethingClient(&s.Connection)

	var todo todo.Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, err = client.AddTodo(context.Background(), &todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"success"}`))
}

func (s ApiService) StartServer() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/todos", s.ListTodosHandler)
	router.HandleFunc("/new", s.NewTodoHandler)
	return router
}
