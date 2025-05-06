package resulto

import (
	"errors"
	"testing"
)

func TestResulto_Success(t *testing.T) {
	r := Success("success")
	if r.IsErr() {
		t.Errorf("Expected success, got failure")
	}
}

func TestResulto_Failure(t *testing.T) {
	r := Failure[string](errors.New("failure"))
	if r.IsOk() {
		t.Errorf("Expected failure, got success")
	}
}
func TestResulto_Success_Empty(t *testing.T) {
	r := SuccessAny()
	if r.IsErr() {
		t.Errorf("Expected success, got failure")
	}
}

func TestResulto_Failure_Empty(t *testing.T) {
	r := FailureAny(errors.New("failure"))
	if r.IsOk() {
		t.Errorf("Expected failure, got success")
	}
}

func TestFailureOf(t *testing.T) {
	r := FailureOf(errors.New("failure"), "value")
	if r.IsOk() {
		t.Errorf("Expected failure, got success")
	}
}

func TestResulto_Unwrap(t *testing.T) {
	r := Success("success")
	if r.Unwrap() != "success" {
		t.Errorf("Expected success, got %v", r.Unwrap())
	}
}

func TestResulto_UnwrapOr(t *testing.T) {
	r := Failure[string](errors.New("failure"))
	if r.UnwrapOr("default") != "default" {
		t.Errorf("Expected default, got %v", r.UnwrapOr("default"))
	}
}

func TestResulto_UnwrapOr_Result(t *testing.T) {
	r := Success("success")
	if r.UnwrapOr("default") != "success" {
		t.Errorf("Expected success, got %v", r.UnwrapOr("default"))
	}
}

func TestResulto_UnwrapErr(t *testing.T) {
	r := Failure[string](errors.New("failure"))
	if r.UnwrapErr().Error() != "failure" {
		t.Errorf("Expected failure, got %v", r.UnwrapErr())
	}
}

func TestResulto_UnwrapErr_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, got nil")
		}
	}()
	r := Success("success")
	err := r.UnwrapErr()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

func TestResulto_Unwrap_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, got nil")
		}
	}()
	r := Failure[string](errors.New("failure"))
	r.Unwrap()
}

func TestResulto_IsOk(t *testing.T) {
	r := Success("success")
	if !r.IsOk() {
		t.Errorf("Expected success, got failure")
	}
}

func TestResulto_IsErr(t *testing.T) {
	r := Failure[string](errors.New("failure"))
	if !r.IsErr() {
		t.Errorf("Expected failure, got success")
	}
}

func TestResulto_IsErr_FailureOf(t *testing.T) {
	r := FailureOf(errors.New("failure"), "value")
	if !r.IsErr() {
		t.Errorf("Expected failure, got success")
	}
}

func TestResulto_IsOk_FailureOf(t *testing.T) {
	r := FailureOf(errors.New("failure"), "value")
	if r.IsOk() {
		t.Errorf("Expected success, got failure")
	}
}
