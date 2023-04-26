package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-10-01-preview/examples/disasterRecoveryConfigs/EHAliasAuthorizationRuleListAll.json
func ExampleDisasterRecoveryConfigsClient_NewListAuthorizationRulesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.AuthorizationRuleListResult = armeventhub.AuthorizationRuleListResult{
		// 	Value: []*armeventhub.AuthorizationRule{
		// 		{
		// 			Name: to.Ptr("RootManageSharedAccessKey"),
		// 			Type: to.Ptr("Microsoft.EventHub/Namespaces/AuthorizationRules"),
		// 			ID: to.Ptr("/subscriptions/exampleSubscriptionId/resourceGroups/exampleResourceGroup/providers/Microsoft.EventHub/namespaces/sdk-Namespace-9080/disasterRecoveryConfigs/sdk-DisasterRecovery-4047/AuthorizationRules/RootManageSharedAccessKey"),
		// 			Properties: &armeventhub.AuthorizationRuleProperties{
		// 				Rights: []*armeventhub.AccessRights{
		// 					to.Ptr(armeventhub.AccessRightsListen),
		// 					to.Ptr(armeventhub.AccessRightsManage),
		// 					to.Ptr(armeventhub.AccessRightsSend)},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("sdk-Authrules-1067"),
		// 				Type: to.Ptr("Microsoft.EventHub/Namespaces/AuthorizationRules"),
		// 				ID: to.Ptr("/subscriptions/exampleSubscriptionId/resourceGroups/exampleResourceGroup/providers/Microsoft.EventHub/namespaces/sdk-Namespace-9080/disasterRecoveryConfigs/sdk-DisasterRecovery-4047/AuthorizationRules/sdk-Authrules-1067"),
		// 				Properties: &armeventhub.AuthorizationRuleProperties{
		// 					Rights: []*armeventhub.AccessRights{
		// 						to.Ptr(armeventhub.AccessRightsListen),
		// 						to.Ptr(armeventhub.AccessRightsSend)},
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("sdk-Authrules-1684"),
		// 					Type: to.Ptr("Microsoft.EventHub/Namespaces/AuthorizationRules"),
		// 					ID: to.Ptr("/subscriptions/exampleSubscriptionId/resourceGroups/exampleResourceGroup/providers/Microsoft.EventHub/namespaces/sdk-Namespace-9080/disasterRecoveryConfigs/sdk-DisasterRecovery-4047/AuthorizationRules/sdk-Authrules-1684"),
		// 					Properties: &armeventhub.AuthorizationRuleProperties{
		// 						Rights: []*armeventhub.AccessRights{
		// 							to.Ptr(armeventhub.AccessRightsListen),
		// 							to.Ptr(armeventhub.AccessRightsSend)},
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("sdk-Authrules-4879"),
		// 						Type: to.Ptr("Microsoft.EventHub/Namespaces/AuthorizationRules"),
		// 						ID: to.Ptr("/subscriptions/exampleSubscriptionId/resourceGroups/exampleResourceGroup/providers/Microsoft.EventHub/namespaces/sdk-Namespace-9080/disasterRecoveryConfigs/sdk-DisasterRecovery-4047/AuthorizationRules/sdk-Authrules-4879"),
		// 						Properties: &armeventhub.AuthorizationRuleProperties{
		// 							Rights: []*armeventhub.AccessRights{
		// 								to.Ptr(armeventhub.AccessRightsListen),
		// 								to.Ptr(armeventhub.AccessRightsSend)},
		// 							},
		// 					}},
		// 				}
	}
}
