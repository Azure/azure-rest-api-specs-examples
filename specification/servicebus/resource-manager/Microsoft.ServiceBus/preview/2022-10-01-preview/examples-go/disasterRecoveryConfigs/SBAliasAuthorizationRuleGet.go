package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5e24fa8e30bff0cb321494a0b550b1c1282a8a3c/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/disasterRecoveryConfigs/SBAliasAuthorizationRuleGet.json
func ExampleDisasterRecoveryConfigsClient_GetAuthorizationRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicebus.NewDisasterRecoveryConfigsClient("exampleSubscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetAuthorizationRule(ctx, "exampleResourceGroup", "sdk-Namespace-9080", "sdk-DisasterRecovery-4879", "sdk-Authrules-4879", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SBAuthorizationRule = armservicebus.SBAuthorizationRule{
	// 	Name: to.Ptr("sdk-Authrules-4879"),
	// 	Type: to.Ptr("Microsoft.ServiceBus/Namespaces/DisasterRecoveryConfig/AuthorizationRules"),
	// 	ID: to.Ptr("/subscriptions/exampleSubscriptionId/resourceGroups/exampleResourceGroup/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-9080/disasterRecoveryConfigs/sdk-DisasterRecovery-4047/AuthorizationRules/sdk-Authrules-4879"),
	// 	Properties: &armservicebus.SBAuthorizationRuleProperties{
	// 		Rights: []*armservicebus.AccessRights{
	// 			to.Ptr(armservicebus.AccessRightsListen),
	// 			to.Ptr(armservicebus.AccessRightsSend)},
	// 		},
	// 	}
}
