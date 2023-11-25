package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/incidents/GetAllIncidentAlerts.json
func ExampleIncidentsClient_ListAlerts() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIncidentsClient().ListAlerts(ctx, "myRg", "myWorkspace", "afbd324f-6c48-459c-8710-8d1e1cd03812", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IncidentAlertList = armsecurityinsights.IncidentAlertList{
	// 	Value: []*armsecurityinsights.SecurityAlert{
	// 		{
	// 			Name: to.Ptr("baa8a239-6fde-4ab7-a093-d09f7b75c58c"),
	// 			Type: to.Ptr("Microsoft.SecurityInsights/Entities"),
	// 			ID: to.Ptr("/subscriptions/bd794837-4d29-4647-9105-6339bfdb4e6a/resourceGroups/myRG/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/Entities/baa8a239-6fde-4ab7-a093-d09f7b75c58c"),
	// 			Kind: to.Ptr(armsecurityinsights.EntityKindEnumSecurityAlert),
	// 			Properties: &armsecurityinsights.SecurityAlertProperties{
	// 				AdditionalData: map[string]any{
	// 					"AlertMessageEnqueueTime": "2020-07-20T18:21:57.304Z",
	// 				},
	// 				FriendlyName: to.Ptr("myAlert"),
	// 				AlertDisplayName: to.Ptr("myAlert"),
	// 				AlertType: to.Ptr("myAlert"),
	// 				ConfidenceLevel: to.Ptr(armsecurityinsights.ConfidenceLevelUnknown),
	// 				EndTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-20T18:21:53.615Z"); return t}()),
	// 				ProcessingEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-20T18:21:53.615Z"); return t}()),
	// 				ProductName: to.Ptr("Azure Security Center"),
	// 				ResourceIdentifiers: []any{
	// 					map[string]any{
	// 						"type": "LogAnalytics",
	// 						"resourceGroup": "myRG",
	// 						"subscriptionId": "bd794837-4d29-4647-9105-6339bfdb4e6a",
	// 						"workspaceId": "c8c99641-985d-4e4e-8e91-fb3466cd0e5b",
	// 				}},
	// 				Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
	// 				StartTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-20T18:21:53.615Z"); return t}()),
	// 				Status: to.Ptr(armsecurityinsights.AlertStatusNew),
	// 				SystemAlertID: to.Ptr("baa8a239-6fde-4ab7-a093-d09f7b75c58c"),
	// 				Tactics: []*armsecurityinsights.AttackTactic{
	// 				},
	// 				TimeGenerated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-20T18:21:53.615Z"); return t}()),
	// 				VendorName: to.Ptr("Microsoft"),
	// 			},
	// 	}},
	// }
}
