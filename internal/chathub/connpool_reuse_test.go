package chathub

import "testing"

func TestConversationReuseRequiresFreshWebSocket(t *testing.T) {
	if shouldReusePooledConnection(Request{ConversationID: "conv", SessionID: "sess"}) {
		t.Fatal("continuing an existing conversation on a pre-warmed websocket can return an immediate empty completion")
	}
	if !shouldReusePooledConnection(Request{}) {
		t.Fatal("fresh first-turn requests should remain eligible for the connection pool")
	}
}
