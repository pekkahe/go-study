package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	handler := PeopleHandler{client.Database("pekkadb").Collection("people")}

	http.Handle("/people", handler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

type PeopleHandler struct {
	people *mongo.Collection
}

func (h PeopleHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	switch request.Method {
	case "GET":
		p, err := h.getPeople()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = io.WriteString(w, fmt.Sprintln("Error: ", err))
			return
		}

		w.WriteHeader(http.StatusOK)
		for _, v := range p {
			_, _ = io.WriteString(w, fmt.Sprintln(v))
		}

	case "POST":
		body := make([]byte, request.ContentLength)

		_, err := io.ReadFull(request.Body, body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = io.WriteString(w, fmt.Sprintln("Error: ", err))
			return
		}

		var jsonv map[string]interface{}

		err = json.Unmarshal(body, &jsonv)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = io.WriteString(w, fmt.Sprintln("Error: ", err))
			return
		}

		firstname, _ := jsonv["firstname"].(string)
		lastname, _ := jsonv["lastname"].(string)

		err = h.addPerson(firstname, lastname)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = io.WriteString(w, fmt.Sprintln("Error: ", err))
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, "New person added.")

	case "DELETE":
		query := request.URL.Query()
		ids, ok := query["id"]
		if ok {
			for _, id := range ids {
				err := h.removePerson(id)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					_, _ = io.WriteString(w, fmt.Sprintln("Error: ", err))
				}
			}

			w.WriteHeader(http.StatusOK)
			_, _ = io.WriteString(w, "Person(s) removed.")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = io.WriteString(w, "Error: Unknown request")
		}
	}
}

func (h PeopleHandler) getPeople() ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := h.people.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var p []bson.M

	for cur.Next(ctx) {
		var person bson.M

		err := cur.Decode(&person)
		if err != nil {
			return nil, err
		}

		p = append(p, person)
	}

	return p, nil
}

func (h PeopleHandler) addPerson(firstname, lastname string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := h.people.InsertOne(ctx, bson.M{
		"firstname": firstname,
		"lastname":  lastname,
	})
	if err != nil {
		return err
	}

	return nil
}

func (h PeopleHandler) removePerson(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = h.people.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}

	return nil
}
