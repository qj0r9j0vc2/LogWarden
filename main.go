package main

import (
	"LogWarden/config"
	logFile "LogWarden/file"
	"LogWarden/util"
	"bufio"
	"fmt"
	"github.com/manifoldco/promptui"
	"log"
	"os"
	"strings"
)

const (
	CronSpec = "0 0 * * *"
)

func main() {
	config.ParseFromConfig()
	//enterValue()
	//setLevel()

	logFile.DetectLog()

	//c := cron.New()
	//err := c.AddFunc(CronSpec, logFile.DetectLog)
	//if err != nil {
	//	return
	//}
	//
	//c.Start()
	//defer c.Stop()
}

func enterValue() {
	util.Clear()

	str := readString("LogFile path(default: ./resources/apigateway.log)")
	if str != "" {
		logFile.LogFilePath = str
	}
}

func readString(description string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter %s : ", description)

	strName, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(strName)
}

func setLevel() {
	levelList := []string{"ERROR", "WARN", "INFO", "DEBUG"}
	prompt := promptui.Select{Label: "Set target Level(this level will be captured.)", Items: levelList}
	i, _, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}
	switch i {
	case 0:
		logFile.TargetLevel = logFile.ERROR_LEVEL
		break
	case 1:
		logFile.TargetLevel = logFile.WARN_LEVEL
		break
	case 2:
		logFile.TargetLevel = logFile.INFO_LEVEL
		break
	case 3:
		logFile.TargetLevel = logFile.DEBUG_LEVEL
		return
	}
}
