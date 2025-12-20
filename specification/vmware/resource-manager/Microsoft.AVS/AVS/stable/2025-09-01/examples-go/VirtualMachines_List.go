package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: 2025-09-01/VirtualMachines_List.json
func ExampleVirtualMachinesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachinesClient().NewListPager("group1", "cloud1", "cluster1", nil)
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
		// page = armavs.VirtualMachinesClientListResponse{
		// 	VirtualMachinesList: armavs.VirtualMachinesList{
		// 		Value: []*armavs.VirtualMachine{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/virtualMachines/vm-209"),
		// 				Name: to.Ptr("vm-209"),
		// 				Type: to.Ptr("Microsoft.AVS/privateClouds/clusters/virtualMachines"),
		// 				Properties: &armavs.VirtualMachineProperties{
		// 					DisplayName: to.Ptr("contoso-vm1"),
		// 					MoRefID: to.Ptr("vm-209"),
		// 					FolderPath: to.Ptr("vm/folder-1"),
		// 					RestrictMovement: to.Ptr(armavs.VirtualMachineRestrictMovementStateDisabled),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/virtualMachines/vm-128"),
		// 				Name: to.Ptr("vm-128"),
		// 				Type: to.Ptr("Microsoft.AVS/privateClouds/clusters/virtualMachines"),
		// 				Properties: &armavs.VirtualMachineProperties{
		// 					DisplayName: to.Ptr("contoso-vm2"),
		// 					MoRefID: to.Ptr("vm-128"),
		// 					FolderPath: to.Ptr("vm"),
		// 					RestrictMovement: to.Ptr(armavs.VirtualMachineRestrictMovementStateEnabled),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
