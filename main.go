package main

import (
  "encoding/json"
  "log"
  "net/http"

  "github.com/gorilla/mux"
  "winks/plugins/cpu"
  "winks/plugins/mem"
)

func GetCpu(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(cpu.Get())
}

func GetMem(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(mem.Get())
}

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/api/3/cpu", GetCpu).Methods("GET")
  router.HandleFunc("/api/3/mem", GetMem).Methods("GET")
  log.Fatal(http.ListenAndServe(":61208", router))
}
