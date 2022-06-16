package armmanagementgroups_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementgroups/armmanagementgroups"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/GetSubscriptionFromManagementGroup.json
func ExampleManagementGroupSubscriptionsClient_GetSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmanagementgroups.NewManagementGroupSubscriptionsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetSubscription(ctx,
		"Group",
		"728bcbe4-8d56-4510-86c2-4921b8beefbc",
		&armmanagementgroups.ManagementGroupSubscriptionsClientGetSubscriptionOptions{CacheControl: to.Ptr("no-cache")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
