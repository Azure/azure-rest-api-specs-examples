package armmongocluster_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mongocluster/armmongocluster"
)

// Generated from example definition: 2026-09-01/MongoClusters_PatchTags.json
func ExampleMongoClustersClient_BeginUpdate_updatesStringValuedResourceTagsWithoutChangingTheExistingPremiumSSDStorageType() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmongocluster.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMongoClustersClient().BeginUpdate(ctx, "TestResourceGroup", "myMongoCluster", armmongocluster.Update{
		Tags: map[string]*string{
			"environment": to.Ptr("production"),
			"costCenter":  to.Ptr("1234"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmongocluster.MongoClustersClientUpdateResponse{
	// 	MongoCluster: armmongocluster.MongoCluster{
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DocumentDB/mongoClusters/myMongoCluster"),
	// 		Name: to.Ptr("myMongoCluster"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/mongoClusters"),
	// 		Location: to.Ptr("westus2"),
	// 		Tags: map[string]*string{
	// 			"environment": to.Ptr("production"),
	// 			"costCenter": to.Ptr("1234"),
	// 		},
	// 		Properties: &armmongocluster.Properties{
	// 			ProvisioningState: to.Ptr(armmongocluster.ProvisioningStateSucceeded),
	// 			Storage: &armmongocluster.StorageProperties{
	// 				SizeGb: to.Ptr[int64](128),
	// 				Type: to.Ptr(armmongocluster.StorageTypePremiumSSD),
	// 			},
	// 		},
	// 	},
	// }
}
