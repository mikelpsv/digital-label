package model

import (
	"context"
	"encoding/json"
	kafka "github.com/segmentio/kafka-go"
	"log"
	"strconv"
)

type IncomingMessage struct {
	NumberTE string
	DataType int
	Payload  string
}

type ReaderConfig struct {
	Brokers []string
	GroupId string
	Topic   string
}

type DataReader struct {
	Config ReaderConfig
}

func (dr *DataReader) getKafkaReader() *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  dr.Config.Brokers,
		GroupID:  dr.Config.GroupId,
		Topic:    dr.Config.Topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}

func (dr *DataReader) Reading() {
	reader := dr.getKafkaReader()
	defer reader.Close()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		code, err := ConvertMessage(msg)
		err = code.Write()
	}
}

func ConvertMessage(msg kafka.Message) (*Code, error) {
	code := new(Code)
	inMsg := IncomingMessage{}
	err := json.Unmarshal(msg.Value, &inMsg)
	if err != nil {
		return nil, err
	}
	intVal, err := strconv.Atoi(inMsg.NumberTE)
	if err != nil {
		return nil, err
	}
	code.KeyLink = NewEnc62("").Encode(uint64(intVal))
	code.KeyData = inMsg.NumberTE
	code.Payload = inMsg.Payload
	return code, nil
}
