package main

import (
	"github.com/pborman/getopt/v2"
	"os"
	"engine/util"
	"fmt"
	
)

func main() {
	optProfile := getopt.StringLong("profile",'p',"","Profile name")
	optConfig := getopt.StringLong("config",'c',"","Config name")
	optConfigPath := getopt.StringLong("configPath",'o',"","Path to configs directory")
    optAction := getopt.StringLong("action",'a',"","backup|list|status")
    optHelp := getopt.BoolLong("help", 0, "Help")
	getopt.Parse()

    if *optHelp {
        getopt.Usage()
        os.Exit(0)
	}
	if getopt.IsSet("profile") != true {
		fmt.Println("ERROR: missing parameter --profile")
		getopt.Usage()
		os.Exit(1)	
	} else if getopt.IsSet("config") != true {
		fmt.Println("ERROR: missing parameter --config")
		getopt.Usage()
		os.Exit(1)
	} else if getopt.IsSet("action") != true {
		fmt.Println("ERROR: missing parameter --action")
		getopt.Usage()
		os.Exit(1)
	}

	var configPath string
	if getopt.IsSet("configPath") == true {
		configPath = *optConfigPath + "/" + *optProfile + "/" + *optConfig
	} else {
		configPath = "configs/" + *optProfile + "/" + *optConfig
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Println(err,"\n" + "ERROR: Profile of Config don't exist")
		os.Exit(1)
	}

	var config util.Config = util.ReadConfig(configPath)
	fmt.Println(config.BackupRetentions[1].Policy)

	if *optAction == "backup" {

		var result []util.Result
		result = util.StartBackupWorkflow(config)

		for index, item := range result {
			fmt.Println(index, item.Time, item.Code, item.Stdout, item.Stderr)
		}	


	} else if *optAction == "list" {

	} else if *optAction == "status" {
		fmt.Println("Checking status of services")

		var workflowStatus util.Status
		workflowStatus = util.GetWorkflowServiceStatus()
		fmt.Println("Workflow Service:", workflowStatus)

		var appStatus util.Status
		appStatus = util.GetAppServiceStatus()
		fmt.Println("App Service:", appStatus)

		var storageStatus util.Status
		storageStatus = util.GetStorageServiceStatus()
		fmt.Println("Storage Service:", storageStatus)

	}
}