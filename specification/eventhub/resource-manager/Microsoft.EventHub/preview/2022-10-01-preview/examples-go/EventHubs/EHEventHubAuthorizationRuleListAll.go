package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-10-01-preview/examples/EventHubs/EHEventHubAuthorizationRuleListAll.json
func ExampleEventHubsClient_NewListAuthorizationRulesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEventHubsClient().NewListAuthorizationRulesPager("ArunMonocle", "sdk-Namespace-960", "sdk-EventHub-532", nil)
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
		// 			Name: to.Ptr("sdk-Authrules-2513"),
		// 			Type: to.Ptr("Microsoft.EventHub/Namespaces/EventHubs/AuthorizationRules"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/ArunMonocle/providers/Microsoft.EventHub/namespaces/sdk-Namespace-960/eventhubs/sdk-EventHub-532/authorizationRules/sdk-Authrules-2513"),
		// 			Properties: &armeventhub.AuthorizationRuleProperties{
		// 				Rights: []*armeventhub.AccessRights{
		// 					to.Ptr(armeventhub.AccessRightsListen),
		// 					to.Ptr(armeventhub.AccessRightsSend)},
		// 				},
		// 		}},
		// 	}
	}
}
