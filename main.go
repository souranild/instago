package main

import (
	"encoding/json"
	"net/http"
	"regexp"
	"sync"
)

var (
	getuserREGEX     = regexp.MustCompile(`^\/users\/(\d+)$`)
	createuserREGEX  = regexp.MustCompile(`^\/users[\/]*$`)
	getpostREGEX     = regexp.MustCompile(`^\/posts\/(\d+)$`)
	createpostREGEX  = regexp.MustCompile(`^\/posts[\/]*$`)
	getuserpostREGEX = regexp.MustCompile(`^\/posts\/\/users\/(\d+)$`)
)

// User struct
type user struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//Post struct
type post struct {
	ID        string `json:"id"`
	Caption   string `json:"caption"`
	Imgurl    string `json:"imgurl"`
	Timestamp string `json:"timestamp"`
}

//user datastore
type userdatastore struct {
	m map[string]user
	*sync.RWMutex
}

//posts datastore
type postdatastore struct {
	m map[string]post
	*sync.RWMutex
}

//user data storage handler
type userHandler struct {
	store *userdatastore
}

//post data storage handler
type postHandler struct {
	store *postdatastore
}

//User endpoints
func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {

	case r.Method == http.MethodGet && getuserREGEX.MatchString(r.URL.Path):
		h.GetUser(w, r)
		return
	case r.Method == http.MethodPost && createuserREGEX.MatchString(r.URL.Path):
		h.CreateUser(w, r)
		return
	default:
		notFound(w, r)
		return
	}
}

//post endpoints
func (h *postHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {

	case r.Method == http.MethodGet && getpostREGEX.MatchString(r.URL.Path):
		h.GetPost(w, r)
		return
	case r.Method == http.MethodPost && createpostREGEX.MatchString(r.URL.Path):
		h.CreatePost(w, r)
		return
	default:
		notFound(w, r)
		return
	}
}

//Get users By ID
func (h *userHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	matches := getuserREGEX.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		notFound(w, r)
		return
	}
	h.store.RLock()
	u, ok := h.store.m[matches[1]]
	h.store.RUnlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	jsonBytes, err := json.Marshal(u)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

//Create a user
func (h *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u user
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		internalServerError(w, r)
		return
	}
	h.store.Lock()
	h.store.m[u.ID] = u
	h.store.Unlock()
	jsonBytes, err := json.Marshal(u)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

//Get post By ID
func (h *postHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	matches := getpostREGEX.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		notFound(w, r)
		return
	}
	h.store.RLock()
	p, ok := h.store.m[matches[1]]
	h.store.RUnlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("post not found"))
		return
	}
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

//Create a post
func (h *postHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var p post
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		internalServerError(w, r)
		return
	}
	h.store.Lock()
	h.store.m[p.ID] = p
	h.store.Unlock()
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

//Errors
func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}

//Main function
func main() {
	mux := http.NewServeMux()

	//Unit tests
	userH := &userHandler{
		store: &userdatastore{
			m: map[string]user{
				"1": user{ID: "1", Name: "bob", Email: "bob@bob.bob", Password: "pass"},
			},
			RWMutex: &sync.RWMutex{},
		},
	}

	postH := &postHandler{
		store: &postdatastore{
			m: map[string]post{
				"1": post{ID: "1", Caption: "Some Image", Imgurl: "https://image.jpeg", Timestamp: "2019-11-10 02:00:00"},
			},
			RWMutex: &sync.RWMutex{},
		},
	}

	mux.Handle("/users", userH)
	mux.Handle("/users/", userH)
	mux.Handle("/posts", postH)
	mux.Handle("/posts/", postH)

	http.ListenAndServe("localhost:8080", mux)
}
