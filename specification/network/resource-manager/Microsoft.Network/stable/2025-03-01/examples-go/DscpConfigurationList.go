package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/DscpConfigurationList.json
func ExampleDscpConfigurationClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDscpConfigurationClient().NewListPager("rg1", nil)
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
		// page.DscpConfigurationListResult = armnetwork.DscpConfigurationListResult{
		// 	Value: []*armnetwork.DscpConfiguration{
		// 		{
		// 			Name: to.Ptr("mydscpConfig"),
		// 			Type: to.Ptr("Microsoft.Network/dscpConfiguration"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/dscpConfiguration/mydscpConfig"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armnetwork.DscpConfigurationPropertiesFormat{
		// 				AssociatedNetworkInterfaces: []*armnetwork.Interface{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/test-nic1"),
		// 					},
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/test-nic2"),
		// 				}},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				QosCollectionID: to.Ptr("0f8fad5b-d9cb-469f-a165-70867728950e"),
		// 				QosDefinitionCollection: []*armnetwork.QosDefinition{
		// 					{
		// 						DestinationIPRanges: []*armnetwork.QosIPRange{
		// 							{
		// 								EndIP: to.Ptr("127.0.10.2"),
		// 								StartIP: to.Ptr("127.0.10.1"),
		// 						}},
		// 						DestinationPortRanges: []*armnetwork.QosPortRange{
		// 							{
		// 								End: to.Ptr[int32](62),
		// 								Start: to.Ptr[int32](61),
		// 						}},
		// 						Markings: []*int32{
		// 							to.Ptr[int32](1)},
		// 							SourceIPRanges: []*armnetwork.QosIPRange{
		// 								{
		// 									EndIP: to.Ptr("127.0.0.2"),
		// 									StartIP: to.Ptr("127.0.0.1"),
		// 							}},
		// 							SourcePortRanges: []*armnetwork.QosPortRange{
		// 								{
		// 									End: to.Ptr[int32](12),
		// 									Start: to.Ptr[int32](11),
		// 							}},
		// 							Protocol: to.Ptr(armnetwork.ProtocolTypeTCP),
		// 						},
		// 						{
		// 							DestinationIPRanges: []*armnetwork.QosIPRange{
		// 								{
		// 									EndIP: to.Ptr("12.0.10.2"),
		// 									StartIP: to.Ptr("12.0.10.1"),
		// 							}},
		// 							DestinationPortRanges: []*armnetwork.QosPortRange{
		// 								{
		// 									End: to.Ptr[int32](52),
		// 									Start: to.Ptr[int32](51),
		// 							}},
		// 							Markings: []*int32{
		// 								to.Ptr[int32](2)},
		// 								SourceIPRanges: []*armnetwork.QosIPRange{
		// 									{
		// 										EndIP: to.Ptr("12.0.0.2"),
		// 										StartIP: to.Ptr("12.0.0.1"),
		// 								}},
		// 								SourcePortRanges: []*armnetwork.QosPortRange{
		// 									{
		// 										End: to.Ptr[int32](12),
		// 										Start: to.Ptr[int32](11),
		// 								}},
		// 								Protocol: to.Ptr(armnetwork.ProtocolTypeUDP),
		// 						}},
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("mydscpConfig2"),
		// 					Type: to.Ptr("Microsoft.Network/dscpConfiguration"),
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/dscpConfiguration/mydscpConfig2"),
		// 					Location: to.Ptr("eastus"),
		// 					Properties: &armnetwork.DscpConfigurationPropertiesFormat{
		// 						AssociatedNetworkInterfaces: []*armnetwork.Interface{
		// 							{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/test-nic3"),
		// 							},
		// 							{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/test-nic4"),
		// 						}},
		// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 						QosCollectionID: to.Ptr("9as24mf6-d9cb-7a7f-a165-70867728950e"),
		// 						QosDefinitionCollection: []*armnetwork.QosDefinition{
		// 							{
		// 								DestinationIPRanges: []*armnetwork.QosIPRange{
		// 									{
		// 										EndIP: to.Ptr("127.0.10.2"),
		// 										StartIP: to.Ptr("127.0.10.1"),
		// 								}},
		// 								DestinationPortRanges: []*armnetwork.QosPortRange{
		// 									{
		// 										End: to.Ptr[int32](62),
		// 										Start: to.Ptr[int32](61),
		// 								}},
		// 								Markings: []*int32{
		// 									to.Ptr[int32](1)},
		// 									SourceIPRanges: []*armnetwork.QosIPRange{
		// 										{
		// 											EndIP: to.Ptr("127.0.0.2"),
		// 											StartIP: to.Ptr("127.0.0.1"),
		// 									}},
		// 									SourcePortRanges: []*armnetwork.QosPortRange{
		// 										{
		// 											End: to.Ptr[int32](12),
		// 											Start: to.Ptr[int32](11),
		// 									}},
		// 									Protocol: to.Ptr(armnetwork.ProtocolTypeTCP),
		// 								},
		// 								{
		// 									DestinationIPRanges: []*armnetwork.QosIPRange{
		// 										{
		// 											EndIP: to.Ptr("12.0.10.2"),
		// 											StartIP: to.Ptr("12.0.10.1"),
		// 									}},
		// 									DestinationPortRanges: []*armnetwork.QosPortRange{
		// 										{
		// 											End: to.Ptr[int32](52),
		// 											Start: to.Ptr[int32](51),
		// 									}},
		// 									Markings: []*int32{
		// 										to.Ptr[int32](2)},
		// 										SourceIPRanges: []*armnetwork.QosIPRange{
		// 											{
		// 												EndIP: to.Ptr("12.0.0.2"),
		// 												StartIP: to.Ptr("12.0.0.1"),
		// 										}},
		// 										SourcePortRanges: []*armnetwork.QosPortRange{
		// 											{
		// 												End: to.Ptr[int32](12),
		// 												Start: to.Ptr[int32](11),
		// 										}},
		// 										Protocol: to.Ptr(armnetwork.ProtocolTypeUDP),
		// 								}},
		// 							},
		// 					}},
		// 				}
	}
}
