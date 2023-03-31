package armnotificationhubs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/notificationhubs/armnotificationhubs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/Namespaces/NHNameSpaceAuthorizationRuleGet.json
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
	// 	Type: to.Ptr("Microsoft.NotificationHubs/Namespaces/AuthorizationRules"),
	// 	ID: to.Ptr("/subscriptions/29cfa613-cbbc-4512-b1d6-1b3a92c7fa40/resourceGroups/5ktrial/providers/Microsoft.NotificationHubs/namespaces/nh-sdk-ns/AuthorizationRules/RootManageSharedAccessKey"),
	// 	Location: to.Ptr("South Central US"),
	// 	Properties: &armnotificationhubs.SharedAccessAuthorizationRuleProperties{
	// 		ClaimType: to.Ptr("SharedAccessKey"),
	// 		ClaimValue: to.Ptr("None"),
	// 		CreatedTime: to.Ptr("2018-05-02T18:24:51.0690674Z"),
	// 		KeyName: to.Ptr("RootManageSharedAccessKey"),
	// 		ModifiedTime: to.Ptr("2018-05-02T18:24:51.0690674Z"),
	// 		PrimaryKey: to.Ptr("############################################"),
	// 		Rights: []*armnotificationhubs.AccessRights{
	// 			to.Ptr(armnotificationhubs.AccessRightsListen),
	// 			to.Ptr(armnotificationhubs.AccessRightsManage),
	// 			to.Ptr(armnotificationhubs.AccessRightsSend)},
	// 			SecondaryKey: to.Ptr("############################################"),
	// 		},
	// 	}
}
