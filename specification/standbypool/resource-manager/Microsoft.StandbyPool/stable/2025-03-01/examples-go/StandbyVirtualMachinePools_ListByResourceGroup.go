package armstandbypool_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/standbypool/armstandbypool/v2"
)

// Generated from example definition: 2025-03-01/StandbyVirtualMachinePools_ListByResourceGroup.json
func ExampleStandbyVirtualMachinePoolsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstandbypool.NewClientFactory("00000000-0000-0000-0000-000000000009", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStandbyVirtualMachinePoolsClient().NewListByResourceGroupPager("rgstandbypool", nil)
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
		// page = armstandbypool.StandbyVirtualMachinePoolsClientListByResourceGroupResponse{
		// 	StandbyVirtualMachinePoolResourceListResult: armstandbypool.StandbyVirtualMachinePoolResourceListResult{
		// 		Value: []*armstandbypool.StandbyVirtualMachinePoolResource{
		// 			{
		// 				Properties: &armstandbypool.StandbyVirtualMachinePoolResourceProperties{
		// 					ElasticityProfile: &armstandbypool.StandbyVirtualMachinePoolElasticityProfile{
		// 						MaxReadyCapacity: to.Ptr[int64](304),
		// 						MinReadyCapacity: to.Ptr[int64](300),
		// 					},
		// 					VirtualMachineState: to.Ptr(armstandbypool.VirtualMachineStateRunning),
		// 					AttachedVirtualMachineScaleSetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000009/resourceGroups/rgstandbypool/providers/Microsoft.Compute/virtualMachineScaleSets/myVmss"),
		// 					ProvisioningState: to.Ptr(armstandbypool.ProvisioningStateSucceeded),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 				Location: to.Ptr("West US"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000009/resourceGroups/rgstandbypool/providers/Microsoft.StandbyPool/standbyVirtualMachinePools/pool"),
		// 				Name: to.Ptr("pool"),
		// 				Type: to.Ptr("Microsoft.StandbyPool/standbyVirtualMachinePools"),
		// 				SystemData: &armstandbypool.SystemData{
		// 					CreatedBy: to.Ptr("pooluser@microsoft.com"),
		// 					CreatedByType: to.Ptr(armstandbypool.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-07T16:33:22.210Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("pooluser@microsoft.com"),
		// 					LastModifiedByType: to.Ptr(armstandbypool.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-07T16:33:22.210Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://example.com/subscriptions/00000000-0000-0000-0000-000000000009/resourceGroups/rgstandbypool/providers/Microsoft.StandbyPool/standbyVirtualMachinePools"),
		// 	},
		// }
	}
}
