package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-01-01-preview/examples/NameSpaces/SBNameSpaceAuthorizationRuleCreate.json
func ExampleNamespacesClient_CreateOrUpdateAuthorizationRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicebus.NewNamespacesClient("5f750a97-50d9-4e36-8081-c9ee4c0210d4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdateAuthorizationRule(ctx,
		"ArunMonocle",
		"sdk-Namespace-6914",
		"sdk-AuthRules-1788",
		armservicebus.SBAuthorizationRule{
			Properties: &armservicebus.SBAuthorizationRuleProperties{
				Rights: []*armservicebus.AccessRights{
					to.Ptr(armservicebus.AccessRightsListen),
					to.Ptr(armservicebus.AccessRightsSend)},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
