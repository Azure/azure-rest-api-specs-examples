package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric/v2"
)

// Generated from example definition: 2025-07-15/InternetGateways_Update.json
func ExampleInternetGatewaysClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("0000ABCD-0A0B-0000-0000-000000ABCDEF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewInternetGatewaysClient().BeginUpdate(ctx, "example-rg", "example-internetGateway", armmanagednetworkfabric.InternetGatewayPatch{
		Tags: map[string]*string{
			"keyId": to.Ptr("KeyValue"),
		},
		Properties: &armmanagednetworkfabric.InternetGatewayPatchProperties{
			InternetGatewayRuleID: to.Ptr("/subscriptions/xxxx-xxxx-xxxx-xxxx/providers/Microsoft.ManagedNetworkFabric/internetGatewayRules/example-internetGatewayRule"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmanagednetworkfabric.InternetGatewaysClientUpdateResponse{
	// 	InternetGateway: &armmanagednetworkfabric.InternetGateway{
	// 		Properties: &armmanagednetworkfabric.InternetGatewayProperties{
	// 			Annotation: to.Ptr("annotation"),
	// 			InternetGatewayRuleID: to.Ptr("/subscriptions/xxxx-xxxx-xxxx-xxxx/providers/Microsoft.ManagedNetworkFabric/internetGatewayRules/example-internetGatewayRule"),
	// 			IPv4Address: to.Ptr("10.10.10.10"),
	// 			Port: to.Ptr[int32](25),
	// 			Type: to.Ptr(armmanagednetworkfabric.GatewayTypeInfrastructure),
	// 			InternetGatewayType: to.Ptr(armmanagednetworkfabric.GatewayTypeInfrastructure),
	// 			NetworkFabricControllerID: to.Ptr("/subscriptions/xxxx-xxxx-xxxx-xxxx/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabricControllers/example-networkFabricController"),
	// 			ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateAccepted),
	// 		},
	// 		Tags: map[string]*string{
	// 			"keyId": to.Ptr("KeyValue"),
	// 		},
	// 		Location: to.Ptr("eastus"),
	// 		ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/internetGateways/example-internetGateway"),
	// 		Name: to.Ptr("example-internetGateway"),
	// 		Type: to.Ptr("microsoft.managednetworkfabric/internetGateways"),
	// 		SystemData: &armmanagednetworkfabric.SystemData{
	// 			CreatedBy: to.Ptr("email@address.com"),
	// 			CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-09T04:51:41.251Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("UserId"),
	// 			LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-09T04:51:41.251Z"); return t}()),
	// 		},
	// 	},
	// }
}
