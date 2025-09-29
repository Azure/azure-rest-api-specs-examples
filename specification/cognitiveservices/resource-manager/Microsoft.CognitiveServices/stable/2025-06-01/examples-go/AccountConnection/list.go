package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1004eed4202d64b48157c084fe2830760f8190f4/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/AccountConnection/list.json
func ExampleAccountConnectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountConnectionsClient().NewListPager("resourceGroup-1", "account-1", &armcognitiveservices.AccountConnectionsClientListOptions{Target: to.Ptr("[tartget url]"),
		Category:   to.Ptr("ContainerRegistry"),
		IncludeAll: nil,
	})
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
		// page.ConnectionPropertiesV2BasicResourceArmPaginatedResult = armcognitiveservices.ConnectionPropertiesV2BasicResourceArmPaginatedResult{
		// 	Value: []*armcognitiveservices.ConnectionPropertiesV2BasicResource{
		// 		{
		// 			Name: to.Ptr("connection-1"),
		// 			Type: to.Ptr("Microsoft.CognitiveServices/accounts/connections"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup-1/providers/Microsoft.CognitiveServices/accounts/account-1/linkedaccounts/connection-1"),
		// 			Properties: &armcognitiveservices.PATAuthTypeConnectionProperties{
		// 				AuthType: to.Ptr(armcognitiveservices.ConnectionAuthTypePAT),
		// 				Category: to.Ptr(armcognitiveservices.ConnectionCategoryContainerRegistry),
		// 				ExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-15T14:30:00.000Z"); return t}()),
		// 				Target: to.Ptr("[tartget url]"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("connection-2"),
		// 			Type: to.Ptr("Microsoft.CognitiveServices/accounts/connections"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup-1/providers/Microsoft.CognitiveServices/accounts/account-1/linkedaccounts/connection-2"),
		// 			Properties: &armcognitiveservices.PATAuthTypeConnectionProperties{
		// 				AuthType: to.Ptr(armcognitiveservices.ConnectionAuthTypePAT),
		// 				Category: to.Ptr(armcognitiveservices.ConnectionCategoryContainerRegistry),
		// 				Target: to.Ptr("[tartget url]"),
		// 			},
		// 	}},
		// }
	}
}
