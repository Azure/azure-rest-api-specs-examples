package armstorage_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v4"
)

// Generated from example definition: 2026-06-01/FileShareAclsPatch.json
func ExampleFileSharesClient_Update_updateShareAcls() {
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
			SignedIdentifiers: []*armstorage.SignedIdentifier{
				{
					AccessPolicy: &armstorage.AccessPolicy{
						ExpiryTime: to.Ptr(time.Date(2021, time.May, 1, 8, 49, 37, 0, time.UTC)),
						Permission: to.Ptr("rwd"),
						StartTime:  to.Ptr(time.Date(2021, time.April, 1, 8, 49, 37, 0, time.UTC)),
					},
					ID: to.Ptr("MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI"),
				},
			},
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
	// 			SignedIdentifiers: []*armstorage.SignedIdentifier{
	// 				{
	// 					AccessPolicy: &armstorage.AccessPolicy{
	// 						ExpiryTime: to.Ptr(time.Date(2021, time.May, 1, 8, 49, 37, 0, time.UTC)),
	// 						Permission: to.Ptr("rwd"),
	// 						StartTime: to.Ptr(time.Date(2021, time.April, 1, 8, 49, 37, 0, time.UTC)),
	// 					},
	// 					ID: to.Ptr("MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
