package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: 2025-09-01/PlacementPolicies_Update.json
func ExamplePlacementPoliciesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPlacementPoliciesClient().BeginUpdate(ctx, "group1", "cloud1", "cluster1", "policy1", armavs.PlacementPolicyUpdate{
		Properties: &armavs.PlacementPolicyUpdateProperties{
			State: to.Ptr(armavs.PlacementPolicyStateDisabled),
			VMMembers: []*string{
				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/virtualMachines/vm-128"),
				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/virtualMachines/vm-256"),
			},
			HostMembers: []*string{
				to.Ptr("fakehost22.nyc1.kubernetes.center"),
				to.Ptr("fakehost23.nyc1.kubernetes.center"),
				to.Ptr("fakehost24.nyc1.kubernetes.center"),
			},
			AffinityStrength:       to.Ptr(armavs.AffinityStrengthMust),
			AzureHybridBenefitType: to.Ptr(armavs.AzureHybridBenefitTypeSQLHost),
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
	// res = armavs.PlacementPoliciesClientUpdateResponse{
	// 	PlacementPolicy: &armavs.PlacementPolicy{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/placementPolicies/policy1"),
	// 		Name: to.Ptr("policy1"),
	// 		Type: to.Ptr("Microsoft.AVS/privateClouds/clusters/placementPolicies"),
	// 		Properties: &armavs.VMHostPlacementPolicyProperties{
	// 			DisplayName: to.Ptr("policy1"),
	// 			Type: to.Ptr(armavs.PlacementPolicyTypeVMHost),
	// 			State: to.Ptr(armavs.PlacementPolicyStateDisabled),
	// 			VMMembers: []*string{
	// 				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/virtualMachines/vm-128"),
	// 				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/virtualMachines/vm-256"),
	// 			},
	// 			HostMembers: []*string{
	// 				to.Ptr("fakehost22.nyc1.kubernetes.center"),
	// 				to.Ptr("fakehost23.nyc1.kubernetes.center"),
	// 				to.Ptr("fakehost24.nyc1.kubernetes.center"),
	// 			},
	// 			AffinityType: to.Ptr(armavs.AffinityTypeAntiAffinity),
	// 			AffinityStrength: to.Ptr(armavs.AffinityStrengthMust),
	// 			AzureHybridBenefitType: to.Ptr(armavs.AzureHybridBenefitTypeSQLHost),
	// 			ProvisioningState: to.Ptr(armavs.PlacementPolicyProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
