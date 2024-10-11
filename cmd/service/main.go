package main

import (
	"github.com/mikelpsv/digital-label/internal/app"
	"github.com/mikelpsv/digital-label/pkg/config"
)

func main() {
	cfg := config.ReadEnv()
	app.Init(cfg)

	//go func() {
	//	reader := kafka.NewReader(kafka.ReaderConfig{
	//		Brokers:         []string{config.Cfg.KafkaHost0},
	//		GroupID:         config.Cfg.KafkaDataGroup,
	//		Topic:           config.Cfg.KafkaDataTopic,
	//		MinBytes:        10e3, // 10KB
	//		MaxBytes:        10e6, // 10MB
	//		ReadLagInterval: 500 * time.Millisecond,
	//	})
	//	defer reader.Close()
	//	for {
	//		msg, err := reader.ReadMessage(ctx)
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//		code, err := model.ConvertMessage(msg)
	//		err = code.Write()
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//	}
	//
	//}()

}
