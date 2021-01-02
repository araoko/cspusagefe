package main

import (
	"fmt"

	"github.com/araoko/cspusagefe/config"
	"github.com/araoko/cspusagefe/rest"
	"github.com/araoko/cspusagefe/window"
	"github.com/icza/gowut/gwu"
)

const (
	defConfigPath = "cspusagefe.json"
)

func startHTTPServer(c *config.Config, logger *ServiceLogger) gwu.Server {

	certFilePath := getPath(c.CertFilePath)
	keyFilePath := getPath(c.KeyFilePath)
	r := rest.NewRestClient(c.BEServerName, c.BEServerPort, certFilePath,
		keyFilePath)
	server := gwu.NewServerTLS(c.AppRoot, fmt.Sprintf("%s:%d", c.ServerName,
		c.ServerPort), certFilePath, keyFilePath)
	//gwu.NewServer(c.AppRoot, fmt.Sprintf("%s:%d", c.ServerName, c.ServerPort))
	wm := window.NewWindowMaker(r, server)

	homeWin := wm.HomeWindow("home", "Home Window")

	server.SetText("Azure CSP Usage")
	server.AddSessCreatorName("login", "Login Window")
	server.AddSHandler(sessHandler{wm: wm})
	server.AddStaticDir("image", getPath("media"))
	server.AddWin(homeWin)
	fmt.Println("About to start gowut server")
	go func() {
		err := server.Start()
		if err != nil {
			errMessage := fmt.Sprintf("Gwu Server Error: %s", err.Error())
			logger.Error(errMessage, true)
		}
	}()
	fmt.Println("About to start gowut server serted")
	return server

}

func ce(err error) {
	if err != nil {
		panic(err)
	}
}

type sessHandler struct {
	wm window.WindowMaker
}

func (h sessHandler) Created(s gwu.Session) {
	//fmt.Println("building login window")
	window.BuildLoginWindow(s, h.wm)
}

func (h sessHandler) Removed(s gwu.Session) {

}
