package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	FirstName string             `json:"firstname,omitempty"`
	LastName  string             `json:"lastname,omitempty"`
}

type PeopleHandler struct {
	people *mongo.Collection
}

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

func (h PeopleHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	okResponse := func(v interface{}) {
		w.WriteHeader(http.StatusOK)
		e := json.NewEncoder(w)
		_ = e.Encode(v)
	}
	errorResponse := func(v interface{}) {
		w.WriteHeader(http.StatusInternalServerError)
		e := json.NewEncoder(w)
		_ = e.Encode(v)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	switch request.Method {
	case "GET":
		p, err := h.getPeople()
		if err != nil {
			errorResponse(err)
			return
		}

		b, err := json.Marshal(p)
		if err != nil {
			errorResponse(err)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(b)

	case "POST":
		body := make([]byte, request.ContentLength)

		_, err := io.ReadFull(request.Body, body)
		if err != nil {
			errorResponse(err)
			return
		}

		var p Person

		err = json.Unmarshal(body, &p)
		if err != nil {
			errorResponse(err)
			return
		}

		err = h.addPerson(p.FirstName, p.LastName)
		if err != nil {
			errorResponse(err)
			return
		}

		okResponse("New person added.")

	case "DELETE":
		query := request.URL.Query()
		ids, ok := query["id"]
		if ok {
			for _, id := range ids {
				err := h.removePerson(id)
				if err != nil {
					errorResponse(err)
				}
			}

			okResponse("Person(s) removed.")
		} else {
			errorResponse(errors.New("Unknown query parameter(s)"))
		}
	}
}

func (h PeopleHandler) getPeople() (people []Person, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := h.people.Find(ctx, bson.D{})
	if err != nil {
		return
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var p Person

		err = cur.Decode(&p)
		if err != nil {
			return
		}

		people = append(people, p)
	}

	return
}

func (h PeopleHandler) addPerson(firstname, lastname string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = h.people.InsertOne(ctx, bson.M{
		"firstname": firstname,
		"lastname":  lastname,
	})

	return
}

func (h PeopleHandler) removePerson(id string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	oid, _ := primitive.ObjectIDFromHex(id)
	_, err = h.people.DeleteOne(ctx, bson.M{"_id": oid})

	return
}
