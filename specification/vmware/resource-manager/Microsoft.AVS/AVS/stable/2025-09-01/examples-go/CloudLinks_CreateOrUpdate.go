package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: 2025-09-01/CloudLinks_CreateOrUpdate.json
func ExampleCloudLinksClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCloudLinksClient().BeginCreateOrUpdate(ctx, "group1", "cloud1", "cloudLink1", armavs.CloudLink{
		Properties: &armavs.CloudLinkProperties{
			LinkedCloud: to.Ptr("/subscriptions/12341234-1234-1234-1234-123412341234/resourceGroups/mygroup/providers/Microsoft.AVS/privateClouds/cloud2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armavs.CloudLinksClientCreateOrUpdateResponse{
	// 	CloudLink: &armavs.CloudLink{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/cloudLinks/cloudLink1"),
	// 		Name: to.Ptr("cloudLink1"),
	// 		Properties: &armavs.CloudLinkProperties{
	// 			Status: to.Ptr(armavs.CloudLinkStatusActive),
	// 			LinkedCloud: to.Ptr("/subscriptions/12341234-1234-1234-1234-123412341234/resourceGroups/mygroup/providers/Microsoft.AVS/privateClouds/cloud2"),
	// 		},
	// 		Type: to.Ptr("Microsoft.AVS/privateClouds/cloudLinks"),
	// 	},
	// }
}
