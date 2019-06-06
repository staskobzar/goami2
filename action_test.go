package goami2

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
		action.ActionId() + "\r\n\r\n"
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
	assert.Contains(t, string(pack), "ActionID: "+action.ActionId())
}

func TestAction(t *testing.T) {
	action = NewAction()

	t.Run("Action Login", loginAction)
	t.Run("Action with single field", actionSingleField)
	t.Run("Action with multiple fields", multiHeaders)
}
