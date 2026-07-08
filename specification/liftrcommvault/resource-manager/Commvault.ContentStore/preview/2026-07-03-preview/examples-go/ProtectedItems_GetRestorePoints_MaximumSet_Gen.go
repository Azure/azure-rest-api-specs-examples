package armcommvaultcontentstore_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/commvaultcontentstore/armcommvaultcontentstore"
)

// Generated from example definition: 2026-07-03-preview/ProtectedItems_GetRestorePoints_MaximumSet_Gen.json
func ExampleProtectedItemsClient_GetRestorePoints() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommvaultcontentstore.NewClientFactory("65D4E6D7-7063-4C4B-BAC5-13C45474009E", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProtectedItemsClient().GetRestorePoints(ctx, "rgcommvault", "sample-cloudAccountName", "sample-protectionGroupName", "sample-protectedItemName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcommvaultcontentstore.ProtectedItemsClientGetRestorePointsResponse{
	// 	RestorePoints: armcommvaultcontentstore.RestorePoints{
	// 		RestoreTimes: []*int64{
	// 			to.Ptr[int64](6),
	// 		},
	// 	},
	// }
}
