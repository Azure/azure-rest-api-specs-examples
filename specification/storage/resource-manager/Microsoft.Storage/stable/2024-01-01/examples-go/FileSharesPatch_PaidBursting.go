package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9d3547622288137fd36f086afcdaea5408dbe48c/specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/FileSharesPatch_PaidBursting.json
func ExampleFileSharesClient_Update_updateSharePaidBursting() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFileSharesClient().Update(ctx, "res3376", "sto328", "share6185", armstorage.FileShare{
		FileShareProperties: &armstorage.FileShareProperties{
			FileSharePaidBursting: &armstorage.FileSharePropertiesFileSharePaidBursting{
				PaidBurstingEnabled:           to.Ptr(true),
				PaidBurstingMaxBandwidthMibps: to.Ptr[int32](10340),
				PaidBurstingMaxIops:           to.Ptr[int32](102400),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FileShare = armstorage.FileShare{
	// 	Name: to.Ptr("share6185"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/fileServices/shares"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res3376/providers/Microsoft.Storage/storageAccounts/sto328/fileServices/default/shares/share6185"),
	// 	FileShareProperties: &armstorage.FileShareProperties{
	// 		FileSharePaidBursting: &armstorage.FileSharePropertiesFileSharePaidBursting{
	// 			PaidBurstingEnabled: to.Ptr(true),
	// 			PaidBurstingMaxBandwidthMibps: to.Ptr[int32](10340),
	// 			PaidBurstingMaxIops: to.Ptr[int32](102400),
	// 		},
	// 	},
	// }
}
