package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: 2025-09-01/VirtualMachines_Get.json
func ExampleVirtualMachinesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachinesClient().Get(ctx, "group1", "cloud1", "cluster1", "vm-209", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armavs.VirtualMachinesClientGetResponse{
	// 	VirtualMachine: &armavs.VirtualMachine{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/virtualMachines/vm-209"),
	// 		Name: to.Ptr("vm-209"),
	// 		Type: to.Ptr("Microsoft.AVS/privateClouds/clusters/virtualMachines"),
	// 		Properties: &armavs.VirtualMachineProperties{
	// 			DisplayName: to.Ptr("contoso-vm"),
	// 			MoRefID: to.Ptr("vm-209"),
	// 			FolderPath: to.Ptr("vm/folder-1"),
	// 			RestrictMovement: to.Ptr(armavs.VirtualMachineRestrictMovementStateDisabled),
	// 		},
	// 	},
	// }
}
