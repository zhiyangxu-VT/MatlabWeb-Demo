package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
    "strconv"
    "os"
    "os/exec"
    //"encoding/json"
    "io"
    //"strings"

    "github.com/gorilla/mux"
)

func main() {
	router_init()
}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func router_init() {
	r := mux.NewRouter()
    r.HandleFunc("/", welcome_request)
    r.HandleFunc("/{cmd}", table_request)

    http.Handle("/", r)
    err := http.ListenAndServe(":3000", nil)
    checkerr(err)
}

func prepare_response(w http.ResponseWriter) http.ResponseWriter {
    w.Header().Set("Access-Control-Allow-Origin", "*")
	return w
}

func welcome_request(w http.ResponseWriter, r *http.Request) {
	w =  prepare_response(w)
    log.Println("Sending welcome massage")
    fmt.Fprintf(w, "{\"message\": \"Welcom to MatlabWeb\"}")
}

func table_request(w http.ResponseWriter, r *http.Request) {
	w = prepare_response(w)
	vars := mux.Vars(r)
	switch r.Method {
	case "GET":
	case "POST":
		file := store_file(r)
        result := matlab_analyse(vars["cmd"], file)
        fmt.Fprintf(w, "%v", result)
	case "PATCH":
	case "DELETE":
	}
}

func store_file(r *http.Request) string {
    file, handle, err := r.FormFile("uploadfile")
    checkerr(err)
    defer file.Close()

    target_name := "uploaded-" + strconv.FormatInt(time.Now().Unix(), 10) + "-" + handle.Filename
    target_folder := "./uploaded_files/"
    target_file := target_folder + target_name
    if _, err := os.Stat(target_folder); os.IsNotExist(err) {
        os.Mkdir(target_folder, 0700)
    }

    f, err := os.OpenFile(target_file, os.O_WRONLY|os.O_CREATE, 0666)
    checkerr(err)
    defer f.Close()

    io.Copy(f, file)
    checkerr(err)
    log.Println("Upload finished")

    return target_file
}

func matlab_analyse(command string, file_name string) string {
    cmd := exec.Command("./bin/"+command, file_name)
    log.Println("Started command")
    result, err := cmd.StdoutPipe()
    err = cmd.Wait()
    log.Println("Command finished")
    checkerr(err)

    return "{\"result\":\"sleeped\", \"" + command + "\":\"" + file_name + "\"}"
}
