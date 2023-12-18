package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	h "github.com/aakash-rajur/http"
	"github.com/aakash-rajur/http/params"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	env := loadFromOS()

	port := env.Get("PORT", "8080")

	address := ":" + port

	books := []Book{
		{
			Id:          1,
			Name:        "The Alchemist",
			Description: "lorem ipsum",
		},
		{
			Id:          2,
			Name:        "The Monk Who Sold His Ferrari",
			Description: "lorem ipsum",
		},
		{
			Id:          3,
			Name:        "The Subtle Art of Not Giving a F*ck",
			Description: "lorem ipsum",
		},
		{
			Id:          4,
			Name:        "The 5 AM Club",
			Description: "lorem ipsum",
		},
		{
			Id:          5,
			Name:        "The Power of Now",
			Description: "lorem ipsum",
		},
		{
			Id:          6,
			Name:        "The Secret",
			Description: "lorem ipsum",
		},
		{
			Id:          7,
			Name:        "The 7 Habits of Highly Effective People",
			Description: "lorem ipsum",
		},
		{
			Id:          8,
			Name:        "The 4-Hour Workweek",
			Description: "lorem ipsum",
		},
		{
			Id:          9,
			Name:        "The 48 Laws of Power",
			Description: "lorem ipsum",
		},
		{
			Id:          10,
			Name:        "The 10X Rule",
			Description: "lorem ipsum",
		},
	}

	users := []User{
		{
			Id:    1,
			Name:  "Aakash Rajur",
			Email: "aakashrajur@example.com",
		},
		{
			Id:    2,
			Name:  "John Doe",
			Email: "johndoe@example.com",
		},
		{
			Id:    3,
			Name:  "Jane Doe",
			Email: "janedoe@example.com",
		},
	}

	router := h.NewRouter()

	router.Use(h.Logger(h.LoggerConfig{}))

	router.GetFunc(
		"/health",
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)

			_, _ = w.Write([]byte("OK"))
		},
	)

	router.GetFunc(
		"/api/v2/books",
		func(w http.ResponseWriter, r *http.Request) {
			buffer, err := json.Marshal(books)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			w.Header().Set("Content-Type", "application/json")

			w.WriteHeader(http.StatusOK)

			_, _ = w.Write(buffer)
		},
	)

	router.GetFunc(
		"/api/v2/books/:id",
		func(w http.ResponseWriter, r *http.Request) {
			p, ok := params.FromContext(r.Context())

			if !ok {
				http.Error(w, "unable to parse param", http.StatusInternalServerError)

				return
			}

			idString := p.Get("id", "")

			id, err := strconv.Atoi(idString)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)

				return
			}

			book := books[id-1]

			buffer, err := json.Marshal(book)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)

				return
			}

			w.Header().Set("Content-Type", "application/json")

			w.WriteHeader(http.StatusOK)

			_, _ = w.Write(buffer)
		},
	)

	router.GetFunc(
		"/api/v2/users",
		func(w http.ResponseWriter, r *http.Request) {
			buffer, err := json.Marshal(users)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			w.Header().Set("Content-Type", "application/json")

			w.WriteHeader(http.StatusOK)

			_, _ = w.Write(buffer)
		},
	)

	router.GetFunc(
		"/api/v2/users/:id",
		func(w http.ResponseWriter, r *http.Request) {
			p, ok := params.FromContext(r.Context())

			if !ok {
				http.Error(w, "unable to parse param", http.StatusInternalServerError)

				return
			}

			idString := p.Get("id", "")

			id, err := strconv.Atoi(idString)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)

				return
			}

			user := users[id-1]

			buffer, err := json.Marshal(user)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)

				return
			}

			w.Header().Set("Content-Type", "application/json")

			w.WriteHeader(http.StatusOK)

			_, _ = w.Write(buffer)
		},
	)

	router.GetFunc(
		"/api/v2/users/:id/books",
		func(w http.ResponseWriter, r *http.Request) {
			p, ok := params.FromContext(r.Context())

			if !ok {
				http.Error(w, "unable to parse param", http.StatusInternalServerError)

				return
			}

			idString := p.Get("id", "")

			id, err := strconv.Atoi(idString)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)

				return
			}

			user := users[id-1]

			payload := map[string]interface{}{
				"user":  user,
				"books": books,
			}

			buffer, err := json.Marshal(payload)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)

				return
			}

			w.Header().Set("Content-Type", "application/json")

			w.WriteHeader(http.StatusOK)

			_, _ = w.Write(buffer)
		},
	)

	router.GetFunc(
		"/identity",
		func(w http.ResponseWriter, r *http.Request) {
			now := time.Now()

			h := sha256.New()

			h.Write([]byte(now.Format(time.RFC3339)))

			buffer := h.Sum(nil)

			hex := fmt.Sprintf("%x", buffer)

			shortHex := hex[:8]

			path := fmt.Sprintf("/identity/%s", shortHex)

			http.Redirect(w, r, path, http.StatusMovedPermanently)
		},
	)

	router.GetFunc(
		"/identity/:id",
		func(w http.ResponseWriter, r *http.Request) {
			p, ok := params.FromContext(r.Context())

			if !ok {
				http.Error(w, "unable to parse param", http.StatusInternalServerError)

				return
			}

			id := p.Get("id", "")

			w.WriteHeader(http.StatusOK)

			_, _ = w.Write([]byte(id))
		},
	)

	modulusLength := 2048

	hash := crypto.SHA256

	ak, err := rsa.GenerateKey(rand.Reader, modulusLength)

	if err != nil {
		panic(err)
	}

	akPublicBuffer, err := x509.MarshalPKIXPublicKey(ak.Public())

	if err != nil {
		panic(err)
	}

	akPublic := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: akPublicBuffer,
		},
	)

	if err != nil {
		panic(err)
	}

	router.GetFunc(
		"/public",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			w.WriteHeader(http.StatusOK)

			payload := map[string]interface{}{
				"modulusLength": modulusLength,
				"hash":          hash.String(),
				"publicKey":     string(akPublic),
			}

			jsonPayload, err := json.Marshal(payload)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)

				return
			}

			_, _ = w.Write(jsonPayload)
		},
	)

	router.PostFunc(
		"/private",
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)

			buffer, err := io.ReadAll(r.Body)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)

				return
			}

			_ = buffer

			decrypted, err := ak.Decrypt(nil, buffer, &rsa.OAEPOptions{Hash: hash})

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)

				return
			}

			payload := map[string]interface{}{
				"message": string(decrypted),
			}

			jsonPayload, err := json.Marshal(payload)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)

				return
			}

			w.Header().Set("Content-Type", "application/json")

			_, _ = w.Write(jsonPayload)
		},
	)

	certs, err := tls.LoadX509KeyPair("local/app.local.pem", "local/app.local-key.pem")

	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Addr:    address,
		Handler: router,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{certs},
		},
	}

	err = server.ListenAndServeTLS("", "")

	if err != nil {
		panic(err)
	}
}

type Book struct {
	Id          int64  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type User struct {
	Id    int64  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

func loadFromOS() Env {
	env := make(Env)

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)

		env[pair[0]] = pair[1]
	}

	return env
}

type Env map[string]string

func (e Env) Get(key string, defaultValue string) string {
	value, ok := e[key]

	if !ok {
		return defaultValue
	}

	return value
}
