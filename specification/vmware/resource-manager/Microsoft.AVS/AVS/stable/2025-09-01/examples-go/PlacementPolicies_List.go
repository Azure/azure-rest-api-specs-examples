package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: 2025-09-01/PlacementPolicies_List.json
func ExamplePlacementPoliciesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPlacementPoliciesClient().NewListPager("group1", "cloud1", "cluster1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armavs.PlacementPoliciesClientListResponse{
		// 	PlacementPoliciesList: armavs.PlacementPoliciesList{
		// 		Value: []*armavs.PlacementPolicy{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/placementPolicies/policy1"),
		// 				Name: to.Ptr("policy1"),
		// 				Type: to.Ptr("Microsoft.AVS/privateClouds/clusters/placementPolicies"),
		// 				Properties: &armavs.VMHostPlacementPolicyProperties{
		// 					DisplayName: to.Ptr("policy1"),
		// 					Type: to.Ptr(armavs.PlacementPolicyTypeVMHost),
		// 					State: to.Ptr(armavs.PlacementPolicyStateEnabled),
		// 					VMMembers: []*string{
		// 						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/virtualMachines/vm-128"),
		// 						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/virtualMachines/vm-256"),
		// 					},
		// 					HostMembers: []*string{
		// 						to.Ptr("fakehost22.nyc1.kubernetes.center"),
		// 						to.Ptr("fakehost23.nyc1.kubernetes.center"),
		// 						to.Ptr("fakehost24.nyc1.kubernetes.center"),
		// 					},
		// 					AffinityType: to.Ptr(armavs.AffinityTypeAntiAffinity),
		// 					AffinityStrength: to.Ptr(armavs.AffinityStrengthMust),
		// 					AzureHybridBenefitType: to.Ptr(armavs.AzureHybridBenefitTypeSQLHost),
		// 					ProvisioningState: to.Ptr(armavs.PlacementPolicyProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/placementPolicies/policy2"),
		// 				Name: to.Ptr("policy2"),
		// 				Type: to.Ptr("Microsoft.AVS/privateClouds/clusters/placementPolicies"),
		// 				Properties: &armavs.VMPlacementPolicyProperties{
		// 					DisplayName: to.Ptr("policy2"),
		// 					Type: to.Ptr(armavs.PlacementPolicyTypeVMVM),
		// 					State: to.Ptr(armavs.PlacementPolicyStateEnabled),
		// 					VMMembers: []*string{
		// 						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/virtualMachines/vm-128"),
		// 						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/virtualMachines/vm-256"),
		// 					},
		// 					AffinityType: to.Ptr(armavs.AffinityTypeAffinity),
		// 					ProvisioningState: to.Ptr(armavs.PlacementPolicyProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
