package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountSessions_List.json
func ExampleIntegrationAccountSessionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIntegrationAccountSessionsClient().NewListPager("testrg123", "testia123", &armlogic.IntegrationAccountSessionsClientListOptions{Top: nil,
		Filter: nil,
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
		// page.IntegrationAccountSessionListResult = armlogic.IntegrationAccountSessionListResult{
		// 	Value: []*armlogic.IntegrationAccountSession{
		// 		{
		// 			Name: to.Ptr("IntegrationAccountSession1662"),
		// 			Type: to.Ptr("Microsoft.Logic/integrationAccounts/sessions"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Logic/integrationAccounts/testia123/sessions/IntegrationAccountSession1662"),
		// 			Properties: &armlogic.IntegrationAccountSessionProperties{
		// 				ChangedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-21T23:06:24.015Z"); return t}()),
		// 				Content: map[string]any{
		// 					"controlNumber": "1234",
		// 					"controlNumberChangedTime": "2017-02-21T22:30:11.9923759Z",
		// 				},
		// 				CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-21T23:06:24.015Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("IntegrationAccountSession6808"),
		// 			Type: to.Ptr("Microsoft.Logic/integrationAccounts/sessions"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Logic/integrationAccounts/testia123/sessions/IntegrationAccountSession6808"),
		// 			Properties: &armlogic.IntegrationAccountSessionProperties{
		// 				ChangedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-21T23:06:23.754Z"); return t}()),
		// 				Content: map[string]any{
		// 					"controlNumber": "1234",
		// 					"controlNumberChangedTime": "2017-02-21T22:30:11.9923759Z",
		// 				},
		// 				CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-21T23:06:23.753Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("IntegrationAccountSession7315"),
		// 			Type: to.Ptr("Microsoft.Logic/integrationAccounts/sessions"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Logic/integrationAccounts/testia123/sessions/IntegrationAccountSession7315"),
		// 			Properties: &armlogic.IntegrationAccountSessionProperties{
		// 				ChangedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-21T23:06:23.837Z"); return t}()),
		// 				Content: map[string]any{
		// 					"controlNumber": "1234",
		// 					"controlNumberChangedTime": "2017-02-21T22:30:11.9923759Z",
		// 				},
		// 				CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-21T23:06:23.836Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
