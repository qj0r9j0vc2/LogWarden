package file

import (
	"LogWarden/config"
	"LogWarden/email"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	file        *LogFile
	LogFilePath = "./resources/apigateway.log"
	TargetLevel = ERROR_LEVEL
)

const (
	ERROR_LEVEL   = "ERROR"
	WARN_LEVEL    = "WARN"
	INFO_LEVEL    = "INFO"
	DEBUG_LEVEL   = "DEBUG"
	TMP_FILE_PATH = "./resources/logwarden.txt"
)

type LogFile struct {
	Reader *bufio.Reader
}

func (logFile *LogFile) check(matchStr string) []string {
	var resultList []string
	for {
		bytes, isPrefix, err := logFile.Reader.ReadLine()
		if isPrefix || err != nil {
			break
		}

		comparator := string(bytes[:40])
		if strings.Contains(comparator, matchStr) {
			resultList = append(resultList, string(bytes))
		}
	}
	return resultList
}

func DetectLog() {
	open, err := os.Open(LogFilePath)
	if err != nil {
		log.Fatalln(err.Error())
	}

	file = &LogFile{Reader: bufio.NewReader(open)}

	resultList := file.check(TargetLevel)
	if err != nil {
		log.Fatalln(err.Error())
	}
	if len(resultList) == 0 {
		log.Println("LogWarden didn't detected.")
	} else {
		var body string
		for _, result := range resultList {
			body += result
		}
		fmt.Println(body)
		//util.SendEmail(body)
		f1, err := os.Create(TMP_FILE_PATH)
		_, err = fmt.Fprintf(f1, body)
		if err != nil {
			return
		}

		sender := email.New()
		m := email.NewMessage("LogWarden", "LogWarden detected anomaly.")
		m.To = []string{config.AppConfig.Email.Receiver}
		err = m.AttachFile(TMP_FILE_PATH)
		if err != nil {
			log.Fatalln(err.Error())
		}
		err = sender.Send(m)
		if err != nil {
			log.Fatalln(err.Error())
		}
	}

}
