package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/bxcodec/faker/v3"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Seed type
type Seed struct {
	db *sqlx.DB
}

func (s Seed) SalesSeed() {

	for i := 0; i < 10; i++ {
		//prepare the statement
		stmt, _ := s.db.Prepare(`INSERT INTO ar_sales(company_code, document_number, status, created_by, updated_by, created_time, last_update, document_date) VALUES ($1, $2, $3, $4, $5, now(), now(), now())`)
		// execute query
		docName := "DOC00"
		concatenated := fmt.Sprintf("%s%d", docName, i)

		_, err := stmt.Exec(faker.Username(), concatenated, 0, faker.Name(), faker.Name())
		if err != nil {
			panic(err)
		}
	}
}

// Execute will executes the given seeder method
func Execute(db *sqlx.DB, seedMethodNames ...string) {
	s := Seed{db}

	seedType := reflect.TypeOf(s)

	// Execute all seeders if no method name is given
	if len(seedMethodNames) == 0 {
		log.Println("Running all seeder...")
		// We are looping over the method on a Seed struct
		for i := 0; i < seedType.NumMethod(); i++ {
			// Get the method in the current iteration
			method := seedType.Method(i)
			// Execute seeder
			seed(s, method.Name)
		}
	}

	// Execute only the given method names
	for _, item := range seedMethodNames {
		seed(s, item)
	}
}

func seed(s Seed, seedMethodName string) {
	// Get the reflect value of the method
	m := reflect.ValueOf(s).MethodByName(seedMethodName)
	// Exit if the method doesn't exist
	if !m.IsValid() {
		log.Fatal("No method called ", seedMethodName)
	}
	// Execute the method
	log.Println("Seeding", seedMethodName, "...")
	m.Call(nil)
	log.Println("Seed", seedMethodName, "succedd")
}

func handleArgs() error {
	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			datasource := fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable",
				os.Getenv("DB_DIALECT"),
				os.Getenv("DB_USERNAME"),
				os.Getenv("DB_PASSWORD"),
				os.Getenv("DB_HOST"),
				os.Getenv("DB_NAME"))
			db, err := sqlx.Connect(os.Getenv("DB_DIALECT"), datasource)
			if err != nil {
				log.Fatalf("Error opening DB: %v", err)
				return err
			}
			Execute(db, args[1:]...)
			os.Exit(0)
		}
	}
	fmt.Println("Seeding Success..!")
	return nil
}

func main() {
	fmt.Println("Loading ENV...")
	godotenv.Load()

	fmt.Println("Seeding DB...")
	err := handleArgs()
	if err != nil {
		log.Fatalf("Error seeding DB: %v", err)
	}
}
