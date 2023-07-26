package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/925ba149e17454ce91ecd3f9f4134effb2f97844/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkTaps_Get_MaximumSet_Gen.json
func ExampleNetworkTapsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNetworkTapsClient().Get(ctx, "example-rg", "example-networkTap", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkTap = armmanagednetworkfabric.NetworkTap{
	// 	Name: to.Ptr("example-networkTap"),
	// 	Type: to.Ptr("microsoft.managednetworkfabric/networkTaps"),
	// 	ID: to.Ptr("/subscriptions/xxxx-xxxx-xxxx-xxxx/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkTaps/example-networkTap"),
	// 	SystemData: &armmanagednetworkfabric.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-08T16:02:57.578Z"); return t}()),
	// 		CreatedBy: to.Ptr("email@address.com"),
	// 		CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-08T16:02:57.579Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("email@address.com"),
	// 		LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastuseuap"),
	// 	Tags: map[string]*string{
	// 		"key6024": to.Ptr("1234"),
	// 	},
	// 	Properties: &armmanagednetworkfabric.NetworkTapProperties{
	// 		Annotation: to.Ptr("annotation"),
	// 		AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
	// 		ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
	// 		Destinations: []*armmanagednetworkfabric.NetworkTapPropertiesDestinationsItem{
	// 			{
	// 				Name: to.Ptr("example-destinaionName"),
	// 				DestinationID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/l3IsloationDomains/example-l3Domain/internalNetworks/example-internalNetwork"),
	// 				DestinationTapRuleID: to.Ptr("/subscriptions/xxxx-xxxx-xxxx-xxxx/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkTapRules/example-destinationTapRule"),
	// 				DestinationType: to.Ptr(armmanagednetworkfabric.DestinationTypeIsolationDomain),
	// 				IsolationDomainProperties: &armmanagednetworkfabric.IsolationDomainProperties{
	// 					Encapsulation: to.Ptr(armmanagednetworkfabric.EncapsulationNone),
	// 					NeighborGroupIDs: []*string{
	// 						to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/neighborGroups/example-neighborGroup")},
	// 					},
	// 			}},
	// 			NetworkPacketBrokerID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkPacketBrokers/example-networkPacketBroker"),
	// 			PollingType: to.Ptr(armmanagednetworkfabric.PollingTypePull),
	// 			ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
	// 			SourceTapRuleID: to.Ptr("/subscriptions/xxxx-xxxx-xxxx-xxxx/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkTapRules/example-sourceTapRule"),
	// 		},
	// 	}
}
