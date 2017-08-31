package main

import (
    "fmt"
    "log"
    "net/http"
    "database/sql"
    "encoding/json"
    "io/ioutil"
    "strings"

	_ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "daigou:myDGpasswd32@@/daigou")
	checkerr(err)
	err = db.Ping()
	checkerr(err)
	defer db.Close()
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
    r.HandleFunc("/{table}", table_request)
    r.HandleFunc("/{table}/{id}", records_request)

    http.Handle("/", r)
    err := http.ListenAndServe(":9090", nil)
    checkerr(err)
}

func prepare_response(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	return w
}

func welcome_request(w http.ResponseWriter, r *http.Request) {
	w =  prepare_response(w)
	fmt.Fprintf(w, "Welcom to Daigou")
}

func table_request(w http.ResponseWriter, r *http.Request) {
	w = prepare_response(w)
	vars := mux.Vars(r)
	switch r.Method {
	case "GET":
		
	case "POST":
		data := parse_request_body_single(r)
		fmt.Fprintf(w, "%v", add_record(vars["table"], data))
	case "PATCH":
	case "DELETE":
	}
}

func records_request(w http.ResponseWriter, r *http.Request) {
	w = prepare_response(w)
	vars := mux.Vars(r)
	switch r.Method {
	case "GET":
		if vars["id"] == "all" {
			table_data := get_table(vars["table"])
			table_json, err := json.Marshal(table_data)
			checkerr(err)
			fmt.Fprint(w, string(table_json))
		} else {
			//return filtered records
		}
	case "POST":
	case "PATCH":
		data := parse_request_body_single(r)
		fmt.Fprintf(w, "%v", update_record(vars["table"], vars["id"], data))
	case "DELETE":
		fmt.Fprintf(w, "%v", delete_record(vars["table"], vars["id"]))
	}
}

func parse_request_body_single(r *http.Request) map[string]interface{} {
	body, err := ioutil.ReadAll(r.Body)
    checkerr(err)
    log.Println(body)

	var data_holder interface{}
    err = json.Unmarshal(body, &data_holder)
    checkerr(err)
    data := data_holder.(map[string]interface{})
    log.Println(data)

    return data
}

func get_table(table string) []map[string]string {
	query := "Select * from " + table
	records, err := db.Query(query)
	checkerr(err)
	defer records.Close()
	
	var columns []string
	columns, err = records.Columns()
	checkerr(err)

	colNum := len(columns)
	record := make([]interface{}, colNum)
	for i := range record {
		var data sql.NullString
		record[i] = &data
	}
	var table_data []map[string]string
	for records.Next() {	
		records_m := make(map[string]string)
		records.Scan(record...)
		for i, c := range columns {
			records_m[c] = ""
			nullstring_holder := *(record[i].(*sql.NullString))
			if nullstring_holder.Valid {
				records_m[c] = nullstring_holder.String	
			}
		}
		table_data = append(table_data, records_m)
	}
	checkerr(err)

	return table_data
}

func add_record(table string, data map[string]interface{}) int {
	var keys, values []string
	for key, value := range data {
		keys = append(keys, key)
		values = append(values, "\""+value.(string)+"\"")
	}

	keys_str := "(" + strings.Join(keys, ",") + ")"
	values_str := "(" + strings.Join(values, ",") + ")"

	log.Println(keys_str)
	log.Println(values_str)

	query := "insert into " + table + " " + keys_str + " values " + values_str
	result, err := db.Exec(query)
	checkerr(err)

	var affected int64
	affected, err = result.RowsAffected()
	checkerr(err)

	return int(affected)
}

func update_record(table string, id string, data map[string]interface{}) int {
	var assignments []string
	for key, value := range data {
		assignment := key + "=\""+value.(string)+"\""
		assignments = append(assignments, assignment)
	}
	log.Println(assignments)

	query := "update " + table + " set " + strings.Join(assignments, ", ") + " where id=\"" + id + "\""
	log.Println(query)

	result, err := db.Exec(query)
	checkerr(err)

	var affected int64
	affected, err = result.RowsAffected()
	checkerr(err)
	log.Printf("Rows Affected: %d", int(affected))

	return int(affected)
}

func delete_record(table string, id string) int {
	query := "Delete from " + table + " where id='" + id + "'"
	
	result, err := db.Exec(query)
	checkerr(err)
	affected, err := result.RowsAffected()
	checkerr(err)

	return int(affected)
}
