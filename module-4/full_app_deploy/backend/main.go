package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Data struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
	Email   string `json:"email"`
}

func main() {
	db, err := sql.Open("postgres", "postgresql://admin:admin@postgres-conn:5432/bootcamp?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		var data Data
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			log.Fatal(err)
		}

		conn, err := db.Conn(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		defer conn.Close()

		id := uuid.NewString()

		if _, err := conn.ExecContext(
			context.Background(),
			"INSERT INTO comments (id, name, comment, email) VALUES ($1, $2, $3, $4)",
			id, data.Name, data.Comment, data.Email,
		); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"data":    err,
			})
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true,
		})
	})

	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}
