package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/pradist/hometic/logger"
	"log"
	"net/http"
	"os"
)

type Pair struct {
	DeviceID int64
	UserID   int64
}

func main() {
	if err := run(); err != nil {
		log.Fatal("Can't start application ", err)
	}
}

func run() error {
	fmt.Println("Hello hometic: I'm gopher!!!")
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.Use(logger.Middleware)
	r.Handle("/pair-device", CustomHandleFunc(PairDeviceHandler(NewCreatePairDevice(db)))).Methods(http.MethodPost)

	addr := fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT"))
	fmt.Println("addr: ", addr)
	server := http.Server{
		Addr:    addr,
		Handler: r,
	}

	log.Println("starting...")
	return server.ListenAndServe()
}

type CustomHandleFunc func(w CustomResponseWriter, r *http.Request)

func (handler CustomHandleFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler(&JSONResponseWriter{w}, r)
}

type CustomResponseWriter interface {
	JSON(statusCode int, data interface{})
}

type JSONResponseWriter struct {
	http.ResponseWriter
}

func (w *JSONResponseWriter) JSON(statusCode int, data interface{})  {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func PairDeviceHandler(device Device) func(w CustomResponseWriter, r *http.Request) {
	return func (w CustomResponseWriter, r *http.Request) {
		logger.L(r.Context()).Info("pair-device")

		var p Pair
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			w.JSON(http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		err = device.Pair(p)
		if err != nil {
			w.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		//w.JSON(http.StatusOK, []byte(`{"status":"active"}`))
		w.JSON(http.StatusOK, map[string]interface{}{"status":"active"})
	}
}

type Device interface {
	Pair(p Pair) error
}

type CreatePairDeviceFunc func(p Pair) error

func (fn CreatePairDeviceFunc) Pair(p Pair) error {
	return fn(p)
}

type DB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

func NewCreatePairDevice(db DB) CreatePairDeviceFunc {
	return func(p Pair) error {
		_, err := db.Exec("INSERT INTO pairs VALUES ($1, $2);", p.DeviceID, p.UserID)
		return err
	}
}