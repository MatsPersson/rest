package rest

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "testing"
    "strings"
    "os"
    "strconv"
)

var (
    server *httptest.Server
    reader io.Reader
    tasksURL string
    addTaskURL string
    deleteTaskURL string
    listTasksURL string
)

func post(json string, t *testing.T){
    reader = strings.NewReader(json) //Convert string to reader

    request, err := http.NewRequest("POST", addTaskURL, reader) //Create request with JSON body
    request.Header.Set("Content-Type", "application/json")
    res, err := http.DefaultClient.Do(request)

    if err != nil {
        t.Error(err) //Something is wrong while sending request
    }

    if res.StatusCode != 201 {
        t.Errorf("Success expected: %d", res.StatusCode) //Uh-oh this means our test failed
    } 
}

func list(t *testing.T){
    request, err := http.Get(listTasksURL)
    if err != nil {
        t.Error(err) //Something is wrong while sending request
    } else {
        defer request.Body.Close()
        contents, err := ioutil.ReadAll(request.Body)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }
        fmt.Printf("%s\n", string(contents))
    }
}

func put(t *testing.T, id int, json string){   
    reader = strings.NewReader(json) //Convert string to reader
    request, err := http.NewRequest("PUT", tasksURL + "/" + strconv.Itoa(id), reader) //Create request with JSON body
    request.Header.Set("Content-Type", "application/json")
    res, err := http.DefaultClient.Do(request)
    if err != nil {
        t.Error(err) //Something is wrong while sending request
    }

    if res.StatusCode != 201 {
        t.Errorf("Success expected: %d", res.StatusCode) //Uh-oh this means our test failed
    }
}

func deleteTask(t *testing.T, id int){
    _, err := http.NewRequest("DELETE", deleteTaskURL  + strconv.Itoa(id), nil)
    if err != nil {
        t.Error(err) //Something is wrong while sending request
    }
}

func read(t *testing.T, id int){
    request, err := http.Get(tasksURL + "/" + strconv.Itoa(id));
    if err != nil {
        t.Error(err) //Something is wrong while sending request
    } else {
        defer request.Body.Close()
        contents, err := ioutil.ReadAll(request.Body)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }
        fmt.Printf("%s\n", string(contents))
    }
}

func TestCreateTasks(t *testing.T) {
    post("{\"task\":\"Learn golang\",\"duedate\":\"2016-03-11T11:45:26.371Z\", \"completed\": false}", t)
    post("{\"task\":\"golang Rest server\",\"duedate\":\"2016-03-11T11:46:26.371Z\", \"completed\": false}", t)
    list(t);
    read(t, 1)
    read(t, 2)
    deleteTask(t, 1)
    read(t, 1)
    list(t)
    put(t, 2, "{\"task\":\"golang Rest server\",\"duedate\":\"2016-03-11T11:46:26.371Z\", \"completed\": true}")
    list(t)
}

func TestMain(m *testing.M) {
    server = httptest.NewServer(Handlers()) // Creating new server with the user handlers
    tasksURL = fmt.Sprintf("%s/task", server.URL) 
    addTaskURL = tasksURL + "/add"
    deleteTaskURL = tasksURL + "/delete/"
    listTasksURL = tasksURL + "/list"
    exitVal := m.Run()
    os.Exit(exitVal)
}
