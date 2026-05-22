package armnotificationhubs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/notificationhubs/armnotificationhubs/v2"
)

// Generated from example definition: 2023-10-01-preview/NotificationHubs/AuthorizationRuleList.json
func ExampleClient_NewListAuthorizationRulesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnotificationhubs.NewClientFactory("29cfa613-cbbc-4512-b1d6-1b3a92c7fa40", cred, nil)
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
		// page = armnotificationhubs.ClientListAuthorizationRulesResponse{
		// 	SharedAccessAuthorizationRuleListResult: armnotificationhubs.SharedAccessAuthorizationRuleListResult{
		// 		Value: []*armnotificationhubs.SharedAccessAuthorizationRuleResource{
		// 			{
		// 				Name: to.Ptr("DefaultListenSharedAccessSignature"),
		// 				Type: to.Ptr("Microsoft.NotificationHubs/namespaces/notificationHubs/authorizationRules"),
		// 				ID: to.Ptr("/subscriptions/29cfa613-cbbc-4512-b1d6-1b3a92c7fa40/resourceGroups/5ktrial/providers/Microsoft.NotificationHubs/namespaces/nh-sdk-ns/notificationHubs/test/authorizationRules/DefaultListenSharedAccessSignature"),
		// 				Properties: &armnotificationhubs.SharedAccessAuthorizationRuleProperties{
		// 					CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-26T10:43:00.5320365+00:00"); return t}()),
		// 					ModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-26T10:43:00.5320368+00:00"); return t}()),
		// 					Rights: []*armnotificationhubs.AccessRights{
		// 						to.Ptr(armnotificationhubs.AccessRightsListen),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("DefaultFullSharedAccessSignature"),
		// 				Type: to.Ptr("Microsoft.NotificationHubs/namespaces/notificationHubs/authorizationRules"),
		// 				ID: to.Ptr("/subscriptions/29cfa613-cbbc-4512-b1d6-1b3a92c7fa40/resourceGroups/5ktrial/providers/Microsoft.NotificationHubs/namespaces/nh-sdk-ns/notificationHubs/test/authorizationRules/DefaultFullSharedAccessSignature"),
		// 				Properties: &armnotificationhubs.SharedAccessAuthorizationRuleProperties{
		// 					CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-26T10:43:00.5320386+00:00"); return t}()),
		// 					ModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-26T10:43:00.5320387+00:00"); return t}()),
		// 					Rights: []*armnotificationhubs.AccessRights{
		// 						to.Ptr(armnotificationhubs.AccessRightsManage),
		// 						to.Ptr(armnotificationhubs.AccessRightsListen),
		// 						to.Ptr(armnotificationhubs.AccessRightsSend),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
