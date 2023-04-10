package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Rules/RuleListBySubscription.json
func ExampleRulesClient_NewListBySubscriptionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRulesClient().NewListBySubscriptionsPager("ArunMonocle", "sdk-Namespace-1319", "sdk-Topics-2081", "sdk-Subscriptions-8691", &armservicebus.RulesClientListBySubscriptionsOptions{Skip: nil,
		Top: nil,
	})
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
		// page.RuleListResult = armservicebus.RuleListResult{
		// 	Value: []*armservicebus.Rule{
		// 		{
		// 			Name: to.Ptr("sdk-Rules-6571"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces/Topics/Subscriptions/Rules"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/ArunMonocle/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-1319/topics/sdk-Topics-2081/subscriptions/sdk-Subscriptions-8691/rules/sdk-Rules-6571"),
		// 			Properties: &armservicebus.Ruleproperties{
		// 				Action: &armservicebus.Action{
		// 				},
		// 				FilterType: to.Ptr(armservicebus.FilterTypeSQLFilter),
		// 				SQLFilter: &armservicebus.SQLFilter{
		// 					CompatibilityLevel: to.Ptr[int32](20),
		// 					SQLExpression: to.Ptr("1=1"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
