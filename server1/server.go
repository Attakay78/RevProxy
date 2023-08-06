package server1

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Item struct{
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
}

type Server struct{
	*mux.Router
	shoppingItems []Item
}

func NewServer() *Server{
	s := &Server{
		Router: mux.NewRouter(),
		shoppingItems: []Item{},
	}
	s.routes()
	return s
}

func (s *Server) routes(){
	s.HandleFunc("/shopping-items", s.listShoppingItems()).Methods("GET")
	s.HandleFunc("/shopping-items", s.createShoppingItem()).Methods("POST")
	s.HandleFunc("/shopping-items/{id}", s.removeShoppingItem()).Methods("DELETE")
}

func (s *Server) createShoppingItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		var i Item
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		i.ID = uuid.New()
		s.shoppingItems = append(s.shoppingItems, i)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(i); err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) listShoppingItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.shoppingItems); err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) removeShoppingItem() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		idStr := mux.Vars(r)["id"]
		id, err := uuid.Parse(idStr)

		if err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		for i, item := range s.shoppingItems{
			if item.ID == id {
				s.shoppingItems = append(s.shoppingItems[:i], s.shoppingItems[i+1:]...)
				break
			}
		}
	}
}