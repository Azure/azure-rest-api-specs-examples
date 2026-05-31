package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-03-15-preview/UpdateCompute.json
func ExampleComputesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewComputesClient().BeginUpdate(ctx, "rgcognitiveservices", "myAccount", "myCompute", armcognitiveservices.Compute{
		Properties: &armcognitiveservices.ClusterComputeProperties{
			ComputeType: to.Ptr(armcognitiveservices.ComputeTypeCluster),
			Pools: []*armcognitiveservices.Pool{
				{
					Name:         to.Ptr("default"),
					VMPriority:   to.Ptr(armcognitiveservices.VMPriorityRegular),
					InstanceType: to.Ptr("Standard_DS3_v2"),
					NodeCount:    to.Ptr[int32](4),
				},
			},
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
	// res = armcognitiveservices.ComputesClientUpdateResponse{
	// 	Compute: armcognitiveservices.Compute{
	// 		Properties: &armcognitiveservices.ClusterComputeProperties{
	// 			ComputeType: to.Ptr(armcognitiveservices.ComputeTypeCluster),
	// 			Pools: []*armcognitiveservices.Pool{
	// 				{
	// 					Name: to.Ptr("default"),
	// 					VMPriority: to.Ptr(armcognitiveservices.VMPriorityRegular),
	// 					InstanceType: to.Ptr("Standard_DS3_v2"),
	// 					NodeCount: to.Ptr[int32](4),
	// 				},
	// 			},
	// 			SubnetArmID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rgcognitiveservices/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/default"),
	// 			ProvisioningState: to.Ptr(armcognitiveservices.ComputeProvisioningStateAccepted),
	// 			Errors: []*armcognitiveservices.ErrorDetail{
	// 			},
	// 			CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-24T11:39:39.795Z"); return t}()),
	// 		},
	// 		Etag: to.Ptr("\"00000000-0000-0000-0000-000000000000\""),
	// 		Location: to.Ptr("eastus"),
	// 		Identity: &armcognitiveservices.Identity{
	// 			Type: to.Ptr(armcognitiveservices.ResourceIdentityTypeNone),
	// 			TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 			PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000001"),
	// 			UserAssignedIdentities: map[string]*armcognitiveservices.UserAssignedIdentity{
	// 			},
	// 		},
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgcognitiveservices/providers/Microsoft.CognitiveServices/accounts/myAccount/computes/myCompute"),
	// 		Name: to.Ptr("myCompute"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts/computes"),
	// 		SystemData: &armcognitiveservices.SystemData{
	// 			CreatedBy: to.Ptr("xxx@microsoft.com"),
	// 			CreatedByType: to.Ptr(armcognitiveservices.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-24T11:39:39.795Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("xxx@microsoft.com"),
	// 			LastModifiedByType: to.Ptr(armcognitiveservices.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-24T11:39:39.795Z"); return t}()),
	// 		},
	// 	},
	// }
}
