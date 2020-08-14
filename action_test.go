package goami2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var action *Action

func loginAction(t *testing.T) {
	login := action.Login("admin", "pa55w0rd")
	assert.Contains(t, string(login), "Action: Login\r\n")
	assert.Contains(t, string(login), "Username: admin\r\n")
	assert.Contains(t, string(login), "Secret: pa55w0rd\r\n")
}

func actionSingleField(t *testing.T) {
	action.New("CoreStatus")
	pack := action.Message()
	expected := "Action: CoreStatus\r\nActionID: " +
		action.ActionID() + "\r\n\r\n"
	assert.Equal(t, string(pack), expected)
}

func multiHeaders(t *testing.T) {
	action.New("ExtensionState")
	action.Field("Exten", "5522")
	action.Field("Context", "default")
	pack := action.Message()

	assert.Contains(t, string(pack), "Action: ExtensionState\r\n")
	assert.Contains(t, string(pack), "Exten: 5522\r\n")
	assert.Contains(t, string(pack), "Context: default\r\n")
	assert.Contains(t, string(pack), "ActionID: "+action.ActionID())
}

func createFromJson(t *testing.T) {
	jsonStr := `{"action":"QueueStatus","queue":"Books_12",` +
		`"actionid":"674a1c86ab9bf@okon.ferry.clusterpbx.xyz"}`
	a, err := ActionFromJSON(jsonStr)
	assert.Nil(t, err)

	assert.Contains(t, string(a.Message()), "Action: QueueStatus\r\n")
	assert.Contains(t, string(a.Message()), "Queue: Books_12\r\n")
	assert.Contains(t, string(a.Message()),
		"Actionid: 674a1c86ab9bf@okon.ferry.clusterpbx.xyz\r\n")
	assert.Equal(t, "674a1c86ab9bf@okon.ferry.clusterpbx.xyz",
		a.ActionID())
}

func createFromJsonWithActionId(t *testing.T) {
	jsonStr := `{"action":"ConfbridgeKick","conference":"Sales",` +
		`"channel":"Local/Sales-65f4a00b-001"}`
	a, err := ActionFromJSON(jsonStr)
	assert.Nil(t, err)
	assert.Contains(t, string(a.Message()), "Action: ConfbridgeKick\r\n")
	assert.Contains(t, string(a.Message()), "Conference: Sales\r\n")
	assert.Contains(t, string(a.Message()),
		"Channel: Local/Sales-65f4a00b-001\r\n")
	assert.Contains(t, string(a.Message()),
		"ActionID: "+a.ActionID()+"\r\n")
}

func createFromJsonWithVars(t *testing.T) {
	jsonStr := `{"action":"Originate","channel":"Local/5555@confbridge-invite",` +
		`"exten":"MEETME_5555","priority":10,"context":"confbridge-join-member",` +
		`"callerid":"\"Alice\" <7777>","variable":` +
		`{"realm":"sip.conference.com","pin":"1234","confid":4567}}`
	a, err := ActionFromJSON(jsonStr)
	assert.Nil(t, err)
	assert.Contains(t, string(a.Message()), "Action: Originate\r\n")
	assert.Contains(t, string(a.Message()), "Variable: pin=1234\r\n")
	assert.Contains(t, string(a.Message()), "Variable: confid=4567\r\n")
}

func createFromJsonWithChanVars(t *testing.T) {
	jsonStr := `{"action":"Originate","channel":"Local/5555@confbridge-invite",` +
		`"exten":"MEETME_5555","priority":10,"context":"confbridge-join-member",` +
		`"variable":{"pin":"","confid":5990},` +
		`"chanvariable":{"external":true,"prefix":"000"}}`
	a, err := ActionFromJSON(jsonStr)
	assert.Nil(t, err)
	assert.Contains(t, string(a.Message()), "Action: Originate\r\n")
	assert.Contains(t, string(a.Message()), "Priority: 10\r\n")
	assert.Contains(t, string(a.Message()), "Variable: pin=\r\n")
	assert.Contains(t, string(a.Message()), "Variable: confid=5990\r\n")
	assert.Contains(t, string(a.Message()), "Chanvariable: external=true\r\n")
	assert.Contains(t, string(a.Message()), "Chanvariable: prefix=000\r\n")
}

func TestAction(t *testing.T) {
	action = NewAction()

	t.Run("Action Login", loginAction)
	t.Run("Action with single field", actionSingleField)
	t.Run("Action with multiple fields", multiHeaders)
	t.Run("Action generate from JSON", createFromJson)
	t.Run("Action generate from JSON and add ActionID",
		createFromJsonWithActionId)
	t.Run("Action generate from JSON with variables",
		createFromJsonWithVars)
	t.Run("Action generate from JSON with ChanVariable",
		createFromJsonWithChanVars)
}
