package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/NameSpaces/VirtualNetworkRule/SBNetworkRuleSetGet.json
func ExampleNamespacesClient_GetNetworkRuleSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicebus.NewNamespacesClient("Subscription", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetNetworkRuleSet(ctx,
		"ResourceGroup",
		"sdk-Namespace-6019",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
