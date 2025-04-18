package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/97ee23a6db6078abcbec7b75bf9af8c503e9bb8b/specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/FileSharesGet_ProvisionedV2.json
func ExampleFileSharesClient_Get_getShareProvisionedV2() {
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
	// 		IncludedBurstIops: to.Ptr[int32](15000),
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-25T01:50:50.000Z"); return t}()),
	// 		MaxBurstCreditsForIops: to.Ptr[int64](36000000),
	// 		NextAllowedProvisionedBandwidthDowngradeTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2024-10-25T01:48:09.000Z"); return t}()),
	// 		NextAllowedProvisionedIopsDowngradeTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2024-10-25T01:48:09.000Z"); return t}()),
	// 		NextAllowedQuotaDowngradeTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2024-10-26T01:50:50.000Z"); return t}()),
	// 		ProvisionedBandwidthMibps: to.Ptr[int32](200),
	// 		ProvisionedIops: to.Ptr[int32](5000),
	// 		ShareQuota: to.Ptr[int32](100),
	// 	},
	// }
}
