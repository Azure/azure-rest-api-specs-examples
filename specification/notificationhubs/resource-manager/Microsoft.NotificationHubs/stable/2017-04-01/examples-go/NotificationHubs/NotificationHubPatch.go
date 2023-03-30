package armnotificationhubs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/notificationhubs/armnotificationhubs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/NotificationHubs/NotificationHubPatch.json
func ExampleClient_Patch() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnotificationhubs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().Patch(ctx, "sdkresourceGroup", "nh-sdk-ns", "sdk-notificationHubs-8708", &armnotificationhubs.ClientPatchOptions{Parameters: &armnotificationhubs.NotificationHubPatchParameters{}})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NotificationHubResource = armnotificationhubs.NotificationHubResource{
	// 	Name: to.Ptr("nh-sdk-hub"),
	// 	Type: to.Ptr("Microsoft.NotificationHubs/namespaces/notificationHubs"),
	// 	ID: to.Ptr("/subscriptions/29cfa613-cbbc-4512-b1d6-1b3a92c7fa40/resourceGroups/sdkresourceGroup/providers/Microsoft.NotificationHubs/namespaces/nh-sdk-ns/notificationHubs/nh-sdk-hub"),
	// 	Location: to.Ptr("South Central US"),
	// 	Properties: &armnotificationhubs.NotificationHubProperties{
	// 		AuthorizationRules: []*armnotificationhubs.SharedAccessAuthorizationRuleProperties{
	// 		},
	// 		RegistrationTTL: to.Ptr("10675199.02:48:05.4775807"),
	// 	},
	// }
}
