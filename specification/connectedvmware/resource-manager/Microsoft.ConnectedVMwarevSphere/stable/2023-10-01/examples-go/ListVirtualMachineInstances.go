package armconnectedvmware_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/ListVirtualMachineInstances.json
func ExampleVirtualMachineInstancesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconnectedvmware.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachineInstancesClient().NewListPager("subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM", nil)
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
		// page.VirtualMachineInstancesList = armconnectedvmware.VirtualMachineInstancesList{
		// 	Value: []*armconnectedvmware.VirtualMachineInstance{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.ConnectedVMwarevSphere/VirtualMachines"),
		// 			ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM/providers/Microsoft.ConnectedVMwarevSphere/virtualMachineInstances/default"),
		// 			ExtendedLocation: &armconnectedvmware.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.ExtendedLocation/customLocations/contoso"),
		// 				Type: to.Ptr("customLocation"),
		// 			},
		// 			Properties: &armconnectedvmware.VirtualMachineInstanceProperties{
		// 				HardwareProfile: &armconnectedvmware.HardwareProfile{
		// 					MemorySizeMB: to.Ptr[int32](4196),
		// 					NumCPUs: to.Ptr[int32](4),
		// 				},
		// 				InfrastructureProfile: &armconnectedvmware.InfrastructureProfile{
		// 					MoRefID: to.Ptr("aaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
		// 					TemplateID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/VirtualMachineTemplates/WebFrontEndTemplate"),
		// 					VCenterID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/VCenters/ContosoVCenter"),
		// 				},
		// 				OSProfile: &armconnectedvmware.OsProfileForVMInstance{
		// 					ComputerName: to.Ptr("DemoVM"),
		// 					OSType: to.Ptr(armconnectedvmware.OsTypeWindows),
		// 				},
		// 				PlacementProfile: &armconnectedvmware.PlacementProfile{
		// 					ResourcePoolID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/ResourcePools/HRPool"),
		// 				},
		// 				ProvisioningState: to.Ptr(armconnectedvmware.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
