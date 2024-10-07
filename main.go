package main

import (
	"fmt"

	"github.com/quickfixgo/quickfix"
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
	cnfgFile := "./config.cfg"
	settings, err := quickfix.NewSessionSettingsFromFile(cnfgFile)
	if err != nil {
		fmt.Printf("Error loading session settings from file: %s\n", err)
		return
	}

	application := new(MyTradeClient)
	storeFactory := quickfix.NewMemoryStoreFactory()
	logFactory := quickfix.NewScreenLogFactory(settings)
	initiator, err := quickfix.NewInitiator(application, storeFactory, settings, logFactory)
	if err != nil {
		fmt.Printf("Unable to create initiator: %s\n", err)
		return
	}

	err = initiator.Start()
	if err != nil {
		fmt.Printf("Error starting initiator: %s\n", err)
		return
	}

	// Wait for a condition to stop the initiator (example: a signal or event)
	// This can be implemented based on your application's requirements
	select {}
}
