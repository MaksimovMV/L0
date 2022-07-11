package main

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"test/internal/model"
	"time"
)

func main() {
	sc, err := stan.Connect("test-cluster", "clientID2")
	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.Open("model.json")
	if err != nil {
		log.Fatalln(err)
	}

	b, err := ioutil.ReadAll(f)

	order := model.Order{}

	if err = json.Unmarshal(b, &order); err != nil {
		log.Fatalln(err)
	}

	id := order.OrderUid

	for i := 0; ; i++ {
		order.OrderUid = id + strconv.Itoa(i)
		message, err := json.Marshal(order)
		if err != nil {
			log.Fatalln(err)
		}
		sc.Publish("subject1", message)
		log.Printf("%v has been sent\n", order.OrderUid)
		time.Sleep(2 * time.Second)
	}

}
