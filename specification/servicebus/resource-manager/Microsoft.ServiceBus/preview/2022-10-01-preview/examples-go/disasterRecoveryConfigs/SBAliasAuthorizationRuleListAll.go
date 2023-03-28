package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5e24fa8e30bff0cb321494a0b550b1c1282a8a3c/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/disasterRecoveryConfigs/SBAliasAuthorizationRuleListAll.json
func ExampleDisasterRecoveryConfigsClient_NewListAuthorizationRulesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDisasterRecoveryConfigsClient().NewListAuthorizationRulesPager("exampleResourceGroup", "sdk-Namespace-9080", "sdk-DisasterRecovery-4047", nil)
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
		// page.SBAuthorizationRuleListResult = armservicebus.SBAuthorizationRuleListResult{
		// 	Value: []*armservicebus.SBAuthorizationRule{
		// 		{
		// 			Name: to.Ptr("RootManageSharedAccessKey"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces/DisasterRecoveryConfig/AuthorizationRules"),
		// 			ID: to.Ptr("/subscriptions/exampleSubscriptionId/resourceGroups/exampleResourceGroup/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-9080/disasterRecoveryConfigs/sdk-DisasterRecovery-4047/AuthorizationRules/RootManageSharedAccessKey"),
		// 			Properties: &armservicebus.SBAuthorizationRuleProperties{
		// 				Rights: []*armservicebus.AccessRights{
		// 					to.Ptr(armservicebus.AccessRightsListen),
		// 					to.Ptr(armservicebus.AccessRightsManage),
		// 					to.Ptr(armservicebus.AccessRightsSend)},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("sdk-Authrules-1067"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces/DisasterRecoveryConfig/AuthorizationRules"),
		// 				ID: to.Ptr("/subscriptions/exampleSubscriptionId/resourceGroups/exampleResourceGroup/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-9080/disasterRecoveryConfigs/sdk-DisasterRecovery-4047/AuthorizationRules/sdk-Authrules-1067"),
		// 				Properties: &armservicebus.SBAuthorizationRuleProperties{
		// 					Rights: []*armservicebus.AccessRights{
		// 						to.Ptr(armservicebus.AccessRightsListen),
		// 						to.Ptr(armservicebus.AccessRightsSend)},
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("sdk-Authrules-1684"),
		// 					Type: to.Ptr("Microsoft.ServiceBus/Namespaces/DisasterRecoveryConfig/AuthorizationRules"),
		// 					ID: to.Ptr("/subscriptions/exampleSubscriptionId/resourceGroups/exampleResourceGroup/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-9080/disasterRecoveryConfigs/sdk-DisasterRecovery-4047/AuthorizationRules/sdk-Authrules-1684"),
		// 					Properties: &armservicebus.SBAuthorizationRuleProperties{
		// 						Rights: []*armservicebus.AccessRights{
		// 							to.Ptr(armservicebus.AccessRightsListen),
		// 							to.Ptr(armservicebus.AccessRightsSend)},
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("sdk-Authrules-4879"),
		// 						Type: to.Ptr("Microsoft.ServiceBus/Namespaces/DisasterRecoveryConfig/AuthorizationRules"),
		// 						ID: to.Ptr("/subscriptions/exampleSubscriptionId/resourceGroups/exampleResourceGroup/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-9080/disasterRecoveryConfigs/sdk-DisasterRecovery-4047/AuthorizationRules/sdk-Authrules-4879"),
		// 						Properties: &armservicebus.SBAuthorizationRuleProperties{
		// 							Rights: []*armservicebus.AccessRights{
		// 								to.Ptr(armservicebus.AccessRightsListen),
		// 								to.Ptr(armservicebus.AccessRightsSend)},
		// 							},
		// 					}},
		// 				}
	}
}
