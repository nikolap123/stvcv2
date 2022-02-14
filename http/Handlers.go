package http

import "C"
import (
	"encoding/json"
	"net/http"
	"stvcv2/executor"
)

func HandleRunCommand(w http.ResponseWriter, r *http.Request) {

	setupResponse(&w, r)

	if r.Method == "OPTIONS" {
		return
	}

	var RCR RunCommandRequest
	var ExecService executor.ExecService

	err := json.NewDecoder(r.Body).Decode(&RCR)

	if err != nil {
		http.Error(w, "Request cannot be decoded!", http.StatusBadRequest)
		return
	}

	ExecService.Init(RCR.SequenceId, RCR.ApplicationId, RCR.DeviceId)

	ExecService.Run()

	if err != nil {
		json, _ := json.Marshal(Response{Message: err.Error(), Code: 500, Data: nil})

		w.WriteHeader(http.StatusBadRequest)
		w.Write(json)

		return
	}

	json, _ := json.Marshal(Response{Data: ExecService.GetResult(), Code: 200, Message: "Success!"})

	w.Write(json)
	return
}
