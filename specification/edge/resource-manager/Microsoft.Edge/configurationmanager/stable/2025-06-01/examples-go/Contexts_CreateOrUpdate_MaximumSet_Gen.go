package armworkloadorchestration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadorchestration/armworkloadorchestration"
)

// Generated from example definition: 2025-06-01/Contexts_CreateOrUpdate_MaximumSet_Gen.json
func ExampleContextsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadorchestration.NewClientFactory("9D54FE4C-00AF-4836-8F48-B6A9C4E47192", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewContextsClient().BeginCreateOrUpdate(ctx, "rgconfigurationmanager", "testname", armworkloadorchestration.Context{
		Properties: &armworkloadorchestration.ContextProperties{
			Capabilities: []*armworkloadorchestration.Capability{
				{
					Name:        to.Ptr("tpylinjcmlnycfpofpxjtqmt"),
					Description: to.Ptr("banbenutsngwytoqh"),
					State:       to.Ptr(armworkloadorchestration.ResourceStateActive),
				},
			},
			Hierarchies: []*armworkloadorchestration.Hierarchy{
				{
					Name:        to.Ptr("upqe"),
					Description: to.Ptr("vg"),
				},
			},
		},
		Tags: map[string]*string{
			"key3046": to.Ptr("clcnhzwypk"),
		},
		Location: to.Ptr("pkquwbplcp"),
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
	// res = armworkloadorchestration.ContextsClientCreateOrUpdateResponse{
	// 	Context: &armworkloadorchestration.Context{
	// 		Properties: &armworkloadorchestration.ContextProperties{
	// 			Capabilities: []*armworkloadorchestration.Capability{
	// 				{
	// 					Name: to.Ptr("tpylinjcmlnycfpofpxjtqmt"),
	// 					Description: to.Ptr("banbenutsngwytoqh"),
	// 					State: to.Ptr(armworkloadorchestration.ResourceStateActive),
	// 				},
	// 			},
	// 			Hierarchies: []*armworkloadorchestration.Hierarchy{
	// 				{
	// 					Name: to.Ptr("upqe"),
	// 					Description: to.Ptr("vg"),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armworkloadorchestration.ProvisioningStateSucceeded),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key3046": to.Ptr("clcnhzwypk"),
	// 		},
	// 		Location: to.Ptr("pkquwbplcp"),
	// 		ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"),
	// 		Name: to.Ptr("dtpqgxwo"),
	// 		Type: to.Ptr("odtmcknffvktfmdhxzhjfrvupezs"),
	// 		SystemData: &armworkloadorchestration.SystemData{
	// 			CreatedBy: to.Ptr("nvjczgdguyvllp"),
	// 			CreatedByType: to.Ptr(armworkloadorchestration.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-09T10:11:50.747Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("uzbznzjgvaspvtqhyg"),
	// 			LastModifiedByType: to.Ptr(armworkloadorchestration.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-09T10:11:50.747Z"); return t}()),
	// 		},
	// 	},
	// }
}
