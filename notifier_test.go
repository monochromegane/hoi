package hoi

import "testing"

func TestNewNotifier(t *testing.T) {
	conf := Notification{To: "slack"}
	n := NewNotifier(conf)
	if s, ok := n.(*SlackNotifier); !ok {
		t.Errorf("Type should be *SlackNotifier, but %v", s)
	}

	conf = Notification{To: "takosan"}
	n = NewNotifier(conf)
	if s, ok := n.(*TakosanNotifier); !ok {
		t.Errorf("Type should be *TakosanNotifier, but %v", s)
	}
}
