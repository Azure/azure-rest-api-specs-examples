package armnotificationhubs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/notificationhubs/armnotificationhubs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/Namespaces/NHNameSpaceAuthorizationRuleCreate.json
func ExampleNamespacesClient_CreateOrUpdateAuthorizationRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnotificationhubs.NewNamespacesClient("29cfa613-cbbc-4512-b1d6-1b3a92c7fa40", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdateAuthorizationRule(ctx,
		"5ktrial",
		"nh-sdk-ns",
		"sdk-AuthRules-1788",
		armnotificationhubs.SharedAccessAuthorizationRuleCreateOrUpdateParameters{
			Properties: &armnotificationhubs.SharedAccessAuthorizationRuleProperties{
				Rights: []*armnotificationhubs.AccessRights{
					to.Ptr(armnotificationhubs.AccessRightsListen),
					to.Ptr(armnotificationhubs.AccessRightsSend)},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
