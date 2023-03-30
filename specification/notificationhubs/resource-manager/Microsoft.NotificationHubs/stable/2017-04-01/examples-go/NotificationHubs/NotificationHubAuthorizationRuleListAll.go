package armnotificationhubs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/notificationhubs/armnotificationhubs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/NotificationHubs/NotificationHubAuthorizationRuleListAll.json
func ExampleClient_NewListAuthorizationRulesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnotificationhubs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListAuthorizationRulesPager("5ktrial", "nh-sdk-ns", "nh-sdk-hub", nil)
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
		// page.SharedAccessAuthorizationRuleListResult = armnotificationhubs.SharedAccessAuthorizationRuleListResult{
		// 	Value: []*armnotificationhubs.SharedAccessAuthorizationRuleResource{
		// 		{
		// 			Name: to.Ptr("DefaultListenSharedAccessSignature"),
		// 			Type: to.Ptr("Microsoft.NotificationHubs/namespaces/notificationHubs/authorizationRules"),
		// 			ID: to.Ptr("/subscriptions/29cfa613-cbbc-4512-b1d6-1b3a92c7fa40/resourceGroups/5ktrial/providers/Microsoft.NotificationHubs/namespaces/nh-sdk-ns/NotificationHubs/nh-sdk-hub/AuthorizationRules/DefaultListenSharedAccessSignature"),
		// 			Location: to.Ptr("West Europe"),
		// 			Properties: &armnotificationhubs.SharedAccessAuthorizationRuleProperties{
		// 				ClaimType: to.Ptr("SharedAccessKey"),
		// 				ClaimValue: to.Ptr("None"),
		// 				CreatedTime: to.Ptr("2018-05-02T00:45:22.0150024Z"),
		// 				KeyName: to.Ptr("DefaultListenSharedAccessSignature"),
		// 				ModifiedTime: to.Ptr("2018-05-02T00:45:22.0150024Z"),
		// 				PrimaryKey: to.Ptr("#################################"),
		// 				Rights: []*armnotificationhubs.AccessRights{
		// 					to.Ptr(armnotificationhubs.AccessRightsListen)},
		// 					SecondaryKey: to.Ptr("#################################"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("DefaultFullSharedAccessSignature"),
		// 				Type: to.Ptr("Microsoft.NotificationHubs/namespaces/notificationHubs/authorizationRules"),
		// 				ID: to.Ptr("/subscriptions/29cfa613-cbbc-4512-b1d6-1b3a92c7fa40/resourceGroups/5ktrial/providers/Microsoft.NotificationHubs/namespaces/nh-sdk-ns/NotificationHubs/nh-sdk-hub/AuthorizationRules/DefaultFullSharedAccessSignature"),
		// 				Location: to.Ptr("West Europe"),
		// 				Properties: &armnotificationhubs.SharedAccessAuthorizationRuleProperties{
		// 					ClaimType: to.Ptr("SharedAccessKey"),
		// 					ClaimValue: to.Ptr("None"),
		// 					CreatedTime: to.Ptr("2018-05-02T00:45:22.0150024Z"),
		// 					KeyName: to.Ptr("DefaultFullSharedAccessSignature"),
		// 					ModifiedTime: to.Ptr("2018-05-02T00:45:22.0150024Z"),
		// 					PrimaryKey: to.Ptr("#################################"),
		// 					Rights: []*armnotificationhubs.AccessRights{
		// 						to.Ptr(armnotificationhubs.AccessRightsListen),
		// 						to.Ptr(armnotificationhubs.AccessRightsManage),
		// 						to.Ptr(armnotificationhubs.AccessRightsSend)},
		// 						SecondaryKey: to.Ptr("#################################"),
		// 					},
		// 			}},
		// 		}
	}
}
