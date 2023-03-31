package armrelay_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/Relay/RelayAuthorizationRuleGet.json
func ExampleWCFRelaysClient_GetAuthorizationRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelay.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWCFRelaysClient().GetAuthorizationRule(ctx, "resourcegroup", "example-RelayNamespace-01", "example-Relay-wcf-01", "example-RelayAuthRules-01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AuthorizationRule = armrelay.AuthorizationRule{
	// 	Name: to.Ptr("example-RelayAuthRules-01"),
	// 	Type: to.Ptr("Microsoft.Relay/Namespaces/WcfRelay/AuthorizationRules"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/resourcegroup/providers/Microsoft.Relay/namespaces/example-RelayNamespace-01/WcfRelays/example-Relay-Wcf-01/AuthorizationRules/example-RelayAuthRules-01"),
	// 	Properties: &armrelay.AuthorizationRuleProperties{
	// 		Rights: []*armrelay.AccessRights{
	// 			to.Ptr(armrelay.AccessRightsListen)},
	// 		},
	// 	}
}
