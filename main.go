// Simulateur de capteurs météo
// données https://www.kaggle.com/datasets/vanvalkenberg/historicalweatherdataforindiancities

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const Periode = 10
const Decalage = 1

// Callback MQTT
var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

// simulateur de capteurs
func capteur(mainTopic string, client mqtt.Client) {
	defer wg.Done()
	for {
		for _, d := range readWeatherList("data/" + mainTopic + ".csv") {
			d.Time = time.Now().Format("2006-01-02T15:04:05Z07:00")
			jsonData, err := json.MarshalIndent(d, "", "  ")
			if err != nil {
				log.Fatal(err)
			}
			token := client.Publish("topic/"+mainTopic, 0, false, jsonData)
			token.Wait()

			time.Sleep(Periode * time.Second)
		}
	}
}

var wg sync.WaitGroup

func main() {
	// configuration
	cfgPath, err := ParseFlags()
	config, err := NewConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	// Initialisation MQTT
	var broker = config.Mqtt.Broker
	var port = config.Mqtt.Port
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", broker, port))
	opts.SetClientID("sim-mesures")
	opts.SetUsername(config.Mqtt.Username)
	opts.SetPassword(config.Mqtt.Password)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Création des capteur virtuel
	for _, c := range []string{"Bangalore", "Bhubhneshwar", "Delhi", "Lucknow", "Mumbai", "Rajasthan", "Rourkela"} {
		token := client.Subscribe("topic/"+c, 1, nil)
		token.Wait()
		fmt.Printf("Subscribed to topic: %s", "topic/"+c)

		wg.Add(1)
		go capteur(c, client)
		time.Sleep(Decalage * time.Second)
	}

	// attente de la fin du monde
	wg.Wait()
}
