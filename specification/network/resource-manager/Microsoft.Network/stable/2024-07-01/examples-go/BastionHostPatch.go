package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c58fa033619b12c7cfa8a0ec5a9bf03bb18869ab/specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/BastionHostPatch.json
func ExampleBastionHostsClient_BeginUpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBastionHostsClient().BeginUpdateTags(ctx, "rg1", "bastionhosttenant", armnetwork.TagsObject{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
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
	// res.BastionHost = armnetwork.BastionHost{
	// 	Name: to.Ptr("bastionhosttenant"),
	// 	Type: to.Ptr("Microsoft.Network/bastionHosts"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/bastionHosts/bastionhosttenant"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.BastionHostPropertiesFormat{
	// 		DisableCopyPaste: to.Ptr(false),
	// 		DNSName: to.Ptr("bst-9d89d361-100e-4c01-b92d-466548c476dc.bastion.azure.com"),
	// 		EnableIPConnect: to.Ptr(false),
	// 		EnableKerberos: to.Ptr(false),
	// 		EnablePrivateOnlyBastion: to.Ptr(false),
	// 		EnableSessionRecording: to.Ptr(false),
	// 		EnableShareableLink: to.Ptr(false),
	// 		EnableTunneling: to.Ptr(false),
	// 		IPConfigurations: []*armnetwork.BastionHostIPConfiguration{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/bastionHosts/bastionhosttenant/bastionHostIpConfigurations/bastionHostIpConfiguration"),
	// 				Name: to.Ptr("bastionHostIpConfiguration"),
	// 				Type: to.Ptr("Microsoft.Network/bastionHosts/bastionHostIpConfigurations"),
	// 				Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 				Properties: &armnetwork.BastionHostIPConfigurationPropertiesFormat{
	// 					PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					PublicIPAddress: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/pipName"),
	// 					},
	// 					Subnet: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet2/subnets/BastionHostSubnet"),
	// 					},
	// 				},
	// 		}},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		ScaleUnits: to.Ptr[int32](2),
	// 	},
	// 	SKU: &armnetwork.SKU{
	// 		Name: to.Ptr(armnetwork.BastionHostSKUNameStandard),
	// 	},
	// }
}
