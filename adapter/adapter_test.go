package adapter

import "testing"

func TestAdapterByEmbedded(t *testing.T) {
	var decoretor Decorator

	decoretor = NewEmbeddedDecorateBanner("委譲")

	if str := decoretor.Decorate(); str != "*委譲*" {
		t.Errorf("Expected decorated 委譲 to %s but %s", "*委譲*", str)
	}
}
