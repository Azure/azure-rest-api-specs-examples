package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/network/resource-manager/Microsoft.Network/stable/2022-09-01/examples/NetworkVirtualApplianceGet.json
func ExampleVirtualAppliancesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualAppliancesClient().Get(ctx, "rg1", "nva", &armnetwork.VirtualAppliancesClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualAppliance = armnetwork.VirtualAppliance{
	// 	Name: to.Ptr("nva"),
	// 	Type: to.Ptr("Microsoft.Network/networkVirtualAppliances"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkVirtualAppliances/nva"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 	},
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Identity: &armnetwork.ManagedServiceIdentity{
	// 		Type: to.Ptr(armnetwork.ResourceIdentityTypeUserAssigned),
	// 		UserAssignedIdentities: map[string]*armnetwork.Components1Jq1T4ISchemasManagedserviceidentityPropertiesUserassignedidentitiesAdditionalproperties{
	// 			"/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": &armnetwork.Components1Jq1T4ISchemasManagedserviceidentityPropertiesUserassignedidentitiesAdditionalproperties{
	// 			},
	// 		},
	// 	},
	// 	Properties: &armnetwork.VirtualAppliancePropertiesFormat{
	// 		AddressPrefix: to.Ptr("192.168.1.0/16"),
	// 		BootStrapConfigurationBlobs: []*string{
	// 			to.Ptr("https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrbootstrapconfig")},
	// 			CloudInitConfigurationBlobs: []*string{
	// 				to.Ptr("https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrcloudinitconfig")},
	// 				InboundSecurityRules: []*armnetwork.SubResource{
	// 					{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkVirtualAppliances/nva/InboundSecurityRules/rule1"),
	// 				}},
	// 				NvaSKU: &armnetwork.VirtualApplianceSKUProperties{
	// 					BundledScaleUnit: to.Ptr("1"),
	// 					MarketPlaceVersion: to.Ptr("12.1"),
	// 					Vendor: to.Ptr("Cisco SDWAN"),
	// 				},
	// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 				VirtualApplianceAsn: to.Ptr[int64](10000),
	// 				VirtualApplianceNics: []*armnetwork.VirtualApplianceNicProperties{
	// 					{
	// 						Name: to.Ptr("managementNic"),
	// 						PrivateIPAddress: to.Ptr("192.168.12.1"),
	// 						PublicIPAddress: to.Ptr("40.30.2.2"),
	// 					},
	// 					{
	// 						Name: to.Ptr("privateNic-1"),
	// 						PrivateIPAddress: to.Ptr("192.168.12.2"),
	// 				}},
	// 				VirtualApplianceSites: []*armnetwork.SubResource{
	// 					{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networtkVirtualAppliances/nva/virtualApplianceSites/site1"),
	// 				}},
	// 				VirtualHub: &armnetwork.SubResource{
	// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1"),
	// 				},
	// 			},
	// 		}
}
