package main

import (
  "fmt"
  "log"

  "database/sql"
  _ "code.google.com/p/go-sqlite/go1/sqlite3"

  "storedb"
)

const (
  DbDriverName = "sqlite3"
  DbDataSourceName = "store.db"
)

func main() {
  fmt.Print("Setting up db... ")
  db, err := GetAndPingDbConnection(DbDriverName, DbDataSourceName)
  defer db.Close()

  err = InitializeDb(db)
  if err != nil {
    log.Fatal(err)
    return
  }

  fmt.Print("Success!\n")
}

func GetAndPingDbConnection(driverName string, dataSourceName string) (db *sql.DB, err error) {
  db, err = sql.Open(driverName, dataSourceName)
  if err != nil {
    log.Fatal(err)
    return
  }

  err = db.Ping()
  if err != nil {
    log.Fatal(err)
    return
  }

  return
}

func InitializeDb(db *sql.DB) error {
  err := storedb.ConfigureSqlite3(db)
  if err != nil {
    return err
  }

  err = storedb.CreateAllTables(db)
  if err != nil {
    return err
  }

  return nil
}

/*
func registerHandlers() {
  http.HandleFunc("/", RootHandler)
  http.HandleFunc("/error", ErrorHandler)

  http.HandleFunc("/get/", GetHandler)
 
  http.HandleFunc("/set/", SetHandler)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
  err := r.ParseForm()
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "error: %v", err)
    return
  }

  fmt.Fprintf(w, "Hello, %q\n", html.EscapeString(r.URL.Path))
  fmt.Fprintf(w, "%v\n", r.Form)
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusInternalServerError)
  fmt.Fprintf(w, "haha!")
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
  if !ReadParams(w, r) {
    return
  }

  fmt.Printf("%s\n", r.URL)
}

func SetHandler(w http.ResponseWriter, r *http.Request) {
  if !ReadParams(w, r) {
    return
  }

  fmt.Printf("%s\n", r.URL)
}

func ReadParams(w http.ResponseWriter, r *http.Request) (ok bool) {
  ok = true
  err := r.ParseForm()
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    fmt.Fprintf(w, "Could not read parameters\n")
    ok = false
  }

  return
}
*/
