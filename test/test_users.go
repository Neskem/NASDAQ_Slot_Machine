package test

import (
	"NASDAQ_Slot_Machine/controller"
	unitTest "github.com/Valiben/gin_unit_test"
	"testing"
)

func TestLoginHandler(t *testing.T) {
	user := controller.Login{}
	resp := controller.LoginResponse{}
	err := unitTest.TestHandlerUnMarshalResp("POST", "/users/login/", "form", &user, &resp)
	if err != nil {
		t.Errorf("TestLoginHandler: %v\n", err)
		return
	}
	if resp.Status != 0 {
		t.Errorf("TestLoginHandler: response is not expected\n")
		return
	}

}
