package armdatadog_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c280892951a9e45c059132c05aace25a9c752d48/specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/ApiKeys_List.json
func ExampleMonitorsClient_NewListAPIKeysPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatadog.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.APIKeyListResponse = armdatadog.APIKeyListResponse{
		// 	Value: []*armdatadog.APIKey{
		// 		{
		// 			Name: to.Ptr("<API_KEY_NAME>"),
		// 			Created: to.Ptr("2019-04-05 09:20:30"),
		// 			CreatedBy: to.Ptr("john@example.com"),
		// 			Key: to.Ptr("1111111111111111aaaaaaaaaaaaaaaa"),
		// 		},
		// 		{
		// 			Name: to.Ptr("<API_KEY_NAME_2>"),
		// 			Created: to.Ptr("2019-04-05 09:19:53"),
		// 			CreatedBy: to.Ptr("jane@example.com"),
		// 			Key: to.Ptr("2111111111111111aaaaaaaaaaaaaaaa"),
		// 	}},
		// }
	}
}
