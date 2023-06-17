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
	query := `
		CREATE TABLE IF NOT EXISTS books (
			id INT AUTO_INCREMENT,
			name TEXT NOT NULL,
			slug TEXT NOT NULL,
			author TEXT NOT NULL,
			publisher TEXT NOT NULL,
			isbn TEXT NOT NULL,
			quantity INT NOT NULL,
			created_at DATETIME,
			updated_at DATETIME,
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

	// cria entradas no banco
	fake := faker.New()
	for i := 0; i < 25; i++ {
		go createBook(fake)
	}
}

func createBook(fake faker.Faker) {
	name := fmt.Sprintf("%s %s", fake.Person().FirstName(), fake.Person().LastName() )

	book := models.BookDTO{Name: fake.App().Name(), Author: name, Publisher: fake.Company().Name(), ISBN: fake.Numerify("###-##########"), Quantity: fake.IntBetween(0,100)}
	services.CreateBook(book)
}