package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/baac183ffa684d94f697f0fc6f480e02cfb00f3d/specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/DscpConfigurationGet.json
func ExampleDscpConfigurationClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDscpConfigurationClient().Get(ctx, "rg1", "mydscpConfig", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DscpConfiguration = armnetwork.DscpConfiguration{
	// 	Name: to.Ptr("mydscpConfig"),
	// 	Type: to.Ptr("Microsoft.Network/dscpConfiguration"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/dscpConfiguration/mydscpConfig"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armnetwork.DscpConfigurationPropertiesFormat{
	// 		AssociatedNetworkInterfaces: []*armnetwork.Interface{
	// 			{
	// 			},
	// 			{
	// 		}},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		QosCollectionID: to.Ptr("0f8fad5b-d9cb-469f-a165-70867728950e"),
	// 		QosDefinitionCollection: []*armnetwork.QosDefinition{
	// 			{
	// 				DestinationIPRanges: []*armnetwork.QosIPRange{
	// 					{
	// 						EndIP: to.Ptr("127.0.10.2"),
	// 						StartIP: to.Ptr("127.0.10.1"),
	// 				}},
	// 				DestinationPortRanges: []*armnetwork.QosPortRange{
	// 					{
	// 						End: to.Ptr[int32](62),
	// 						Start: to.Ptr[int32](61),
	// 				}},
	// 				Markings: []*int32{
	// 					to.Ptr[int32](1)},
	// 					SourceIPRanges: []*armnetwork.QosIPRange{
	// 						{
	// 							EndIP: to.Ptr("127.0.0.2"),
	// 							StartIP: to.Ptr("127.0.0.1"),
	// 					}},
	// 					SourcePortRanges: []*armnetwork.QosPortRange{
	// 						{
	// 							End: to.Ptr[int32](12),
	// 							Start: to.Ptr[int32](11),
	// 					}},
	// 					Protocol: to.Ptr(armnetwork.ProtocolTypeTCP),
	// 				},
	// 				{
	// 					DestinationIPRanges: []*armnetwork.QosIPRange{
	// 						{
	// 							EndIP: to.Ptr("12.0.10.2"),
	// 							StartIP: to.Ptr("12.0.10.1"),
	// 					}},
	// 					DestinationPortRanges: []*armnetwork.QosPortRange{
	// 						{
	// 							End: to.Ptr[int32](52),
	// 							Start: to.Ptr[int32](51),
	// 					}},
	// 					Markings: []*int32{
	// 						to.Ptr[int32](2)},
	// 						SourceIPRanges: []*armnetwork.QosIPRange{
	// 							{
	// 								EndIP: to.Ptr("12.0.0.2"),
	// 								StartIP: to.Ptr("12.0.0.1"),
	// 						}},
	// 						SourcePortRanges: []*armnetwork.QosPortRange{
	// 							{
	// 								End: to.Ptr[int32](12),
	// 								Start: to.Ptr[int32](11),
	// 						}},
	// 						Protocol: to.Ptr(armnetwork.ProtocolTypeUDP),
	// 				}},
	// 			},
	// 		}
}
