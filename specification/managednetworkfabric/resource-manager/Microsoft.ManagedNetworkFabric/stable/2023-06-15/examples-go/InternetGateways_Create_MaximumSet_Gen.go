package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/925ba149e17454ce91ecd3f9f4134effb2f97844/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/InternetGateways_Create_MaximumSet_Gen.json
func ExampleInternetGatewaysClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewInternetGatewaysClient().BeginCreate(ctx, "example-rg", "example-internetGateway", armmanagednetworkfabric.InternetGateway{
		Location: to.Ptr("eastus"),
		Tags: map[string]*string{
			"key3540": to.Ptr("1234"),
		},
		Properties: &armmanagednetworkfabric.InternetGatewayProperties{
			Annotation:                to.Ptr("annotation"),
			InternetGatewayRuleID:     to.Ptr("/subscriptions/xxxx-xxxx-xxxx-xxxx/providers/Microsoft.ManagedNetworkFabric/internetGatewayRules/example-internetGatewayRule"),
			Type:                      to.Ptr(armmanagednetworkfabric.GatewayTypeInfrastructure),
			NetworkFabricControllerID: to.Ptr("/subscriptions/xxxx-xxxx-xxxx-xxxx/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabricControllers/example-networkFabricController"),
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
	// res.InternetGateway = armmanagednetworkfabric.InternetGateway{
	// 	Name: to.Ptr("example-internetGateway"),
	// 	Type: to.Ptr("microsoft.managednetworkfabric/internetGateways"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/internetGateways/example-internetGateway"),
	// 	SystemData: &armmanagednetworkfabric.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-23T14:51:13.501Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@mail.com"),
	// 		CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-23T14:51:13.501Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("email@address.com"),
	// 		LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"key3540": to.Ptr("1234"),
	// 	},
	// 	Properties: &armmanagednetworkfabric.InternetGatewayProperties{
	// 		Annotation: to.Ptr("annotation"),
	// 		InternetGatewayRuleID: to.Ptr("/subscriptions/xxxx-xxxx-xxxx-xxxx/providers/Microsoft.ManagedNetworkFabric/internetGatewayRules/example-internetGatewayRule"),
	// 		Type: to.Ptr(armmanagednetworkfabric.GatewayTypeInfrastructure),
	// 		IPv4Address: to.Ptr("10.10.10.10"),
	// 		NetworkFabricControllerID: to.Ptr("/subscriptions/xxxx-xxxx-xxxx-xxxx/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabricControllers/example-networkFabricController"),
	// 		Port: to.Ptr[int32](25),
	// 		ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
	// 	},
	// }
}
