package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/NetworkTaps_ListBySubscription.json
func ExampleNetworkTapsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("0000ABCD-0A0B-0000-0000-000000ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNetworkTapsClient().NewListBySubscriptionPager(nil)
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
		// page = armmanagednetworkfabric.NetworkTapsClientListBySubscriptionResponse{
		// 	NetworkTapsListResult: armmanagednetworkfabric.NetworkTapsListResult{
		// 		Value: []*armmanagednetworkfabric.NetworkTap{
		// 			{
		// 				Properties: &armmanagednetworkfabric.NetworkTapProperties{
		// 					Annotation: to.Ptr("annotation"),
		// 					NetworkPacketBrokerID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkPacketBrokers/example-networkPacketBroker"),
		// 					SourceTapRuleID: to.Ptr("/subscriptions/xxxx-xxxx-xxxx-xxxx/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkTapRules/example-sourceTapRule"),
		// 					NetworkFabricIDs: []*string{
		// 						to.Ptr("/subscriptions/xxxxxx/resourceGroups/resourcegroupname/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-networkFabric"),
		// 					},
		// 					Destinations: []*armmanagednetworkfabric.DestinationProperties{
		// 						{
		// 							Name: to.Ptr("example-destinaionName"),
		// 							DestinationType: to.Ptr(armmanagednetworkfabric.DestinationTypeIsolationDomain),
		// 							DestinationID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/l3IsloationDomains/example-l3Domain/internalNetworks/example-internalNetwork"),
		// 							IsolationDomainProperties: &armmanagednetworkfabric.IsolationDomainProperties{
		// 								Encapsulation: to.Ptr(armmanagednetworkfabric.EncapsulationNone),
		// 								NeighborGroupIDs: []*string{
		// 									to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/neighborGroups/example-neighborGroup"),
		// 								},
		// 							},
		// 							DestinationTapRuleID: to.Ptr("/subscriptions/xxxx-xxxx-xxxx-xxxx/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkTapRules/example-destinationTapRule"),
		// 						},
		// 					},
		// 					PollingType: to.Ptr(armmanagednetworkfabric.PollingTypePull),
		// 					LastOperation: &armmanagednetworkfabric.LastOperationProperties{
		// 						Details: to.Ptr("Succeeded"),
		// 					},
		// 					ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
		// 					ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateAccepted),
		// 					AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
		// 				},
		// 				Identity: &armmanagednetworkfabric.ManagedServiceIdentity{
		// 					PrincipalID: to.Ptr("0000ABCD-0A0B-0000-0000-000000ABCDEF"),
		// 					TenantID: to.Ptr("0000ABCD-0A0B-0000-0000-000000ABCDEF"),
		// 					Type: to.Ptr(armmanagednetworkfabric.ManagedServiceIdentityTypeNone),
		// 					UserAssignedIdentities: map[string]*armmanagednetworkfabric.UserAssignedIdentity{
		// 						"key3673": &armmanagednetworkfabric.UserAssignedIdentity{
		// 							PrincipalID: to.Ptr("0000ABCD-0A0B-0000-0000-000000ABCDEF"),
		// 							ClientID: to.Ptr("0000ABCD-0A0B-0000-0000-000000ABCDEF"),
		// 						},
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 					"keyId": to.Ptr("keyValue"),
		// 				},
		// 				Location: to.Ptr("eastuseuap"),
		// 				ID: to.Ptr("/subscriptions/xxxx-xxxx-xxxx-xxxx/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkTaps/example-networkTap"),
		// 				Name: to.Ptr("example-networkTap"),
		// 				Type: to.Ptr("microsoft.managednetworkfabric/networkTaps"),
		// 				SystemData: &armmanagednetworkfabric.SystemData{
		// 					CreatedBy: to.Ptr("email@address.com"),
		// 					CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-09T04:51:41.251Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("UserId"),
		// 					LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-09T04:51:41.251Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/aukskqxh"),
		// 	},
		// }
	}
}
