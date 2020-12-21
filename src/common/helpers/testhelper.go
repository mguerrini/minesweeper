package helpers

import "testing"

func AssertTrue (t *testing.T, value bool, msg string){
	if value {
		return
	}

	t.Error("Fail - " + msg)
	t.FailNow()
}

func AssertFalse (t *testing.T, value bool, msg string){
	if !value {
		return
	}

	t.Error("Fail - " + msg)
	t.FailNow()
}

func AssertError (t *testing.T, err error){
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func AssertErrorWithMsg (t *testing.T, msg string, err error){
	if err != nil {
		t.Error(msg, err)
		t.FailNow()
	}
}



