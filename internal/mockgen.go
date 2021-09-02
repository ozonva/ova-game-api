package internal

//go:generate mockgen -destination=./mocks/flusher_mock.go -package=mocks github.com/ozonva/ova-game-api/internal/flusher Flusher
//go:generate mockgen -destination=./mocks/metrics_mock.go -package=mocks github.com/ozonva/ova-game-api/internal/metrics Metrics
//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ozonva/ova-game-api/internal/repo HeroRepo
//go:generate mockgen -destination=./mocks/kafka_producer_mock.go -package=mocks github.com/ozonva/ova-game-api/internal/kafka KafkaProducer
