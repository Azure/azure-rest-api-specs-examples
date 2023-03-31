package armrelay_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/HybridConnection/RelayHybridConnectionAuthorizationRuleListAll.json
func ExampleHybridConnectionsClient_NewListAuthorizationRulesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelay.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewHybridConnectionsClient().NewListAuthorizationRulesPager("resourcegroup", "example-RelayNamespace-01", "example-Relay-Hybrid-01", nil)
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
		// page.AuthorizationRuleListResult = armrelay.AuthorizationRuleListResult{
		// 	Value: []*armrelay.AuthorizationRule{
		// 		{
		// 			Name: to.Ptr("example-RelayAuthRules-01"),
		// 			Type: to.Ptr("Microsoft.Relay/Namespaces/Hybridconnections/AuthorizationRules"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/resourcegroup/providers/Microsoft.Relay/namespaces/example-RelayNamespace-01/HybridConnections/example-Relay-Hybrid-01/AuthorizationRules/example-RelayAuthRules-01"),
		// 			Properties: &armrelay.AuthorizationRuleProperties{
		// 				Rights: []*armrelay.AccessRights{
		// 					to.Ptr(armrelay.AccessRightsListen),
		// 					to.Ptr(armrelay.AccessRightsSend)},
		// 				},
		// 		}},
		// 	}
	}
}
