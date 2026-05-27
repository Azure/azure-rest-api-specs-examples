package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2"
)

// Generated from example definition: 2025-05-01-preview/disasterRecoveryConfigs/SBAliasAuthorizationRuleGet.json
func ExampleDisasterRecoveryConfigsClient_GetAuthorizationRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("5f750a97-50d9-4e36-8081-c9ee4c0210d4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDisasterRecoveryConfigsClient().GetAuthorizationRule(ctx, "exampleResourceGroup", "sdk-Namespace-9080", "sdk-DisasterRecovery-4879", "sdk-Authrules-4879", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armservicebus.DisasterRecoveryConfigsClientGetAuthorizationRuleResponse{
	// 	SBAuthorizationRule: armservicebus.SBAuthorizationRule{
	// 		Name: to.Ptr("sdk-Authrules-4879"),
	// 		Type: to.Ptr("Microsoft.ServiceBus/Namespaces/DisasterRecoveryConfig/AuthorizationRules"),
	// 		ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/exampleResourceGroup/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-9080/disasterRecoveryConfigs/sdk-DisasterRecovery-4047/AuthorizationRules/sdk-Authrules-4879"),
	// 		Properties: &armservicebus.SBAuthorizationRuleProperties{
	// 			Rights: []*armservicebus.AccessRights{
	// 				to.Ptr(armservicebus.AccessRightsListen),
	// 				to.Ptr(armservicebus.AccessRightsSend),
	// 			},
	// 		},
	// 	},
	// }
}
