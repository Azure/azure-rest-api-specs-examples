package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/260ed6a52537921f53a18ffaf4020e3b4d510367/specification/storage/resource-manager/Microsoft.Storage/stable/2025-01-01/examples/FileSharesPut_ProvisionedV2.json
func ExampleFileSharesClient_Create_putSharesProvisionedV2() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFileSharesClient().Create(ctx, "res346", "sto666", "share1235", armstorage.FileShare{
		FileShareProperties: &armstorage.FileShareProperties{
			ProvisionedBandwidthMibps: to.Ptr[int32](200),
			ProvisionedIops:           to.Ptr[int32](5000),
			ShareQuota:                to.Ptr[int32](100),
		},
	}, &armstorage.FileSharesClientCreateOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FileShare = armstorage.FileShare{
	// 	Name: to.Ptr("share1235"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/fileServices/shares"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res346/providers/Microsoft.Storage/storageAccounts/sto666/fileServices/default/shares/share1235"),
	// 	FileShareProperties: &armstorage.FileShareProperties{
	// 		IncludedBurstIops: to.Ptr[int32](15000),
	// 		MaxBurstCreditsForIops: to.Ptr[int64](36000000),
	// 		ProvisionedBandwidthMibps: to.Ptr[int32](200),
	// 		ProvisionedIops: to.Ptr[int32](5000),
	// 		ShareQuota: to.Ptr[int32](100),
	// 	},
	// }
}
