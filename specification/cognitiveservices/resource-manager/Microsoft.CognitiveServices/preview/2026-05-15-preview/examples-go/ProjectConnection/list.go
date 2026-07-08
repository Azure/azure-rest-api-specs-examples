package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-05-15-preview/ProjectConnection/list.json
func ExampleProjectConnectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProjectConnectionsClient().NewListPager("resourceGroup-1", "account-1", "project-1", &armcognitiveservices.ProjectConnectionsClientListOptions{
		Category: to.Ptr("ContainerRegistry"),
		Target:   to.Ptr("[target url]")})
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
		// page = armcognitiveservices.ProjectConnectionsClientListResponse{
		// 	ConnectionPropertiesV2BasicResourceArmPaginatedResult: armcognitiveservices.ConnectionPropertiesV2BasicResourceArmPaginatedResult{
		// 		Value: []*armcognitiveservices.ConnectionPropertiesV2BasicResource{
		// 			{
		// 				Name: to.Ptr("connection-1"),
		// 				Type: to.Ptr("Microsoft.CognitiveServices/accounts/projects/connections"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup-1/providers/Microsoft.CognitiveServices/accounts/account-1/projects/project-1/connections/connection-1"),
		// 				Properties: &armcognitiveservices.PATAuthTypeConnectionProperties{
		// 					AuthType: to.Ptr(armcognitiveservices.ConnectionAuthTypePAT),
		// 					Category: to.Ptr(armcognitiveservices.ConnectionCategoryContainerRegistry),
		// 					ExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-15T14:30:00Z"); return t}()),
		// 					Target: to.Ptr("[target url]"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("connection-2"),
		// 				Type: to.Ptr("Microsoft.CognitiveServices/accounts/projects/connections"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup-1/providers/Microsoft.CognitiveServices/accounts/account-1/projects/project-1/connections/connection-2"),
		// 				Properties: &armcognitiveservices.PATAuthTypeConnectionProperties{
		// 					AuthType: to.Ptr(armcognitiveservices.ConnectionAuthTypePAT),
		// 					Category: to.Ptr(armcognitiveservices.ConnectionCategoryContainerRegistry),
		// 					Target: to.Ptr("[target url]"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
