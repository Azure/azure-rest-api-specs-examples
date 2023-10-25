package armresourcemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcemover/armresourcemover"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveResources_Create.json
func ExampleMoveResourcesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcemover.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMoveResourcesClient().BeginCreate(ctx, "rg1", "movecollection1", "moveresourcename1", &armresourcemover.MoveResourcesClientBeginCreateOptions{Body: &armresourcemover.MoveResource{
		Properties: &armresourcemover.MoveResourceProperties{
			DependsOnOverrides: []*armresourcemover.MoveResourceDependencyOverride{
				{
					ID:       to.Ptr("/subscriptions/c4488a3f-a7f7-4ad4-aa72-0e1f4d9c0756/resourceGroups/eastusRG/providers/Microsoft.Network/networkInterfaces/eastusvm140"),
					TargetID: to.Ptr("/subscriptions/c4488a3f-a7f7-4ad4-aa72-0e1f4d9c0756/resourceGroups/westusRG/providers/Microsoft.Network/networkInterfaces/eastusvm140"),
				}},
			ResourceSettings: &armresourcemover.VirtualMachineResourceSettings{
				ResourceType:            to.Ptr("Microsoft.Compute/virtualMachines"),
				TargetResourceName:      to.Ptr("westusvm1"),
				TargetAvailabilitySetID: to.Ptr("/subscriptions/subid/resourceGroups/eastusRG/providers/Microsoft.Compute/availabilitySets/avset1"),
				TargetAvailabilityZone:  to.Ptr(armresourcemover.TargetAvailabilityZoneTwo),
				UserManagedIdentities: []*string{
					to.Ptr("/subscriptions/subid/resourceGroups/eastusRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/umi1")},
			},
			SourceID: to.Ptr("/subscriptions/subid/resourceGroups/eastusRG/providers/Microsoft.Compute/virtualMachines/eastusvm1"),
		},
	},
	})
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
	// res.MoveResource = armresourcemover.MoveResource{
	// 	Name: to.Ptr("moveresourcename1"),
	// 	Type: to.Ptr("Microsoft.Migrate/MoveCollections/MoveResources"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Migrate/MoveCollections/movecollection1/MoveResources/moveresource1"),
	// 	Properties: &armresourcemover.MoveResourceProperties{
	// 		DependsOn: []*armresourcemover.MoveResourceDependency{
	// 			{
	// 				ID: to.Ptr("/subscriptions/c4488a3f-a7f7-4ad4-aa72-0e1f4d9c0756/resourceGroups/eastusRG/providers/Microsoft.Network/networkInterfaces/eastusvm140"),
	// 				ResolutionStatus: to.Ptr("Unresolved"),
	// 				ResolutionType: to.Ptr(armresourcemover.ResolutionTypeAutomatic),
	// 		}},
	// 		DependsOnOverrides: []*armresourcemover.MoveResourceDependencyOverride{
	// 			{
	// 				ID: to.Ptr("/subscriptions/c4488a3f-a7f7-4ad4-aa72-0e1f4d9c0756/resourceGroups/eastusRG/providers/Microsoft.Network/networkInterfaces/eastusvm140"),
	// 				TargetID: to.Ptr("/subscriptions/c4488a3f-a7f7-4ad4-aa72-0e1f4d9c0756/resourceGroups/westusRG/providers/Microsoft.Network/networkInterfaces/eastusvm140"),
	// 		}},
	// 		ProvisioningState: to.Ptr(armresourcemover.ProvisioningStateSucceeded),
	// 		ResourceSettings: &armresourcemover.VirtualMachineResourceSettings{
	// 			ResourceType: to.Ptr("Microsoft.Compute/virtualMachines"),
	// 			TargetResourceGroupName: to.Ptr("rg1"),
	// 			TargetResourceName: to.Ptr("westusvm1"),
	// 			TargetAvailabilitySetID: to.Ptr("/subscriptions/subid/resourceGroups/eastusRG/providers/Microsoft.Compute/availabilitySets/avset1"),
	// 			TargetAvailabilityZone: to.Ptr(armresourcemover.TargetAvailabilityZoneTwo),
	// 			UserManagedIdentities: []*string{
	// 				to.Ptr("/subscriptions/subid/resourceGroups/eastusRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/umi1")},
	// 			},
	// 			SourceID: to.Ptr("/subscriptions/subid/resourceGroups/eastusRG/providers/Microsoft.Compute/virtualMachines/eastusvm1"),
	// 		},
	// 		SystemData: &armresourcemover.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-01T15:06:54.2757906Z"); return t}()),
	// 			CreatedBy: to.Ptr("user@microsoft.com"),
	// 			CreatedByType: to.Ptr(armresourcemover.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-27T17:16:24.3644126Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user@microsoft.com"),
	// 			LastModifiedByType: to.Ptr(armresourcemover.CreatedByTypeUser),
	// 		},
	// 	}
}
