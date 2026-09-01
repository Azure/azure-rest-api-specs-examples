package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-07-15-preview/PutCompute.json
func ExampleComputesClient_BeginCreateOrUpdate_putCompute() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewComputesClient().BeginCreateOrUpdate(ctx, "rgcognitiveservices", "myAccount", "myCompute", armcognitiveservices.Compute{
		Properties: &armcognitiveservices.ClusterComputeProperties{
			ComputeType: to.Ptr(armcognitiveservices.ComputeTypeCluster),
			Location:    to.Ptr("eastus"),
			Pools: []*armcognitiveservices.Pool{
				{
					Name:         to.Ptr("default"),
					VMPriority:   to.Ptr(armcognitiveservices.VMPriorityRegular),
					InstanceType: to.Ptr("Standard_DS3_v2"),
					NodeCount:    to.Ptr[int32](2),
				},
			},
		},
		Identity: &armcognitiveservices.Identity{
			Type: to.Ptr(armcognitiveservices.ResourceIdentityTypeNone),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
}
