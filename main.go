package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/config"
)

type TradeClientApplication interface {
	OnCreate(sessionID quickfix.SessionID)
	OnLogon(sessionID quickfix.SessionID)
	OnLogout(sessionID quickfix.SessionID)
	FromAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) (reject quickfix.MessageRejectError)
	ToAdmin(msg *quickfix.Message, sessionID quickfix.SessionID)
	ToApp(msg *quickfix.Message, sessionID quickfix.SessionID) (err error)
	FromApp(msg *quickfix.Message, sessionID quickfix.SessionID) (reject quickfix.MessageRejectError)
}

type MyTradeClient struct{}

func (t *MyTradeClient) OnCreate(sessionID quickfix.SessionID) {
	fmt.Printf("Session created: %s\n", sessionID.String())
}

func (t *MyTradeClient) OnLogon(sessionID quickfix.SessionID) {
	fmt.Printf("Logged on: %s\n", sessionID.String())
}

func (t *MyTradeClient) OnLogout(sessionID quickfix.SessionID) {
	fmt.Printf("Logged out: %s\n", sessionID.String())
}

func (t *MyTradeClient) FromAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) (reject quickfix.MessageRejectError) {
	fmt.Printf("Received admin message from %s: %v\n", sessionID.String(), msg.String())
	return nil // Replace with rejection logic if needed
}

func (t *MyTradeClient) ToAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) {
	fmt.Printf("Sending admin message to %s: %v\n", sessionID.String(), msg.String())
}

func (t *MyTradeClient) ToApp(msg *quickfix.Message, sessionID quickfix.SessionID) (err error) {
	fmt.Printf("Sending application message to %s: %v\n", sessionID.String(), msg.String())
	return nil // Replace with error handling logic
}

func (t *MyTradeClient) FromApp(msg *quickfix.Message, sessionID quickfix.SessionID) (reject quickfix.MessageRejectError) {
	fmt.Printf("Received application message from %s: %v\n", sessionID.String(), msg.String())
	return nil // Replace with rejection logic if needed
}

func main() {
	// Parse command-line flags
	var cnfgFile string
	flag.StringVar(&cnfgFile, "config", "", "Path to configuration file")
	flag.Parse()

	if cnfgFile == "" {
		log.Fatal("No configuration file provided. Use -config option to specify one.")
	}

	// Read and parse the configuration file
	cfg, err := config.BeginString.
	if err != nil {
		log.Fatalf("Error parsing configuration file: %v", err)
	}

	// Initialize QuickFIX with the parsed configuration
	settings := quickfix.NewSettings()
	err = cfg.Configure(settings)
	if err != nil {
		log.Fatalf("Error configuring QuickFIX settings: %v", err)
	}

	storeFactory := quickfix.NewMemoryStoreFactory()
	logFactory, err := quickfix.NewFileLogFactory(settings)
	if err != nil {
		log.Fatalf("Error creating log factory: %v", err)
	}

	app := &MyTradeClient{}

	initiator, err := quickfix.NewInitiator(app, storeFactory, settings, logFactory)
	if err != nil {
		log.Fatalf("Error creating initiator: %v", err)
	}

	err = initiator.Start()
	if err != nil {
		log.Fatalf("Error starting initiator: %v", err)
	}

	// Wait for a condition to stop the initiator
	// This example uses a simple block; replace with your logic
	select {}
}
