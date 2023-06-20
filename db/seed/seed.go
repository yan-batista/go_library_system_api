package seed

import (
	"fmt"
	"log"

	"example.com/db"
	"example.com/models"
	"example.com/services"
	"github.com/jaswdr/faker"
)

var connection = db.CreateConnection()

func SeedDB() {
	seedBooks()
	seedUsers()
}

func seedUsers() {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT,
			first_name TEXT NOT NULL,
			last_name TEXT NOT NULL,
			email TEXT NOT NULL,
			phone TEXT NOT NULL,
			debt int NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			PRIMARY KEY (id)
		);
	`

	if _, err := connection.Exec(query); err != nil {
		log.Fatal(err)
	}

	if _, err := connection.Exec("truncate users"); err != nil {
		log.Fatal(err)
	}

	fake := faker.New()
	for i := 0; i < 25; i++ {
		go createUser(fake)
	}
}

func createUser(fake faker.Faker) {
	user := models.UserDTO{FirstName: fake.Person().FirstName(), LastName: fake.Person().LastName(), Email: fake.Person().Contact().Email, Phone: fake.Person().Contact().Phone}
	services.CreateUser(user)
}

func seedBooks() {
	query := `
	CREATE TABLE IF NOT EXISTS books (
		id INT AUTO_INCREMENT,
		name TEXT NOT NULL,
		slug TEXT NOT NULL,
		author TEXT NOT NULL,
		publisher TEXT NOT NULL,
		isbn TEXT NOT NULL,
		quantity INT NOT NULL,
		description TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (id)
	);
`

	// executa a query e verifica se houve algum erro
	if _, err := connection.Exec(query); err != nil {
		log.Fatal(err)
	}

	// limpa o banco
	if _, err := connection.Exec("truncate books"); err != nil {
		log.Fatal(err)
	}

	// cria entradas de livros no banco
	fake := faker.New()
	for i := 0; i < 25; i++ {
		go createBook(fake)
	}
}

func createBook(fake faker.Faker) {
	name := fmt.Sprintf("%s %s", fake.Person().FirstName(), fake.Person().LastName() )

	book := models.BookDTO{Name: fake.App().Name(), Author: name, Publisher: fake.Company().Name(), ISBN: fake.Numerify("###-##########"), Quantity: fake.IntBetween(0,100), Description: fake.Lorem().Sentence(fake.IntBetween(1, 10))}
	services.CreateBook(book)
}