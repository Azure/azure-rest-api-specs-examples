package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/PrivateEndpointDnsZoneGroupList.json
func ExamplePrivateDNSZoneGroupsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateDNSZoneGroupsClient().NewListPager("testPe", "rg1", nil)
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
		// page.PrivateDNSZoneGroupListResult = armnetwork.PrivateDNSZoneGroupListResult{
		// 	Value: []*armnetwork.PrivateDNSZoneGroup{
		// 		{
		// 			ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPe/privateDnsZoneGroups/testPdnsgroup1"),
		// 			Name: to.Ptr("testPdnsgroup1"),
		// 			Properties: &armnetwork.PrivateDNSZoneGroupPropertiesFormat{
		// 				PrivateDNSZoneConfigs: []*armnetwork.PrivateDNSZoneConfig{
		// 					{
		// 						Properties: &armnetwork.PrivateDNSZonePropertiesFormat{
		// 							PrivateDNSZoneID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateDnsZones/zone1.com"),
		// 							RecordSets: []*armnetwork.RecordSet{
		// 								{
		// 									Fqdn: to.Ptr("abc.zone1.com"),
		// 									IPAddresses: []*string{
		// 										to.Ptr("10.0.0.4")},
		// 										RecordSetName: to.Ptr("abc"),
		// 										RecordType: to.Ptr("A"),
		// 									},
		// 									{
		// 										Fqdn: to.Ptr("abc2.zone1.com"),
		// 										IPAddresses: []*string{
		// 											to.Ptr("10.0.0.5")},
		// 											RecordSetName: to.Ptr("abc2"),
		// 											RecordType: to.Ptr("A"),
		// 									}},
		// 								},
		// 						}},
		// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					},
		// 				},
		// 				{
		// 					ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPe/privateDnsZoneGroups/testPdnsgroup2"),
		// 					Name: to.Ptr("testPdnsgroup2"),
		// 					Properties: &armnetwork.PrivateDNSZoneGroupPropertiesFormat{
		// 						PrivateDNSZoneConfigs: []*armnetwork.PrivateDNSZoneConfig{
		// 							{
		// 								Properties: &armnetwork.PrivateDNSZonePropertiesFormat{
		// 									PrivateDNSZoneID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateDnsZones/zone2.com"),
		// 									RecordSets: []*armnetwork.RecordSet{
		// 										{
		// 											Fqdn: to.Ptr("abc3.zone2.com"),
		// 											IPAddresses: []*string{
		// 												to.Ptr("10.0.0.6")},
		// 												RecordSetName: to.Ptr("abc3"),
		// 												RecordType: to.Ptr("A"),
		// 											},
		// 											{
		// 												Fqdn: to.Ptr("abc4.zone2.com"),
		// 												IPAddresses: []*string{
		// 													to.Ptr("10.0.0.7")},
		// 													RecordSetName: to.Ptr("abc4"),
		// 													RecordType: to.Ptr("A"),
		// 											}},
		// 										},
		// 								}},
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							},
		// 					}},
		// 				}
	}
}
