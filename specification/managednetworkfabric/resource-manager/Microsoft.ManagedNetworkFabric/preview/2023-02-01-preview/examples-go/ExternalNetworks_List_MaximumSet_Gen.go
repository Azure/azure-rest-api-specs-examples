package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d03c1964cb76ffd6884d10a1871bbe779a2f68ef/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/ExternalNetworks_List_MaximumSet_Gen.json
func ExampleExternalNetworksClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExternalNetworksClient().NewListPager("resourceGroupName", "example-l3domain", nil)
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
		// page.ExternalNetworksList = armmanagednetworkfabric.ExternalNetworksList{
		// 	Value: []*armmanagednetworkfabric.ExternalNetwork{
		// 		{
		// 			Name: to.Ptr("example-externalnetwork"),
		// 			Type: to.Ptr("microsoft.managednetworkfabric/example-externalnetwork"),
		// 			ID: to.Ptr("/subscriptions/xxxxxx/resourceGroups/resourcegroupname/providers/Microsoft.ManagedNetworkFabric/example-l3domain/externalNetworks/example-externalnetwork"),
		// 			SystemData: &armmanagednetworkfabric.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-09T18:35:44.070Z"); return t}()),
		// 				CreatedBy: to.Ptr("email@address.com"),
		// 				CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-09T18:35:44.070Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("User"),
		// 				LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 			},
		// 			Properties: &armmanagednetworkfabric.ExternalNetworkProperties{
		// 				Annotation: to.Ptr("lab1"),
		// 				AdministrativeState: to.Ptr(armmanagednetworkfabric.EnabledDisabledStateEnabled),
		// 				DisabledOnResources: []*string{
		// 					to.Ptr("")},
		// 					ExportRoutePolicyID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
		// 					ImportRoutePolicyID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
		// 					NetworkToNetworkInterconnectID: to.Ptr("/subscriptions/xxxxxx/resourceGroups/resourcegroupname/providers/Microsoft.ManagedNetworkFabric/networkFabrics/FabricName/networkToNetworkInterconnects/DefaultNNI"),
		// 					OptionAProperties: &armmanagednetworkfabric.ExternalNetworkPropertiesOptionAProperties{
		// 						PrimaryIPv4Prefix: to.Ptr("10.1.1.0/30"),
		// 						PrimaryIPv6Prefix: to.Ptr("3FFE:FFFF:0:CD30::a0/126"),
		// 						SecondaryIPv4Prefix: to.Ptr("10.1.1.4/30"),
		// 						SecondaryIPv6Prefix: to.Ptr("3FFE:FFFF:0:CD30::a4/126"),
		// 						BfdConfiguration: &armmanagednetworkfabric.BfdConfiguration{
		// 							AdministrativeState: to.Ptr(armmanagednetworkfabric.EnabledDisabledStateEnabled),
		// 							Interval: to.Ptr[int32](4),
		// 							Multiplier: to.Ptr[int32](13),
		// 						},
		// 						FabricASN: to.Ptr[int32](65048),
		// 						Mtu: to.Ptr[int32](1500),
		// 						PeerASN: to.Ptr[int32](65047),
		// 						VlanID: to.Ptr[int32](1001),
		// 					},
		// 					OptionBProperties: &armmanagednetworkfabric.OptionBProperties{
		// 						ExportRouteTargets: []*string{
		// 							to.Ptr("65046:10039")},
		// 							ImportRouteTargets: []*string{
		// 								to.Ptr("65046:10039")},
		// 							},
		// 							PeeringOption: to.Ptr(armmanagednetworkfabric.PeeringOptionOptionA),
		// 							ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
		// 						},
		// 				}},
		// 			}
	}
}
