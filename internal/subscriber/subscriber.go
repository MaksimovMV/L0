package subscriber

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"log"
	"test/internal/model"
	"test/internal/store/cachestore"
)

func ConnectAndSubscribe(stanClusterID, clientID, subject, durable string, cStore *cachestore.Store) (stan.Conn, stan.Subscription, error) {
	sc, err := stan.Connect(stanClusterID, clientID)
	if err != nil {
		return nil, nil, err
	}

	sb, err := sc.Subscribe(subject, func(msg *stan.Msg) {
		log.Println("Message received")
		order := model.Order{}

		if err := json.Unmarshal(msg.Data, &order); err != nil {
			log.Println(err)
			return
		}

		if err := order.Validate(); err != nil {
			log.Println(err)
			return
		}

		if err := cStore.Create(&order); err != nil {
			log.Println(err)
			return
		}
		msg.Ack()
	}, stan.SetManualAckMode(), stan.DurableName(durable))
	if err != nil {
		return nil, nil, err
	}
	return sc, sb, nil
}
