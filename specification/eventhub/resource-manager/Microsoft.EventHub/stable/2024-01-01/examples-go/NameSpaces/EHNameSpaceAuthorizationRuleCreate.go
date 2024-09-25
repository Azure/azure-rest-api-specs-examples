package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b574e2a41acda14a90ef237006e8bbdda2b63c63/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/NameSpaces/EHNameSpaceAuthorizationRuleCreate.json
func ExampleNamespacesClient_CreateOrUpdateAuthorizationRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNamespacesClient().CreateOrUpdateAuthorizationRule(ctx, "ArunMonocle", "sdk-Namespace-2702", "sdk-Authrules-1746", armeventhub.AuthorizationRule{
		Properties: &armeventhub.AuthorizationRuleProperties{
			Rights: []*armeventhub.AccessRights{
				to.Ptr(armeventhub.AccessRightsListen),
				to.Ptr(armeventhub.AccessRightsSend)},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AuthorizationRule = armeventhub.AuthorizationRule{
	// 	Name: to.Ptr("sdk-Authrules-1746"),
	// 	Type: to.Ptr("Microsoft.EventHub/Namespaces/AuthorizationRules"),
	// 	ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/ArunMonocle/providers/Microsoft.EventHub/namespaces/sdk-Namespace-2702/AuthorizationRules/sdk-Authrules-1746"),
	// 	Properties: &armeventhub.AuthorizationRuleProperties{
	// 		Rights: []*armeventhub.AccessRights{
	// 			to.Ptr(armeventhub.AccessRightsListen),
	// 			to.Ptr(armeventhub.AccessRightsSend)},
	// 		},
	// 	}
}
