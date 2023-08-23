package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/baac183ffa684d94f697f0fc6f480e02cfb00f3d/specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/NetworkVirtualApplianceListByResourceGroup.json
func ExampleVirtualAppliancesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualAppliancesClient().NewListByResourceGroupPager("rg1", nil)
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
		// page.VirtualApplianceListResult = armnetwork.VirtualApplianceListResult{
		// 	Value: []*armnetwork.VirtualAppliance{
		// 		{
		// 			Name: to.Ptr("nva"),
		// 			Type: to.Ptr("Microsoft.Network/networkVirtualAppliances"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkVirtualAppliances/nva"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 			},
		// 			Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 			Identity: &armnetwork.ManagedServiceIdentity{
		// 				Type: to.Ptr(armnetwork.ResourceIdentityTypeUserAssigned),
		// 				UserAssignedIdentities: map[string]*armnetwork.Components1Jq1T4ISchemasManagedserviceidentityPropertiesUserassignedidentitiesAdditionalproperties{
		// 					"/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": &armnetwork.Components1Jq1T4ISchemasManagedserviceidentityPropertiesUserassignedidentitiesAdditionalproperties{
		// 					},
		// 				},
		// 			},
		// 			Properties: &armnetwork.VirtualAppliancePropertiesFormat{
		// 				AddressPrefix: to.Ptr("192.168.1.0/16"),
		// 				BootStrapConfigurationBlobs: []*string{
		// 					to.Ptr("https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrbootstrapconfig")},
		// 					CloudInitConfigurationBlobs: []*string{
		// 						to.Ptr("https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrcloudinitconfig")},
		// 						InboundSecurityRules: []*armnetwork.SubResource{
		// 							{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkVirtualAppliances/nva/InboundSecurityRules/rule1"),
		// 						}},
		// 						NvaSKU: &armnetwork.VirtualApplianceSKUProperties{
		// 							BundledScaleUnit: to.Ptr("1"),
		// 							MarketPlaceVersion: to.Ptr("12.1"),
		// 							Vendor: to.Ptr("Cisco SDWAN"),
		// 						},
		// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 						VirtualApplianceAsn: to.Ptr[int64](10000),
		// 						VirtualApplianceNics: []*armnetwork.VirtualApplianceNicProperties{
		// 							{
		// 								Name: to.Ptr("managementNic"),
		// 								PrivateIPAddress: to.Ptr("192.168.12.1"),
		// 								PublicIPAddress: to.Ptr("40.30.2.2"),
		// 							},
		// 							{
		// 								Name: to.Ptr("privateNic-1"),
		// 								PrivateIPAddress: to.Ptr("192.168.12.2"),
		// 						}},
		// 						VirtualApplianceSites: []*armnetwork.SubResource{
		// 							{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networtkVirtualAppliances/nva/virtualApplianceSites/site1"),
		// 						}},
		// 						VirtualHub: &armnetwork.SubResource{
		// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1"),
		// 						},
		// 					},
		// 			}},
		// 		}
	}
}
