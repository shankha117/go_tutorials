package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"
	"tradier/config"
	"tradier/kafka"
	"tradier/kafka_service"

	"github.com/fatih/structs"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

// type Message struct {
// 	Symbol string `json:"symbol"`

// 	Size string `json:"size"`

// 	Date string `json:"date"`
// 	Last string `json:"last"`
// }

var logger = log.With().Str("pkg", "main").Logger()
var configs config.Configurations = config.Getconfig()

type commandStruct struct {
	Symbols   []string `json:"symbols"`
	SessionId string   `json:"sessionid"`
	Filter    []string `json:"filter"`
}

// func main() {
// 	m := make(map[string]string)
// 	m["hello"] = "data"
// 	m["world"] = "data"

// 	// Initial State
// 	for key, val := range m {
// 		fmt.Println(key, "=>", val)
// 	}

// 	fmt.Println("-----------------------")
// 	fmt.Println("Welcome to GO!")

// 	kafka_producer.Produce_to_kafka(m)
// }

func main() {

	sessionval := marketSession()
	// connect to kafka
	kafkaProducer, err := kafka.Configure(strings.Split(configs.Kafka.KafkaBrokerUrl, ","), configs.Kafka.KafkaClientId, configs.Kafka.KafkaTopic)
	if err != nil {
		logger.Error().Str("error", err.Error()).Msg("unable to configure kafka")
		panic("Kafka Connection Could Not be established !!")
	}
	defer kafkaProducer.Close()

	flag.Parse()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: configs.Tradier.Websocket_uri, Path: configs.Tradier.Market_steam_events}

	log.Printf("Connecting to %s ...", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		logger.Error().Str("error", err.Error()).Msg("could not dial to websocket")
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				logger.Error().Str("error", err.Error()).Msg("could not read message from websocket")
				return
			}
			var dat Message
			if err := json.Unmarshal(message, &dat); err != nil {
				panic(err)
			}
			log.Printf("Data received from socket --------")
			map_data := structs.Map(dat)
			bytes_data_map, err := json.Marshal(map_data)
			if err != nil {
				panic(err)
			}

			// m := map[string]string{}
			// if err := json.Unmarshal(message, &m); err != nil {
			// 	panic(err)
			// }

			// data_map := map[string]string{"sym": m["symbol"], "shares": m["size"], "ts": m["date"], "price": m["last"]}
			// bytes_data_map, err := json.Marshal(data_map)
			// if err != nil {
			// 	panic(err)
			// }
			go kafka_service.Produce_to_kafka(bytes_data_map)
		}
	}()
	command := &commandStruct{
		Symbols:   []string{"SPY"},
		SessionId: sessionval,
		Filter:    []string{"trade"},
	}

	ticker := time.NewTicker(time.Second * 1000)
	defer ticker.Stop()
	connectionErr := c.WriteJSON(command)
	if connectionErr != nil {
		logger.Error().Str("error", connectionErr.Error())
	}

	for {
		select {
		case <-done:
			fmt.Println(1122)
			return
		case <-interrupt:
			logger.Error().Str("error", "interrupt")

			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				logger.Error().Str("error", err.Error()).Msg("write close")
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

func marketSession() string {

	apiUrl := configs.Tradier.Base_uri + "/" + configs.Tradier.Market_steam_events + "/session"
	data := url.Values{}

	u, _ := url.ParseRequestURI(apiUrl)
	urlStr := u.String()

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	r.Header.Add("Authorization", "Bearer "+configs.Tradier.Token)
	r.Header.Add("Accept", "application/json")
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	log.Printf("\nGetting the Session Token ...")

	resp, _ := client.Do(r)
	responseData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		logger.Error().Str("error", err.Error()).Msg("could not get the session ID for Tradier")
	}
	streammap := map[string]map[string]string{}
	json.Unmarshal(responseData, &streammap)
	log.Printf("Session Token Received ..")
	return streammap["stream"]["sessionid"]
}
