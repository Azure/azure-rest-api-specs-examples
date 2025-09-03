package armazurestackhcivm_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhcivm"
)

// Generated from example definition: 2025-06-01-preview/NetworkInterfaces__CreateOrUpdate_CreateFromLocal.json
func ExampleNetworkInterfacesClient_BeginCreateOrUpdate_createNetworkInterfaceFromLocal() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhcivm.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkInterfacesClient().BeginCreateOrUpdate(ctx, "test-rg", "test-nic", armazurestackhcivm.NetworkInterface{
		ExtendedLocation: &armazurestackhcivm.ExtendedLocation{
			Name: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location"),
			Type: to.Ptr(armazurestackhcivm.ExtendedLocationTypesCustomLocation),
		},
		Location: to.Ptr("eastus"),
		Properties: &armazurestackhcivm.NetworkInterfaceProperties{
			IPConfigurations: []*armazurestackhcivm.IPConfiguration{
				{
					Name: to.Ptr("ipconfig-sample"),
					Properties: &armazurestackhcivm.IPConfigurationProperties{
						Subnet: &armazurestackhcivm.LogicalNetworkArmReference{
							ID: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/logicalNetworks/test-lnet"),
						},
					},
				},
			},
			CreateFromLocal: to.Ptr(true),
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
	// res = armazurestackhcivm.NetworkInterfacesClientCreateOrUpdateResponse{
	// 	NetworkInterface: &armazurestackhcivm.NetworkInterface{
	// 		Name: to.Ptr("test-nic"),
	// 		Type: to.Ptr("Microsoft.AzureStackHCI/networkInterfaces"),
	// 		ExtendedLocation: &armazurestackhcivm.ExtendedLocation{
	// 			Name: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location"),
	// 			Type: to.Ptr(armazurestackhcivm.ExtendedLocationTypesCustomLocation),
	// 		},
	// 		ID: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/networkInterfaces/test-nic"),
	// 		Location: to.Ptr("West US2"),
	// 		Properties: &armazurestackhcivm.NetworkInterfaceProperties{
	// 			IPConfigurations: []*armazurestackhcivm.IPConfiguration{
	// 				{
	// 					Name: to.Ptr("ipconfig-sample"),
	// 					Properties: &armazurestackhcivm.IPConfigurationProperties{
	// 						Subnet: &armazurestackhcivm.LogicalNetworkArmReference{
	// 							ID: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/logicalNetworks/test-lnet"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			CreateFromLocal: to.Ptr(true),
	// 			ProvisioningState: to.Ptr(armazurestackhcivm.ProvisioningStateEnumAccepted),
	// 		},
	// 	},
	// }
}
