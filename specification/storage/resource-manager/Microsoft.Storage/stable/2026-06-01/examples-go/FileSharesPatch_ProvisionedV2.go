package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v4"
)

// Generated from example definition: 2026-06-01/FileSharesPatch_ProvisionedV2.json
func ExampleFileSharesClient_Update_updateShareProvisionedV2() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFileSharesClient().Update(ctx, "res3376", "sto328", "share6185", armstorage.FileShare{
		FileShareProperties: &armstorage.FileShareProperties{
			ProvisionedBandwidthMibps: to.Ptr[int32](200),
			ProvisionedIops:           to.Ptr[int32](5000),
			ShareQuota:                to.Ptr[int32](100),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstorage.FileSharesClientUpdateResponse{
	// 	FileShare: armstorage.FileShare{
	// 		Name: to.Ptr("share6185"),
	// 		Type: to.Ptr("Microsoft.Storage/storageAccounts/fileServices/shares"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res3376/providers/Microsoft.Storage/storageAccounts/sto328/fileServices/default/shares/share6185"),
	// 		FileShareProperties: &armstorage.FileShareProperties{
	// 			IncludedBurstIops: to.Ptr[int32](15000),
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
