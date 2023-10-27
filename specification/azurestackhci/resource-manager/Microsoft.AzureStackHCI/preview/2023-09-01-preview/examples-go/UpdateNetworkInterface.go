package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/UpdateNetworkInterface.json
func ExampleNetworkInterfacesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkInterfacesClient().BeginUpdate(ctx, "test-rg", "test-nic", armazurestackhci.NetworkInterfacesUpdateRequest{
		Tags: map[string]*string{
			"additionalProperties": to.Ptr("sample"),
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
	// res.NetworkInterfaces = armazurestackhci.NetworkInterfaces{
	// 	Name: to.Ptr("test-nic"),
	// 	Type: to.Ptr("Microsoft.AzureStackHCI/networkInterfaces"),
	// 	ID: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/networkInterfaces/test-nic"),
	// 	Location: to.Ptr("West US2"),
	// 	ExtendedLocation: &armazurestackhci.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location"),
	// 		Type: to.Ptr(armazurestackhci.ExtendedLocationTypesCustomLocation),
	// 	},
	// 	Properties: &armazurestackhci.NetworkInterfaceProperties{
	// 		IPConfigurations: []*armazurestackhci.IPConfiguration{
	// 			{
	// 				Name: to.Ptr("ipconfig-sample"),
	// 				Properties: &armazurestackhci.IPConfigurationProperties{
	// 					Subnet: &armazurestackhci.IPConfigurationPropertiesSubnet{
	// 						ID: to.Ptr("test-lnet"),
	// 					},
	// 				},
	// 		}},
	// 		ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateEnumSucceeded),
	// 	},
	// }
}
