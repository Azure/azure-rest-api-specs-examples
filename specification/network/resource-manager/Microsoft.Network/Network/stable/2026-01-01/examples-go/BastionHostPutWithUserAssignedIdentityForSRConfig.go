package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v12"
)

// Generated from example definition: 2026-01-01/BastionHostPutWithUserAssignedIdentityForSRConfig.json
func ExampleBastionHostsClient_BeginCreateOrUpdate_createOrUpdateBastionHostWithUserAssignedIdentityForSessionRecordingConfiguration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBastionHostsClient().BeginCreateOrUpdate(ctx, "rg1", "bastionhosttenant", armnetwork.BastionHost{
		Properties: &armnetwork.BastionHostPropertiesFormat{
			EnableSessionRecording: to.Ptr(true),
			SessionRecordingConfiguration: &armnetwork.BastionSessionRecordingConfiguration{
				Identity: &armnetwork.SessionRecordingIdentity{
					Type:                   to.Ptr(armnetwork.SessionRecordingIdentityTypeUserAssigned),
					UserAssignedIdentityID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userassignedmsi"),
				},
				BlobContainerURI: to.Ptr("https://contosostorage.blob.core.windows.net/contosocontainer"),
			},
			IPConfigurations: []*armnetwork.BastionHostIPConfiguration{
				{
					Name: to.Ptr("bastionHostIpConfiguration"),
					Properties: &armnetwork.BastionHostIPConfigurationPropertiesFormat{
						Subnet: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet2/subnets/BastionHostSubnet"),
						},
						PublicIPAddress: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/pipName"),
						},
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
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.BastionHostsClientCreateOrUpdateResponse{
	// 	BastionHost: armnetwork.BastionHost{
	// 		Name: to.Ptr("bastionhosttenant"),
	// 		ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/bastionHosts/bastionhosttenant"),
	// 		Type: to.Ptr("Microsoft.Network/bastionHosts"),
	// 		Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 		Location: to.Ptr("West US"),
	// 		SKU: &armnetwork.SKU{
	// 			Name: to.Ptr(armnetwork.BastionHostSKUNamePremium),
	// 		},
	// 		Properties: &armnetwork.BastionHostPropertiesFormat{
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			DNSName: to.Ptr("bst-9d89d361-100e-4c01-b92d-466548c476dc.bastion.azure.com"),
	// 			ScaleUnits: to.Ptr[int32](2),
	// 			DisableCopyPaste: to.Ptr(false),
	// 			EnableTunneling: to.Ptr(false),
	// 			EnableIPConnect: to.Ptr(false),
	// 			EnableShareableLink: to.Ptr(false),
	// 			EnableKerberos: to.Ptr(false),
	// 			EnableSessionRecording: to.Ptr(true),
	// 			EnablePrivateOnlyBastion: to.Ptr(false),
	// 			SessionRecordingConfiguration: &armnetwork.BastionSessionRecordingConfiguration{
	// 				Identity: &armnetwork.SessionRecordingIdentity{
	// 					Type: to.Ptr(armnetwork.SessionRecordingIdentityTypeUserAssigned),
	// 					UserAssignedIdentityID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userassignedmsi"),
	// 				},
	// 				BlobContainerURI: to.Ptr("https://contosostorage.blob.core.windows.net/contosocontainer"),
	// 			},
	// 			IPConfigurations: []*armnetwork.BastionHostIPConfiguration{
	// 				{
	// 					Name: to.Ptr("bastionHostIpConfiguration"),
	// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/bastionHosts/bastionhosttenant/bastionHostIpConfigurations/bastionHostIpConfiguration"),
	// 					Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 					Type: to.Ptr("Microsoft.Network/bastionHosts/bastionHostIpConfigurations"),
	// 					Properties: &armnetwork.BastionHostIPConfigurationPropertiesFormat{
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 						PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
	// 						Subnet: &armnetwork.SubResource{
	// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet2/subnets/BastionHostSubnet"),
	// 						},
	// 						PublicIPAddress: &armnetwork.SubResource{
	// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/pipName"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		Identity: &armnetwork.ManagedServiceIdentity{
	// 			Type: to.Ptr(armnetwork.ResourceIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armnetwork.ManagedServiceIdentityUserAssignedIdentities{
	// 				"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userassignedmsi": &armnetwork.ManagedServiceIdentityUserAssignedIdentities{
	// 					PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 					ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
