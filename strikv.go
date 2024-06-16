package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/dgraph-io/badger/v4"
)

var db *badger.DB

func handler(w http.ResponseWriter, r *http.Request) {
	variantKey := r.PathValue("key")
	if r.Method == "GET" {
		log.Printf("GET %s", variantKey)
		lsm_size, vlog_size := db.Size()
		log.Printf("LSM: %d, VLOG: %d", lsm_size, vlog_size)
		txn := db.NewTransaction(false)
		defer txn.Discard()
		log.Print("Getting item")
		item, err := txn.Get([]byte(variantKey))
		if err == badger.ErrKeyNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Print("Found item")
		err = item.Value(func(val []byte) error {
			w.Write(val)
			return nil
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	} else if r.Method == "PUT" {
		txn := db.NewTransaction(true)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = txn.Set([]byte(variantKey), []byte(body))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		txn.Commit()
		w.WriteHeader(http.StatusCreated)
		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func initDb() *badger.DB {
	dbPath := os.Getenv("STRIKV_PATH")
	if dbPath == "" {
		dbPath = "/tmp/badger"
	}
	db, err := badger.Open(badger.DefaultOptions(dbPath))
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	db = initDb()
	defer db.Close()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Print("Closing DB")
		db.Close()
		os.Exit(0)
	}()

	lsm_size, vlog_size := db.Size()
	log.Printf("LSM: %d, VLOG: %d", lsm_size, vlog_size)
	http.HandleFunc("/{key}", handler)
	log.Fatal(http.ListenAndServe(":9080", nil))
}
