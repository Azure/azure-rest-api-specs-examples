package armrelay_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay/v2"
)

// Generated from example definition: 2024-01-01/Relay/RelayAuthorizationRuleCreate.json
func ExampleWCFRelaysClient_CreateOrUpdateAuthorizationRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelay.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWCFRelaysClient().CreateOrUpdateAuthorizationRule(ctx, "resourcegroup", "example-RelayNamespace-01", "example-Relay-wcf-01", "example-RelayAuthRules-01", armrelay.AuthorizationRule{
		Properties: &armrelay.AuthorizationRuleProperties{
			Rights: []*armrelay.AccessRights{
				to.Ptr(armrelay.AccessRightsListen),
				to.Ptr(armrelay.AccessRightsSend),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armrelay.WCFRelaysClientCreateOrUpdateAuthorizationRuleResponse{
	// 	AuthorizationRule: &armrelay.AuthorizationRule{
	// 		Name: to.Ptr("example-RelayAuthRules-01"),
	// 		Type: to.Ptr("Microsoft.Relay/Namespaces/WcfRelay/AuthorizationRules"),
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/resourcegroup/providers/Microsoft.Relay/namespaces/example-RelayNamespace-01/WcfRelays/example-Relay-Wcf-01/AuthorizationRules/example-RelayAuthRules-01"),
	// 		Properties: &armrelay.AuthorizationRuleProperties{
	// 			Rights: []*armrelay.AccessRights{
	// 				to.Ptr(armrelay.AccessRightsListen),
	// 				to.Ptr(armrelay.AccessRightsSend),
	// 			},
	// 		},
	// 	},
	// }
}
