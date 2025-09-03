package armazurestackhcivm_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhcivm"
)

// Generated from example definition: 2025-06-01-preview/VirtualMachineInstances_List.json
func ExampleVirtualMachineInstancesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhcivm.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armazurestackhcivm.VirtualMachineInstancesClientListResponse{
		// 	VirtualMachineInstanceListResult: armazurestackhcivm.VirtualMachineInstanceListResult{
		// 		Value: []*armazurestackhcivm.VirtualMachineInstance{
		// 			{
		// 				Name: to.Ptr("default"),
		// 				Type: to.Ptr("Microsoft.AzureStackHCI/virtualMachineInstances"),
		// 				ExtendedLocation: &armazurestackhcivm.ExtendedLocation{
		// 					Name: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location"),
		// 					Type: to.Ptr(armazurestackhcivm.ExtendedLocationTypesCustomLocation),
		// 				},
		// 				ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM/providers/Microsoft.AzureStackHCI/virtualMachineInstances/default"),
		// 				Properties: &armazurestackhcivm.VirtualMachineInstanceProperties{
		// 					HardwareProfile: &armazurestackhcivm.VirtualMachineInstancePropertiesHardwareProfile{
		// 						VMSize: to.Ptr(armazurestackhcivm.VMSizeEnumDefault),
		// 					},
		// 					NetworkProfile: &armazurestackhcivm.VirtualMachineInstancePropertiesNetworkProfile{
		// 						NetworkInterfaces: []*armazurestackhcivm.NetworkInterfaceArmReference{
		// 							{
		// 								ID: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/networkInterfaces/test-nic"),
		// 							},
		// 						},
		// 					},
		// 					OSProfile: &armazurestackhcivm.VirtualMachineInstancePropertiesOsProfile{
		// 						AdminUsername: to.Ptr("localadmin"),
		// 						ComputerName: to.Ptr("luamaster"),
		// 					},
		// 					ProvisioningState: to.Ptr(armazurestackhcivm.ProvisioningStateEnumAccepted),
		// 					StorageProfile: &armazurestackhcivm.VirtualMachineInstancePropertiesStorageProfile{
		// 						ImageReference: &armazurestackhcivm.ImageArmReference{
		// 							ID: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/galleryImages/test-gallery-image"),
		// 						},
		// 						VMConfigStoragePathID: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/storageContainers/test-container"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
