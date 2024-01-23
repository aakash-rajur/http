package main

import (
	"crypto"
	"crypto/sha256"
	"crypto/tls"
	"encoding/json"
	"fmt"
	h "github.com/aakash-rajur/http"
	"github.com/aakash-rajur/http/params"
	"io"
	"net/http"
	"strconv"
	"time"
)

func main() {
	env := loadFromOS()

	port := env.Get("PORT", "8080")

	address := ":" + port

	router := h.NewRouter()

	router.Use(h.Logger(h.LoggerConfig{}))

	router.GetFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)

			_, _ = w.Write([]byte("Hello World!"))
		},
	)

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
		"/api/v2/books/{id}",
		func(w http.ResponseWriter, r *http.Request) {
			p, ok := params.FromRequest(r)

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
		"/api/v2/users/{id}",
		func(w http.ResponseWriter, r *http.Request) {
			p, ok := params.FromRequest(r)

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
		"/api/v2/users/{id}/books",
		func(w http.ResponseWriter, r *http.Request) {
			p, ok := params.FromRequest(r)

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

			enc := sha256.New()

			enc.Write([]byte(now.Format(time.RFC3339)))

			buffer := enc.Sum(nil)

			hex := fmt.Sprintf("%x", buffer)

			shortHex := hex[:8]

			path := fmt.Sprintf("/identity/%s", shortHex)

			http.Redirect(w, r, path, http.StatusMovedPermanently)
		},
	)

	router.GetFunc(
		"/identity/{id}",
		func(w http.ResponseWriter, r *http.Request) {
			p, ok := params.FromRequest(r)

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

	keyPair, err := generateRSAKeyPair(hash, modulusLength)

	router.GetFunc(
		"/public",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			w.WriteHeader(http.StatusOK)

			payload := map[string]interface{}{
				"modulusLength": modulusLength,
				"hash":          hash.String(),
				"publicKey":     keyPair.PublicKeyBytes,
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

			decrypted, err := keyPair.Decrypt(buffer)

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

	router.GetFunc(
		"/settings.json",
		func(w http.ResponseWriter, r *http.Request) {
			keyPair, err := generateRSAKeyPair(hash, modulusLength)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)

				return
			}

			// send the public key to the client along with existing settings object

			w.Header().Set("Content-Type", "application/json")

			w.WriteHeader(http.StatusOK)

			payload := map[string]interface{}{
				"modulusLength": modulusLength,
				"hash":          hash.String(),
				"publicKey":     keyPair.PublicKeyBytes,
				"other":         "stuff",
			}

			jsonPayload, err := json.Marshal(payload)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)

				return
			}

			_, _ = w.Write(jsonPayload)

			// store the keypair somewhere for setup-session to pickup
		},
	)

	certFile := env.Get("CERT_FILE", "")

	keyFile := env.Get("KEY_FILE", "")

	certs, err := tls.LoadX509KeyPair(certFile, keyFile)

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
