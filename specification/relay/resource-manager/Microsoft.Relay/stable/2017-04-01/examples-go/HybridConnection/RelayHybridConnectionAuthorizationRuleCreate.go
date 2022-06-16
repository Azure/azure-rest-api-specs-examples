package armrelay_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/HybridConnection/RelayHybridConnectionAuthorizationRuleCreate.json
func ExampleHybridConnectionsClient_CreateOrUpdateAuthorizationRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armrelay.NewHybridConnectionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdateAuthorizationRule(ctx,
		"<resource-group-name>",
		"<namespace-name>",
		"<hybrid-connection-name>",
		"<authorization-rule-name>",
		armrelay.AuthorizationRule{
			Properties: &armrelay.AuthorizationRuleProperties{
				Rights: []*armrelay.AccessRights{
					to.Ptr(armrelay.AccessRightsListen),
					to.Ptr(armrelay.AccessRightsSend)},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
