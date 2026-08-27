package models

import (
	"encoding/json"
	"testing"
)

func TestWebhookMarshal_sslVerificationFalseIsExplicit(t *testing.T) {
	name := "kea-sync"
	url := "https://example.com/webhook"
	secret := "hmac-secret"
	ssl := false
	tagName := "internal-infra"
	tagSlug := "internal-infra"

	w := &Webhook{
		Name:            &name,
		PayloadURL:      &url,
		Secret:          &secret,
		SslVerification: &ssl,
		Tags: []*NestedTag{
			{Name: &tagName, Slug: &tagSlug},
		},
	}

	b, err := json.Marshal(w)
	if err != nil {
		t.Fatal(err)
	}

	var got map[string]any
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}

	v, ok := got["ssl_verification"]
	if !ok {
		t.Fatalf("ssl_verification omitted from JSON: %s", b)
	}
	if v != false {
		t.Fatalf("ssl_verification = %#v, want false; json=%s", v, b)
	}

	if got["secret"] != secret {
		t.Fatalf("secret = %#v, want %q; json=%s", got["secret"], secret, b)
	}

	tags, ok := got["tags"].([]any)
	if !ok || len(tags) != 1 {
		t.Fatalf("tags = %#v, want 1 entry; json=%s", got["tags"], b)
	}
}

func TestWebhookMarshal_emptySecretSerializes(t *testing.T) {
	name := "kea-sync"
	url := "https://example.com/webhook"
	empty := ""
	w := &Webhook{
		Name:       &name,
		PayloadURL: &url,
		Secret:     &empty,
	}

	b, err := json.Marshal(w)
	if err != nil {
		t.Fatal(err)
	}

	var got map[string]any
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}

	if _, ok := got["secret"]; !ok {
		t.Fatalf("empty secret omitted from JSON (cannot clear): %s", b)
	}
	if got["secret"] != "" {
		t.Fatalf("secret = %#v, want empty string; json=%s", got["secret"], b)
	}
}
