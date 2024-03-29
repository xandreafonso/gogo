package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/xandreafonso/gogo/internal/application/usecase"
	"github.com/xandreafonso/gogo/internal/infra/database"
	rabbitmq "github.com/xandreafonso/gogo/pkg"
)

type Car struct {
	Model string
	Color string
}

func (c Car) Start() {
	println(c.Model + " is started")
}

// func funcaoComum(x, y int) int {
// 	return x + y
// }

// func (c Car) ChangeColor(color string) {
// 	c.Color = color // cópia do color original
// 	println("New color: " + c.Color)
// }

func (c *Car) ChangeColor(color string) {
	c.Color = color // agora muda a color original
	println("New color: " + c.Color)
}

func main() {
	// car := Car{Model: "Ferrari", Color: "Red"}
	// println(car.Model)
	// car.Start()

	// car.ChangeColor("Blue")

	// car.Model = "Fiat"
	// car.Start()
	// println(car.Color)

	// a := 10
	// // b := a // b copiou o valor de a, mas guardou em um espaço próprio de memória
	// b := &a // agora b aponta para o endereço da memória de a
	// // b = 20
	// *b = 20 // assim o a também passa valer 20

	// println(a)
	// println(&a) // endereço onde o valor de a está armazenado
	// println(b)
	// println(a)

	// order, err := entity.NewOrder("1", 50.0, 5.5)

	// if err != nil {
	// 	println(err.Error())
	// } else {
	// 	println(order.Id)
	// }

	db, err := sql.Open("sqlite3", "db.sqlite3")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	orderRepository := database.NewOrderRepositorySQL(db)

	usecaseCFP := usecase.NewCalculateFinalPrice(orderRepository)

	// input := usecase.OrderInput{
	// 	Id:    "12345",
	// 	Price: 10.0,
	// 	Tax:   1.0,
	// }

	// output, err := usecaseCFP.Execute(input)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(output)

	ch, err := rabbitmq.OpenChannel()

	if err != nil {
		panic(err)
	}

	defer ch.Close()

	msgRabbitmqChannel := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, msgRabbitmqChannel)
	rabbitmqWorker(msgRabbitmqChannel, usecaseCFP)
}

func rabbitmqWorker(msgChan chan amqp.Delivery, uc *usecase.CalculateFinalPrice) {
	fmt.Println("Starting rabbitmq")

	for msg := range msgChan {
		var input usecase.OrderInput
		err := json.Unmarshal(msg.Body, &input)

		if err != nil {
			panic(err)
		}

		output, err := uc.Execute(input)

		if err != nil {
			panic(err)
		}

		msg.Ack(false)

		fmt.Println("Mensagem processada e salva no banco: ", output)
	}
}
