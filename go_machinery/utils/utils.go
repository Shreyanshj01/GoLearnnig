package utils

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"go.uber.org/zap"

	tasks "go-machinery/tasks"
)

var (
	Logger *zap.SugaredLogger
)

func init() {
	logger, _ := zap.NewProduction()
	Logger = logger.Sugar()
}
func GetMachineryServer() *machinery.Server {
	Logger.Info("initing task server")

	taskserver, err := machinery.NewServer(&config.Config{
		Broker:        "redis://default:athena@localhost:6381/6",
		ResultBackend: "redis://default:athena@localhost:6381/7",
	})
	// {REDIS_USER}:{REDIS_PASSWORD}@{REDIS_HOST}:{REDIS_PORT}/{REDIS_DATABASE_NUMBER}
	if err != nil {
		Logger.Fatal(err.Error())
	}

	taskserver.RegisterTasks(map[string]interface{}{
		"send_email": tasks.SendMail,
	})

	return taskserver
}
