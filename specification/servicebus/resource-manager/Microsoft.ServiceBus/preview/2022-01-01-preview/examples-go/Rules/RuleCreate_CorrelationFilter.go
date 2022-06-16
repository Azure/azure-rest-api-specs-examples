package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-01-01-preview/examples/Rules/RuleCreate_CorrelationFilter.json
func ExampleRulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicebus.NewRulesClient("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"resourceGroupName",
		"sdk-Namespace-1319",
		"sdk-Topics-2081",
		"sdk-Subscriptions-8691",
		"sdk-Rules-6571",
		armservicebus.Rule{
			Properties: &armservicebus.Ruleproperties{
				CorrelationFilter: &armservicebus.CorrelationFilter{
					Properties: map[string]*string{
						"topicHint": to.Ptr("Crop"),
					},
				},
				FilterType: to.Ptr(armservicebus.FilterTypeCorrelationFilter),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
