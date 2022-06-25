package qsse

// Config configuration struct of QSSE.
type Config struct {
	Topics                    []string `koanf:"topics"`
	Port                      int      `koanf:"port"`
	ClientAcceptorCount       int      `koanf:"client_acceptor_count"`
	ClientAcceptorQueueSize   int      `koanf:"client_acceptor_queue_size"`
	EventDistributorCount     int      `koanf:"event_distributor_count"`
	EventDistributorQueueSize int      `koanf:"event_distributor_queue_size"`
	CleaningInterval          int      `koanf:"cleaning_interval"`
}
