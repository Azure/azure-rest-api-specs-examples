package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountSessions_CreateOrUpdate.json
func ExampleIntegrationAccountSessionsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIntegrationAccountSessionsClient().CreateOrUpdate(ctx, "testrg123", "testia123", "testsession123-ICN", armlogic.IntegrationAccountSession{
		Properties: &armlogic.IntegrationAccountSessionProperties{
			Content: map[string]any{
				"controlNumber":            "1234",
				"controlNumberChangedTime": "2017-02-21T22:30:11.9923759Z",
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IntegrationAccountSession = armlogic.IntegrationAccountSession{
	// 	Name: to.Ptr("testsession123-ICN"),
	// 	Type: to.Ptr("Microsoft.Logic/integrationAccounts/sessions"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Logic/integrationAccounts/testia123/sessions/testsession123-ICN"),
	// 	Properties: &armlogic.IntegrationAccountSessionProperties{
	// 		ChangedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-21T20:41:35.979Z"); return t}()),
	// 		Content: map[string]any{
	// 			"controlNumber": "1234",
	// 			"controlNumberChangedTime": "2017-02-21T22:30:11.9923759Z",
	// 		},
	// 		CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-21T20:41:35.978Z"); return t}()),
	// 	},
	// }
}
