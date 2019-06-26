package main

import (
	"context"
	"encoding/json"
	"errors"
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
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	switch request.Method {
	case "GET":
		h.httpGet(w, request)
	case "POST":
		h.httpPost(w, request)
	case "DELETE":
		h.httpDelete(w, request)
	}
}

func (h PeopleHandler) httpGet(w http.ResponseWriter, request *http.Request) {
	p, err := h.getPeople()
	if err != nil {
		serverError(w, err)
		return
	}

	b, err := json.Marshal(p)
	if err != nil {
		serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(b)
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

func (h PeopleHandler) httpPost(w http.ResponseWriter, request *http.Request) {
	body := make([]byte, request.ContentLength)
	_, err := io.ReadFull(request.Body, body)
	if err != nil {
		badRequest(w, err)
		return
	}

	var p Person
	err = json.Unmarshal(body, &p)
	if err != nil {
		badRequest(w, err)
		return
	}

	err = h.addPerson(&p)
	if err != nil {
		serverError(w, err)
		return
	}

	okResponse(w, "New person added.")
}

func (h PeopleHandler) addPerson(p *Person) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if len(p.FirstName) == 0 || len(p.LastName) == 0 {
		err = errors.New("Missing first or last name.")
		return
	}

	_, err = h.people.InsertOne(ctx, bson.M{
		"firstname": p.FirstName,
		"lastname":  p.LastName,
	})

	return
}

func (h PeopleHandler) httpDelete(w http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	ids, ok := query["id"]
	if !ok {
		badRequest(w, errors.New("Missing 'id' query parameter(s)."))
		return
	}

	for _, id := range ids {
		err := h.removePerson(id)
		if err != nil {
			serverError(w, err)
		}
	}

	okResponse(w, "Person(s) removed.")
}

func (h PeopleHandler) removePerson(id string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	oid, _ := primitive.ObjectIDFromHex(id)
	_, err = h.people.DeleteOne(ctx, bson.M{"_id": oid})

	return
}

func okResponse(w http.ResponseWriter, v interface{}) {
	w.WriteHeader(http.StatusOK)

	e := json.NewEncoder(w)
	_ = e.Encode(v)
}

func serverError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)

	e := json.NewEncoder(w)
	_ = e.Encode(fmt.Sprintf("Error: %v", err))
}

func badRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)

	e := json.NewEncoder(w)
	_ = e.Encode(fmt.Sprintf("Error: %v", err))
}
