package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f9d14b5db982b1d554651348adc9bef4b098bdb/specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/VirtualMachines_Get.json
func ExampleVirtualMachinesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.VirtualMachine = armavs.VirtualMachine{
	// 	Name: to.Ptr("vm-209"),
	// 	Type: to.Ptr("Microsoft.AVS/privateClouds/clusters/virtualMachines"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/virtualMachines/vm-209"),
	// 	Properties: &armavs.VirtualMachineProperties{
	// 		DisplayName: to.Ptr("contoso-vm"),
	// 		FolderPath: to.Ptr("vm/folder-1"),
	// 		MoRefID: to.Ptr("vm-209"),
	// 		RestrictMovement: to.Ptr(armavs.VirtualMachineRestrictMovementStateDisabled),
	// 	},
	// }
}
