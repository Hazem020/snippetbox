package main

import (
	"context"
	"flag"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"html/template"
	"log"
	"net/http"
	"os"
	"snippetbox.hazem/internal/models"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	snippets      *models.SnippetModel
	templateCache map[string]*template.Template
}

func connectDB() (*mongo.Client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	} else {
		log.Println("Connected to MongoDB!")
	}
	return client, err
}
func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	client, _ := connectDB()
	defer client.Disconnect(context.TODO())
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &models.SnippetModel{
			Client:     client,
			Database:   client.Database("Snippetbox"),
			Collection: client.Database("Snippetbox").Collection("snippets"),
		},
		templateCache: templateCache,
	}
	mux := app.routes()

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}
	infoLog.Printf("Starting server on %s", *addr)
	// Call the ListenAndServe() method on our new http.Server struct.
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
