package mux_server

import (
	base64 "encoding/base64"
	"encoding/json"
	"fmt"
	task "go-machinery/tasks"
	"go-machinery/utils"
	"log"
	"net/http"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/gorilla/mux"
)

func StartServer(taskserver *machinery.Server) {

	route := mux.NewRouter()
	s := route.PathPrefix("/api/v1").Subrouter() //Base Path

	s.HandleFunc("/send_email", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request")

		payload := new(task.Payload)
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			fmt.Print(err)
		}

		reqJSON, err := json.Marshal(payload)
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
		json.NewEncoder(w).Encode(res.GetState().TaskUUID)

	}).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", s)) // Run Server

}
