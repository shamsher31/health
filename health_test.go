package health

import "testing"

func TestNewHealth(t *testing.T) {
	h := NewHealth()

	if !h.IsDown() {
		t.Errorf("NewHealth().IsDown() == %t, want %t", h.IsDown(), true)
	}
}

func Test_Health_Unknown(t *testing.T) {
	h := NewHealth()
	h.Unknown()

	if !h.IsUnknown() {
		t.Errorf("NewHealth().IsUnknown() == %t, want %t", h.IsUnknown(), true)
	}
}

func Test_Health_Up(t *testing.T) {
	h := NewHealth()
	h.Up()

	if !h.IsUp() {
		t.Errorf("NewHealth().IsUp() == %t, want %t", h.IsUp(), true)
	}
}

func Test_Health_Down(t *testing.T) {
	h := NewHealth()
	h.Up()
	h.Down()

	if !h.IsDown() {
		t.Errorf("NewHealth().IsDown() == %t, want %t", h.IsDown(), true)
	}
}

func Test_Health_OutOfService(t *testing.T) {
	h := NewHealth()
	h.OutOfService()

	if !h.IsOutOfService() {
		t.Errorf("NewHealth().IsOutOfService() == %t, want %t", h.IsOutOfService(), true)
	}
}

func Test_Health_IsUp(t *testing.T) {
	h := NewHealth()
	h.Up()

	if h.status != up {
		t.Errorf("NewHealth().status == %s, want %s", h.status, up)
	}
}

func Test_Health_IsDown(t *testing.T) {
	h := NewHealth()

	if h.status != down {
		t.Errorf("NewHealth().status == %s, want %s", h.status, down)
	}
}

func Test_Health_IsOutOfService(t *testing.T) {
	h := NewHealth()
	h.OutOfService()

	if h.status != outOfService {
		t.Errorf("NewHealth().status == %s, want %s", h.status, outOfService)
	}
}

func Test_Health_MarshalJSON(t *testing.T) {
	h := NewHealth()
	h.Up()

	h.Info = map[string]interface{}{
		"status": "Should not render",
	}

	json, err := h.MarshalJSON()

	if err != nil {
		t.Errorf("err != nil, wants nil")
	}

	expected := `{"status":"UP"}`

	if string(json) != expected {
		t.Errorf("h.MarshalJSON() == %s, wants %s", string(json), expected)
	}
}
