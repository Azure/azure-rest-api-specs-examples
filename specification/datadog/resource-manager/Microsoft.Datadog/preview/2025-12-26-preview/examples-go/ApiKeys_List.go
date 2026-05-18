package armdatadog_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog/v2"
)

// Generated from example definition: 2025-12-26-preview/ApiKeys_List.json
func ExampleMonitorsClient_NewListAPIKeysPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatadog.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMonitorsClient().NewListAPIKeysPager("myResourceGroup", "myMonitor", nil)
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
		// page = armdatadog.MonitorsClientListAPIKeysResponse{
		// 	APIKeyListResponse: armdatadog.APIKeyListResponse{
		// 		Value: []*armdatadog.APIKey{
		// 			{
		// 				Name: to.Ptr("<API_KEY_NAME>"),
		// 				Created: to.Ptr("2019-04-05 09:20:30"),
		// 				CreatedBy: to.Ptr("john@example.com"),
		// 				Key: to.Ptr("1111111111111111aaaaaaaaaaaaaaaa"),
		// 			},
		// 			{
		// 				Name: to.Ptr("<API_KEY_NAME_2>"),
		// 				Created: to.Ptr("2019-04-05 09:19:53"),
		// 				CreatedBy: to.Ptr("jane@example.com"),
		// 				Key: to.Ptr("2111111111111111aaaaaaaaaaaaaaaa"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
