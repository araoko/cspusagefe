package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/araoko/cspusagefe/config"
	"github.com/icza/gowut/gwu"
	"github.com/kardianos/service"
)

// var fileLogger *log.Logger
// var winEvtlogger service.Logger

const configFileName = "config.json"

type cspUsageSvc struct {
	cspUsgSvr gwu.Server
	c         *config.Config
	l         *ServiceLogger
}

func (mySvc *cspUsageSvc) Start(s service.Service) error {
	err := mySvc.l.InitLogger()
	if err != nil {
		return err
	}
	mySvc.l.Info("Starting Service", true)
	go mySvc.run() // not neat?
	return nil

}

func (mySvc *cspUsageSvc) Stop(s service.Service) error {
	mySvc.l.Info("Stopping Service", true)
	//mySvc.cspUsgSvr. stop server
	mySvc.l.Info("Service Stopped", true)
	mySvc.l.FreeLogger()
	return nil
}

func (mySvc *cspUsageSvc) run() {
	var err error
	mySvc.cspUsgSvr = startHTTPServer(mySvc.c, mySvc.l)
	if err != nil {
		errMessage := fmt.Sprintf("Error starting server: %s", err.Error())
		mySvc.l.Error(errMessage, true)

		return
	}
	mySvc.l.Info("CSP Usage Front-End Server Started", true)

}

// func getServiceLogger() *log.Logger {
// 	return nil
// }

func getPath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(getProgarmDir(), path)
}

func getProgarmDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func main() {

	svcFlag := flag.String("service", "", "Control the system service.")
	flag.Parse()

	svcConfig := &service.Config{
		Name:        "CSPUsageFE",
		DisplayName: "Azure CSP Usage Web Server",
		Description: "Azure CSP Usage Front End Service",
		//Dependencies: []string{"MySQL80"},
	}

	c, err := config.LoadConfigFromFile(filepath.Join(getProgarmDir(), configFileName))
	if err != nil {
		panic(err)
	}

	prg := &cspUsageSvc{
		c: c,
	}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	prg.l = NewServiceLogger(s, service.Interactive(), c.LogFile)

	if len(*svcFlag) != 0 {
		err := service.Control(s, *svcFlag)
		if err != nil {
			log.Printf("Valid actions: %q\n", service.ControlAction)
			log.Fatal(err)
		}
		return
	}

	// TODO do stuff to test all exernal serviceas, DB, AD etc
	err = s.Run()
	if err != nil {
		prg.l.Error(err.Error(), true)
	}

}

type Severity int

const (
	INFO Severity = iota
	WARNING
	ERROR
)

func (s Severity) String() string {
	switch s {
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"

	}
	return "UNKNOWN"
}

type ServiceLogger struct {
	logger      *log.Logger
	winLogger   service.Logger
	interactive bool
	logFile     string
	service     service.Service
}

func (sl *ServiceLogger) log(t Severity, s string, eventLog bool) {

	m := t.String() + ": " + s + "\r\n"
	sl.logger.Print(m)

	if sl.interactive {
		return
	}
	if eventLog {
		sl.logWinEvent(t, s)
	}

}

func (sl *ServiceLogger) logWinEvent(t Severity, s string) {
	switch t {
	case ERROR:
		sl.winLogger.Error(s)
	case INFO:
		sl.winLogger.Info(s)
	case WARNING:
		sl.winLogger.Warning(s)

	}

}

func (sl *ServiceLogger) Error(s string, eventLog bool) {
	sl.log(ERROR, s, eventLog)
}

func (sl *ServiceLogger) Info(s string, eventLog bool) {
	sl.log(INFO, s, eventLog)
}

func (sl *ServiceLogger) Warning(s string, eventLog bool) {
	sl.log(WARNING, s, eventLog)
}

func getEventLogger(s service.Service) (service.Logger, error) {
	return s.Logger(nil)
}

func (sl *ServiceLogger) FreeLogger() {
	if sl.logger == nil {
		return
	}

	if file, isFile := sl.logger.Writer().(*os.File); isFile {
		file.Close()
	}

	sl.logger = nil
}

func (sl *ServiceLogger) InitLogger() error {

	if sl.interactive {
		sl.createLog(os.Stdout)
		return nil
	}
	f, err := os.OpenFile(getPath(sl.logFile), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	sl.createLog(f)
	sl.winLogger, err = getEventLogger(sl.service)
	return err

}

func (sl *ServiceLogger) createLog(w io.Writer) {
	sl.logger = log.New(w, "", log.Ldate|log.Ltime|log.LUTC)
}

func NewServiceLogger(s service.Service, isInteractive bool, logFile string) *ServiceLogger {
	sl := &ServiceLogger{
		interactive: isInteractive,
		service:     s,
		logFile:     logFile,
	}

	return sl
}
