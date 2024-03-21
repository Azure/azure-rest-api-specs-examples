package armnotificationhubs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/notificationhubs/armnotificationhubs/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/85cfba195a19120f309bd292c4261aa53a586adb/specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/Namespaces/AuthorizationRuleGet.json
func ExampleNamespacesClient_GetAuthorizationRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnotificationhubs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNamespacesClient().GetAuthorizationRule(ctx, "5ktrial", "nh-sdk-ns", "RootManageSharedAccessKey", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SharedAccessAuthorizationRuleResource = armnotificationhubs.SharedAccessAuthorizationRuleResource{
	// 	Name: to.Ptr("RootManageSharedAccessKey"),
	// 	Type: to.Ptr("Microsoft.NotificationHubs/namespaces/authorizationRules"),
	// 	ID: to.Ptr("/subscriptions/29cfa613-cbbc-4512-b1d6-1b3a92c7fa40/resourceGroups/5ktrial/providers/Microsoft.NotificationHubs/namespaces/nh-sdk-ns/authorizationRules/RootManageSharedAccessKey"),
	// 	Properties: &armnotificationhubs.SharedAccessAuthorizationRuleProperties{
	// 		CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-26T10:19:36.331Z"); return t}()),
	// 		ModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-26T10:19:36.331Z"); return t}()),
	// 		Rights: []*armnotificationhubs.AccessRights{
	// 			to.Ptr(armnotificationhubs.AccessRightsManage),
	// 			to.Ptr(armnotificationhubs.AccessRightsListen),
	// 			to.Ptr(armnotificationhubs.AccessRightsSend)},
	// 		},
	// 	}
}
