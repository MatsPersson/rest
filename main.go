package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
    "time"
	"github.com/gorilla/mux"
    "strconv"
)

// ToDo task
type ToDo struct {
	Task  string `json:"task"`
	DueDate time.Time `json:"duedate"`
	Completed bool `json:"completed"`
}

var tasks = make(map[string]*ToDo)
var taskCt = 0

func Handlers() *mux.Router {
    router := mux.NewRouter()
	router.HandleFunc("/task/list", handleTasks).Methods("GET")
    router.HandleFunc("/task/add", handlePost).Methods("POST")
	router.HandleFunc("/task/delete/{id}", handleTask).Methods("DELETE")
	router.HandleFunc("/task/{id}", handleTask).Methods("GET", "PUT")
    return router
}

func main() {
	http.ListenAndServe(":8080", Handlers())
}

func handlePost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
    taskCt++
    var id = strconv.Itoa(taskCt)
	task := new(ToDo)
	decoder := json.NewDecoder(req.Body)
	error := decoder.Decode(&task)
	if error != nil {
	    log.Println(error.Error())
		http.Error(res, error.Error(), http.StatusInternalServerError)
		return
	}
	tasks[id] = task
	outgoingJSON, err := json.Marshal(task)
	if err != nil {
		log.Println(error.Error())
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusCreated)
	fmt.Fprint(res, string(outgoingJSON))
}

func handleTask(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
    id := vars["id"]
 	if id == "" {
		res.WriteHeader(http.StatusNotFound)
		fmt.Fprint(res, string("Id of task not specified or not integer"))
	}   
	switch req.Method {
	case "GET":
		task, ok := tasks[id]
		if !ok {
			res.WriteHeader(http.StatusNotFound)
			fmt.Fprint(res, string("Task not found"))
		}
		outgoingJSON, error := json.Marshal(task)
		if error != nil {
			log.Println(error.Error())
			http.Error(res, error.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(res, string(outgoingJSON))
	case "DELETE":
		delete(tasks, id)
		res.WriteHeader(http.StatusNoContent)
    case "PUT":
		task := new(ToDo)
		decoder := json.NewDecoder(req.Body)
		error := decoder.Decode(&task)
		if error != nil {
			log.Println(error.Error())
			http.Error(res, error.Error(), http.StatusInternalServerError)
			return
		}
		tasks[id] = task
		outgoingJSON, err := json.Marshal(task)
		if err != nil {
			log.Println(error.Error())
			http.Error(res, err.Error(), http.StatusInternalServerError)
		} else {
            res.WriteHeader(http.StatusCreated)
		    fmt.Fprint(res, string(outgoingJSON))
        }
	}
}

func handleTasks(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	outgoingJSON, error := json.Marshal(tasks)
	if error != nil {
		log.Println(error.Error())
		http.Error(res, error.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(res, string(outgoingJSON))
}