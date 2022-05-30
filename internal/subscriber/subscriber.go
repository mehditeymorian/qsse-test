package subscriber

import (
	"sync"

	"github.com/snapp-incubator/qsse"
	"go.uber.org/zap"
)

type Subscriber struct {
	WaitGroup     *sync.WaitGroup
	Logger        *zap.Logger
	ServerAddress string
}

func (s Subscriber) Subscribe(cfg Config) {
	for i := 1; i <= cfg.Count; i++ {
		go s.subscribeClient(i, cfg.Topics)
	}

	s.WaitGroup.Wait()
}

func (s Subscriber) subscribeClient(clientID int, topics []string) {
	s.WaitGroup.Add(1)

	client, err := qsse.NewClient(s.ServerAddress, topics, nil)
	if err != nil {
		s.Logger.Warn("failed to create client", zap.Error(err))
		s.WaitGroup.Done()

		return
	}

	client.SetMessageHandler(func(topic string, _ []byte) {
		s.Logger.Info("new event received", zap.String("topic", topic), zap.Int("client_id", clientID))
	})

	select {}
}
