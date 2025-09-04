package armazurestackhcivm_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhcivm"
)

// Generated from example definition: 2025-06-01-preview/VirtualMachineInstances_Update.json
func ExampleVirtualMachineInstancesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhcivm.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineInstancesClient().BeginUpdate(ctx, "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM", armazurestackhcivm.VirtualMachineInstanceUpdateRequest{
		Properties: &armazurestackhcivm.VirtualMachineInstanceUpdateProperties{
			StorageProfile: &armazurestackhcivm.StorageProfileUpdate{
				DataDisks: []*armazurestackhcivm.VirtualHardDiskArmReference{
					{
						ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.AzureStackHCI/virtualHardDisks/test-vhd"),
					},
				},
			},
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
	// res = armazurestackhcivm.VirtualMachineInstancesClientUpdateResponse{
	// 	VirtualMachineInstance: &armazurestackhcivm.VirtualMachineInstance{
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.AzureStackHCI/virtualMachineInstances"),
	// 		ExtendedLocation: &armazurestackhcivm.ExtendedLocation{
	// 			Name: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location"),
	// 			Type: to.Ptr(armazurestackhcivm.ExtendedLocationTypesCustomLocation),
	// 		},
	// 		ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM/providers/Microsoft.AzureStackHCI/virtualMachineInstances/default"),
	// 		Properties: &armazurestackhcivm.VirtualMachineInstanceProperties{
	// 			HardwareProfile: &armazurestackhcivm.VirtualMachineInstancePropertiesHardwareProfile{
	// 				VMSize: to.Ptr(armazurestackhcivm.VMSizeEnumDefault),
	// 			},
	// 			NetworkProfile: &armazurestackhcivm.VirtualMachineInstancePropertiesNetworkProfile{
	// 				NetworkInterfaces: []*armazurestackhcivm.NetworkInterfaceArmReference{
	// 					{
	// 						ID: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/networkInterfaces/test-nic"),
	// 					},
	// 				},
	// 			},
	// 			OSProfile: &armazurestackhcivm.VirtualMachineInstancePropertiesOsProfile{
	// 				AdminUsername: to.Ptr("localadmin"),
	// 				ComputerName: to.Ptr("luamaster"),
	// 			},
	// 			ProvisioningState: to.Ptr(armazurestackhcivm.ProvisioningStateEnumAccepted),
	// 			StorageProfile: &armazurestackhcivm.VirtualMachineInstancePropertiesStorageProfile{
	// 				DataDisks: []*armazurestackhcivm.VirtualHardDiskArmReference{
	// 					{
	// 						ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.AzureStackHCI/virtualHardDisks/test-vhd"),
	// 					},
	// 				},
	// 				ImageReference: &armazurestackhcivm.ImageArmReference{
	// 					ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.AzureStackHCI/galleryImages/test-gallery-image"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
