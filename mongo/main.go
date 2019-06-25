package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
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
	writeOkResponse := func(msg string) {
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, msg)
	}
	writeErrorResponse := func(err error) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = io.WriteString(w, fmt.Sprintln("Error: ", err))
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	switch request.Method {
	case "GET":
		p, err := h.getPeople()
		if err != nil {
			writeErrorResponse(err)
			return
		}

		var b strings.Builder
		for _, v := range p {
			b.WriteString(fmt.Sprintf("%s %s #%s", v.FirstName, v.LastName, v.Id.Hex()))
			b.WriteString("\n")
		}
		writeOkResponse(b.String())

	case "POST":
		body := make([]byte, request.ContentLength)
		_, err := io.ReadFull(request.Body, body)
		if err != nil {
			writeErrorResponse(err)
			return
		}

		var p Person
		err = json.Unmarshal(body, &p)
		if err != nil {
			writeErrorResponse(err)
			return
		}

		err = h.addPerson(p.FirstName, p.LastName)
		if err != nil {
			writeErrorResponse(err)
			return
		}

		writeOkResponse("New person added.")

	case "DELETE":
		query := request.URL.Query()
		ids, ok := query["id"]
		if ok {
			for _, id := range ids {
				err := h.removePerson(id)
				if err != nil {
					writeErrorResponse(err)
				}
			}

			writeOkResponse("Person(s) removed.")
		} else {
			writeErrorResponse(errors.New("Unknown query parameter(s)"))
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
