package common

import (
	"github.com/streadway/amqp"
	"testing"
	"time"
)

func Test_rabbitC(t *testing.T) {
	r := NewRabbitC("amqp://gateway-ws:gateway-ws123@127.0.0.1:5672/gateway-ws", afterConnected, nil)

	if err := r.Start(); err != nil {
		t.Error(err)
		return
	}

	r.SetNewMessageCallBack(receiveMessage)

	for i := 0; i < 1000; i++ {
		// r.PushTransientMessage("gateway-ws", "push_test", []byte("123"))
		r.PushPersistentMessage("gateway-ws", "push_test", []byte("123"))
		time.Sleep(time.Millisecond * 10)
	}

	time.Sleep(time.Minute * 2)

	r.Close()
}

func receiveMessage(d amqp.Delivery) {
	println(string(d.Body))
}

func afterConnected(r *RabbitC, ch *amqp.Channel) error {
	if err := ch.ExchangeDeclare("gateway-ws", amqp.ExchangeDirect, true, false, false, false, nil); err != nil {
		return err
	}

	queuesName := []string{}

	if queue, err := ch.QueueDeclare("test1", true, false, false, false, nil); err != nil {
		return err
	} else {
		if err = ch.QueueBind(queue.Name, "push_test", "gateway-ws", false, nil); err != nil {
			return err
		}
		queuesName = append(queuesName, queue.Name)
	}

	r.SetConsumeQueues(queuesName)

	return nil
}
