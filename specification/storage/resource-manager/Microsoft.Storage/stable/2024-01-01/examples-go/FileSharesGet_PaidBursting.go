package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/86c6306649b02e542117adb46c61e8019dbd78e9/specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/FileSharesGet_PaidBursting.json
func ExampleFileSharesClient_Get_getSharePaidBursting() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFileSharesClient().Get(ctx, "res9871", "sto6217", "share1634", &armstorage.FileSharesClientGetOptions{Expand: nil,
		XMSSnapshot: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FileShare = armstorage.FileShare{
	// 	Name: to.Ptr("share1634"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/fileServices/shares"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9871/providers/Microsoft.Storage/storageAccounts/sto6217/fileServices/default/shares/share1634"),
	// 	Etag: to.Ptr("\"0x8D592D74CC20EBA\""),
	// 	FileShareProperties: &armstorage.FileShareProperties{
	// 		FileSharePaidBursting: &armstorage.FileSharePropertiesFileSharePaidBursting{
	// 			PaidBurstingEnabled: to.Ptr(true),
	// 			PaidBurstingMaxBandwidthMibps: to.Ptr[int32](10340),
	// 			PaidBurstingMaxIops: to.Ptr[int32](102400),
	// 		},
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-26T05:06:14.000Z"); return t}()),
	// 		ShareQuota: to.Ptr[int32](1024),
	// 	},
	// }
}
