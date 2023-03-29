package armlocks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armlocks"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/resources/resource-manager/Microsoft.Authorization/stable/2020-05-01/examples/ManagementLocks_ListAtSubscriptionLevel.json
func ExampleManagementLocksClient_NewListAtSubscriptionLevelPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlocks.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagementLocksClient().NewListAtSubscriptionLevelPager(&armlocks.ManagementLocksClientListAtSubscriptionLevelOptions{Filter: nil})
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
		// page.ManagementLockListResult = armlocks.ManagementLockListResult{
		// 	Value: []*armlocks.ManagementLockObject{
		// 		{
		// 			Name: to.Ptr("testlock"),
		// 			Type: to.Ptr("Microsoft.Authorization/locks"),
		// 			ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourcegroupname/providers/Microsoft.Authorization/locks/testlock"),
		// 			Properties: &armlocks.ManagementLockProperties{
		// 				Level: to.Ptr(armlocks.LockLevelReadOnly),
		// 			},
		// 	}},
		// }
	}
}
