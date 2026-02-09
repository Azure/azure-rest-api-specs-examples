package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkVirtualApplianceListBySubscription.json
func ExampleVirtualAppliancesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualAppliancesClient().NewListPager(nil)
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
		// 				AdditionalNics: []*armnetwork.VirtualApplianceAdditionalNicProperties{
		// 					{
		// 						Name: to.Ptr("exrsdwan"),
		// 						HasPublicIP: to.Ptr(true),
		// 				}},
		// 				AddressPrefix: to.Ptr("192.168.1.0/16"),
		// 				BootStrapConfigurationBlobs: []*string{
		// 					to.Ptr("https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrbootstrapconfig")},
		// 					CloudInitConfigurationBlobs: []*string{
		// 						to.Ptr("https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrcloudinitconfig")},
		// 						InboundSecurityRules: []*armnetwork.SubResource{
		// 							{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkVirtualAppliances/nva/InboundSecurityRules/rule1"),
		// 						}},
		// 						InternetIngressPublicIPs: []*armnetwork.InternetIngressPublicIPsProperties{
		// 							{
		// 								ID: to.Ptr("/subscriptions/{{subscriptionId}}/resourceGroups/{{rg}}/providers/Microsoft.Network/publicIPAddresses/slbip"),
		// 						}},
		// 						NetworkProfile: &armnetwork.VirtualAppliancePropertiesFormatNetworkProfile{
		// 							NetworkInterfaceConfigurations: []*armnetwork.VirtualApplianceNetworkInterfaceConfiguration{
		// 								{
		// 									NicType: to.Ptr(armnetwork.NicTypeInRequestPublicNic),
		// 									Properties: &armnetwork.VirtualApplianceNetworkInterfaceConfigurationProperties{
		// 										IPConfigurations: []*armnetwork.VirtualApplianceIPConfiguration{
		// 											{
		// 												Name: to.Ptr("publicnicipconfig"),
		// 												Properties: &armnetwork.VirtualApplianceIPConfigurationProperties{
		// 													Primary: to.Ptr(true),
		// 												},
		// 											},
		// 											{
		// 												Name: to.Ptr("publicnicipconfig-2"),
		// 												Properties: &armnetwork.VirtualApplianceIPConfigurationProperties{
		// 													Primary: to.Ptr(false),
		// 												},
		// 										}},
		// 									},
		// 								},
		// 								{
		// 									NicType: to.Ptr(armnetwork.NicTypeInRequestPrivateNic),
		// 									Properties: &armnetwork.VirtualApplianceNetworkInterfaceConfigurationProperties{
		// 										IPConfigurations: []*armnetwork.VirtualApplianceIPConfiguration{
		// 											{
		// 												Name: to.Ptr("privatenicipconfig"),
		// 												Properties: &armnetwork.VirtualApplianceIPConfigurationProperties{
		// 													Primary: to.Ptr(true),
		// 												},
		// 											},
		// 											{
		// 												Name: to.Ptr("privatenicipconfig-2"),
		// 												Properties: &armnetwork.VirtualApplianceIPConfigurationProperties{
		// 													Primary: to.Ptr(false),
		// 												},
		// 										}},
		// 									},
		// 							}},
		// 						},
		// 						NvaSKU: &armnetwork.VirtualApplianceSKUProperties{
		// 							BundledScaleUnit: to.Ptr("1"),
		// 							MarketPlaceVersion: to.Ptr("12.1"),
		// 							Vendor: to.Ptr("Cisco SDWAN"),
		// 						},
		// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 						VirtualApplianceAsn: to.Ptr[int64](10000),
		// 						VirtualApplianceNics: []*armnetwork.VirtualApplianceNicProperties{
		// 							{
		// 								Name: to.Ptr("publicnicipconfig"),
		// 								InstanceName: to.Ptr("nva_0"),
		// 								NicType: to.Ptr(armnetwork.NicTypeInResponsePublicNic),
		// 								PrivateIPAddress: to.Ptr("192.168.12.1"),
		// 								PublicIPAddress: to.Ptr("40.30.2.2"),
		// 							},
		// 							{
		// 								Name: to.Ptr("publicnicipconfig-2"),
		// 								InstanceName: to.Ptr("nva_0"),
		// 								NicType: to.Ptr(armnetwork.NicTypeInResponsePublicNic),
		// 								PrivateIPAddress: to.Ptr("192.168.12.2"),
		// 								PublicIPAddress: to.Ptr("40.30.2.3"),
		// 							},
		// 							{
		// 								Name: to.Ptr("privatenicipconfig"),
		// 								InstanceName: to.Ptr("nva_0"),
		// 								NicType: to.Ptr(armnetwork.NicTypeInResponsePrivateNic),
		// 								PrivateIPAddress: to.Ptr("192.168.12.3"),
		// 								PublicIPAddress: to.Ptr(""),
		// 							},
		// 							{
		// 								Name: to.Ptr("privatenicipconfig-2"),
		// 								InstanceName: to.Ptr("nva_0"),
		// 								NicType: to.Ptr(armnetwork.NicTypeInResponsePrivateNic),
		// 								PrivateIPAddress: to.Ptr("192.168.12.4"),
		// 								PublicIPAddress: to.Ptr(""),
		// 							},
		// 							{
		// 								Name: to.Ptr("exrsdwan"),
		// 								InstanceName: to.Ptr("nva_0"),
		// 								NicType: to.Ptr(armnetwork.NicTypeInResponseAdditionalNic),
		// 								PrivateIPAddress: to.Ptr("10.1.113.4"),
		// 								PublicIPAddress: to.Ptr("4.231.25.19"),
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
