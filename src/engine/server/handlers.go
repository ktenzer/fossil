package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"fossil/src/engine/util"
	"net/http"
	"strings"
)

// GetStatus godoc
// @Description Status and version information for the service
// @Accept  json
// @Produce  json
// @Success 200 {string} string 
// @Header 200 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /status [get]
func GetStatus(w http.ResponseWriter, r *http.Request) {
	var status util.Status
	status.Msg = "OK"
	status.Version = version
	
	json.NewEncoder(w).Encode(status)
}

// SendTrapSuccessCmd godoc
// @Description Execute command after successfull workflow execution
// @Param config body util.Config true "config struct"
// @Accept  json
// @Produce  json
// @Success 200 {object} util.Result 
// @Header 200 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /sendTrapSuccessCmd [post]
func SendTrapSuccessCmd(w http.ResponseWriter, r *http.Request) {
	var result util.Result
	config,_ := util.GetConfig(w,r)
	printConfigDebug(config)

	if config.SendTrapSuccessCmd != "" {
		args := strings.Split(config.SendTrapSuccessCmd, ",")
		message := util.SetMessage("INFO", "Performing send trap success command [" + config.SendTrapSuccessCmd + "]")

		result = util.ExecuteCommand(args...)
		result.Messages = util.PrependMessage(message,result.Messages)

		_ = json.NewDecoder(r.Body).Decode(&result)
		json.NewEncoder(w).Encode(result)
	}
}

// SendTrapErrorCmd godoc
// @Description Execute command after failed workflow execution
// @Param config body util.Config true "config struct"
// @Accept  json
// @Produce  json
// @Success 200 {object} util.Result 
// @Header 200 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /sendTrapErrorCmd [post]
func SendTrapErrorCmd(w http.ResponseWriter, r *http.Request) {
	var result util.Result
	config,_ := util.GetConfig(w,r)
	printConfigDebug(config)

	if config.SendTrapSuccessCmd != "" {
		args := strings.Split(config.SendTrapErrorCmd, ",")
		message := util.SetMessage("INFO", "Performing send trap error command [" + config.SendTrapSuccessCmd + "]")

		result = util.ExecuteCommand(args...)
		result.Messages = util.PrependMessage(message,result.Messages)

		_ = json.NewDecoder(r.Body).Decode(&result)
		json.NewEncoder(w).Encode(result)
	}
}

// GetJobs godoc
// @Description Get jobs (workflows) that have executed for a profile/config
// @Param profileName path string true "name of profile"
// @Param configName path string true "name of config"
// @Accept  json
// @Produce  json
// @Success 200 {object} util.Jobs
// @Header 200 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /getJobs/{profileName}/{configName} [get]
func GetJobs(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)	
	var profileName string = params["profileName"]
	var configName string = params["configName"]

	jobsDir := dataDir + "/" + profileName + "/" + configName
	
	var result util.Result
	var messages []util.Message

	var jobs util.Jobs
	var jobList []util.Job
	jobList,err := util.ListJobs(jobsDir)

	if err != nil {
		msg := util.SetMessage("ERROR","Job list failed! " + err.Error())
		messages = append(messages, msg)

		result = util.SetResult(1, messages)
		jobs.Result = result

		_ = json.NewDecoder(r.Body).Decode(&jobs)
		json.NewEncoder(w).Encode(jobs)
	}

	jobs.Result.Code = 0
	jobs.Jobs = jobList

	_ = json.NewDecoder(r.Body).Decode(&jobs)
	json.NewEncoder(w).Encode(jobs)
}