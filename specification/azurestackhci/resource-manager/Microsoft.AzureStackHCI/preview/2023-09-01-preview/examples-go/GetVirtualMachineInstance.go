package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/GetVirtualMachineInstance.json
func ExampleVirtualMachineInstancesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineInstancesClient().Get(ctx, "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/Microsoft.HybridCompute/machines/DemoVM", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineInstance = armazurestackhci.VirtualMachineInstance{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.AzureStackHCI/virtualMachineInstances"),
	// 	ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/Microsoft.HybridCompute/machines/DemoVM/providers/Microsoft.AzureStackHCI/virtualMachineInstances/default"),
	// 	ExtendedLocation: &armazurestackhci.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location"),
	// 		Type: to.Ptr(armazurestackhci.ExtendedLocationTypesCustomLocation),
	// 	},
	// 	Properties: &armazurestackhci.VirtualMachineInstanceProperties{
	// 		HardwareProfile: &armazurestackhci.VirtualMachineInstancePropertiesHardwareProfile{
	// 			VMSize: to.Ptr(armazurestackhci.VMSizeEnumDefault),
	// 		},
	// 		NetworkProfile: &armazurestackhci.VirtualMachineInstancePropertiesNetworkProfile{
	// 			NetworkInterfaces: []*armazurestackhci.VirtualMachineInstancePropertiesNetworkProfileNetworkInterfacesItem{
	// 				{
	// 					ID: to.Ptr("test-nic"),
	// 			}},
	// 		},
	// 		OSProfile: &armazurestackhci.VirtualMachineInstancePropertiesOsProfile{
	// 			AdminUsername: to.Ptr("localadmin"),
	// 			ComputerName: to.Ptr("luamaster"),
	// 		},
	// 		ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateEnumSucceeded),
	// 		StorageProfile: &armazurestackhci.VirtualMachineInstancePropertiesStorageProfile{
	// 			ImageReference: &armazurestackhci.VirtualMachineInstancePropertiesStorageProfileImageReference{
	// 				ID: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/galleryImages/test-gallery-image"),
	// 			},
	// 			VMConfigStoragePathID: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/storageContainers/test-container"),
	// 		},
	// 	},
	// }
}
