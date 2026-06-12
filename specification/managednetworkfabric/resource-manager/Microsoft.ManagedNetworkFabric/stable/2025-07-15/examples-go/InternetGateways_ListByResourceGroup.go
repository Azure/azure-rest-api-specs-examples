package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/InternetGateways_ListByResourceGroup.json
func ExampleInternetGatewaysClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("0000ABCD-0A0B-0000-0000-000000ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewInternetGatewaysClient().NewListByResourceGroupPager("example-rg", nil)
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
		// page = armmanagednetworkfabric.InternetGatewaysClientListByResourceGroupResponse{
		// 	InternetGatewaysListResult: armmanagednetworkfabric.InternetGatewaysListResult{
		// 		Value: []*armmanagednetworkfabric.InternetGateway{
		// 			{
		// 				Properties: &armmanagednetworkfabric.InternetGatewayProperties{
		// 					Annotation: to.Ptr("annotation"),
		// 					InternetGatewayRuleID: to.Ptr("/subscriptions/xxxx-xxxx-xxxx-xxxx/providers/Microsoft.ManagedNetworkFabric/internetGatewayRules/example-internetGatewayRule"),
		// 					IPv4Address: to.Ptr("10.10.10.10"),
		// 					Port: to.Ptr[int32](25),
		// 					Type: to.Ptr(armmanagednetworkfabric.GatewayTypeInfrastructure),
		// 					InternetGatewayType: to.Ptr(armmanagednetworkfabric.GatewayTypeInfrastructure),
		// 					NetworkFabricControllerID: to.Ptr("/subscriptions/xxxx-xxxx-xxxx-xxxx/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabricControllers/example-networkFabricController"),
		// 					LastOperation: &armmanagednetworkfabric.LastOperationProperties{
		// 						Details: to.Ptr("Succeeded"),
		// 					},
		// 					ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateAccepted),
		// 				},
		// 				Tags: map[string]*string{
		// 					"keyId": to.Ptr("KeyValue"),
		// 				},
		// 				Location: to.Ptr("eastus"),
		// 				ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/internetGateways/example-internetGateway"),
		// 				Name: to.Ptr("example-internetGateway"),
		// 				Type: to.Ptr("microsoft.managednetworkfabric/internetGateways"),
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
