package server

import (
	base64 "encoding/base64"
	"encoding/json"
	task "go-machinery/tasks"
	"go-machinery/utils"
	"log"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/gofiber/fiber"
)

func StartServer(taskserver *machinery.Server) {

	app := fiber.New()

	app.Post("/send_email", func(ctx *fiber.Ctx) {
		log.Println("Received request")
		p := new(task.Payload)
		if err := ctx.BodyParser(p); err != nil {
			utils.Logger.Error(err.Error())
		}

		reqJSON, err := json.Marshal(p)
		if err != nil {
			utils.Logger.Error(err.Error())
		}

		b64EncodedReq := base64.StdEncoding.EncodeToString([]byte(reqJSON))
		task := tasks.Signature{
			Name: "send_email",
			Args: []tasks.Arg{
				{
					Type:  "string",
					Value: b64EncodedReq,
				},
			},
		}

		res, err := taskserver.SendTask(&task)
		if err != nil {
			utils.Logger.Error(err.Error())	
		}

		ctx.JSON(&fiber.Map{
			"task_uuid": res.GetState().TaskUUID,
		})

	})

	app.Listen("localhost:5000")

}
