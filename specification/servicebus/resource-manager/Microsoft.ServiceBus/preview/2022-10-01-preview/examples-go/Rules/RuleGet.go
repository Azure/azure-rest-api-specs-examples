package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5e24fa8e30bff0cb321494a0b550b1c1282a8a3c/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/Rules/RuleGet.json
func ExampleRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicebus.NewRulesClient("5f750a97-50d9-4e36-8081-c9ee4c0210d4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "ArunMonocle", "sdk-Namespace-1319", "sdk-Topics-2081", "sdk-Subscriptions-8691", "sdk-Rules-6571", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Rule = armservicebus.Rule{
	// 	Name: to.Ptr("sdk-Rules-6571"),
	// 	Type: to.Ptr("Microsoft.ServiceBus/Namespaces/Topics/Subscriptions/Rules"),
	// 	ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/ArunMonocle/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-1319/topics/sdk-Topics-2081/subscriptions/sdk-Subscriptions-8691/rules/sdk-Rules-6571"),
	// 	Properties: &armservicebus.Ruleproperties{
	// 		Action: &armservicebus.Action{
	// 		},
	// 		FilterType: to.Ptr(armservicebus.FilterTypeSQLFilter),
	// 		SQLFilter: &armservicebus.SQLFilter{
	// 			CompatibilityLevel: to.Ptr[int32](20),
	// 			SQLExpression: to.Ptr("1=1"),
	// 		},
	// 	},
	// }
}
