package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/925ba149e17454ce91ecd3f9f4134effb2f97844/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/InternetGatewayRules_ListBySubscription_MaximumSet_Gen.json
func ExampleInternetGatewayRulesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewInternetGatewayRulesClient().NewListBySubscriptionPager(nil)
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
		// page.InternetGatewayRulesListResult = armmanagednetworkfabric.InternetGatewayRulesListResult{
		// 	Value: []*armmanagednetworkfabric.InternetGatewayRule{
		// 		{
		// 			Name: to.Ptr("example-internetGatewayRule"),
		// 			Type: to.Ptr("microsoft.managednetworkfabric/internetGatewayRules"),
		// 			ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/internetGatewayRules/example-InternetGatewayRule"),
		// 			SystemData: &armmanagednetworkfabric.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-24T18:22:51.443Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@mail.com"),
		// 				CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-05-24T18:22:51.443Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@mail.com"),
		// 				LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"keyID": to.Ptr("keyValue"),
		// 			},
		// 			Properties: &armmanagednetworkfabric.InternetGatewayRuleProperties{
		// 				Annotation: to.Ptr("annotationValue"),
		// 				InternetGatewayIDs: []*string{
		// 					to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/internetGateways/example-internetGateway")},
		// 					ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
		// 					RuleProperties: &armmanagednetworkfabric.RuleProperties{
		// 						Action: to.Ptr(armmanagednetworkfabric.ActionAllow),
		// 						AddressList: []*string{
		// 							to.Ptr("10.10.10.10")},
		// 						},
		// 					},
		// 			}},
		// 		}
	}
}
