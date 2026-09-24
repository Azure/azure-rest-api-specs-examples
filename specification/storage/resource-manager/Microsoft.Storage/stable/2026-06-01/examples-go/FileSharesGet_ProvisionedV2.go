package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v4"
)

// Generated from example definition: 2026-06-01/FileSharesGet_ProvisionedV2.json
func ExampleFileSharesClient_Get_getShareProvisionedV2() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFileSharesClient().Get(ctx, "res9871", "sto6217", "share1634", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstorage.FileSharesClientGetResponse{
	// 	FileShare: armstorage.FileShare{
	// 		Name: to.Ptr("share1634"),
	// 		Type: to.Ptr("Microsoft.Storage/storageAccounts/fileServices/shares"),
	// 		Etag: to.Ptr("\"0x8D592D74CC20EBA\""),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9871/providers/Microsoft.Storage/storageAccounts/sto6217/fileServices/default/shares/share1634"),
	// 		FileShareProperties: &armstorage.FileShareProperties{
	// 			IncludedBurstIops: to.Ptr[int32](15000),
	// 			LastModifiedTime: to.Ptr(time.Date(2024, time.October, 25, 1, 50, 50, 0, time.UTC)),
	// 			MaxBurstCreditsForIops: to.Ptr[int64](36000000),
	// 			NextAllowedProvisionedBandwidthDowngradeTime: to.Ptr(time.Date(2024, time.October, 25, 1, 48, 9, 0, time.UTC)),
	// 			NextAllowedProvisionedIopsDowngradeTime: to.Ptr(time.Date(2024, time.October, 25, 1, 48, 9, 0, time.UTC)),
	// 			NextAllowedQuotaDowngradeTime: to.Ptr(time.Date(2024, time.October, 26, 1, 50, 50, 0, time.UTC)),
	// 			ProvisionedBandwidthMibps: to.Ptr[int32](200),
	// 			ProvisionedIops: to.Ptr[int32](5000),
	// 			ShareQuota: to.Ptr[int32](100),
	// 		},
	// 	},
	// }
}
