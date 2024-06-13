package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/entities/GetSecurityAlertEntityById.json
func ExampleEntitiesClient_Get_getASecurityAlertEntity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEntitiesClient().Get(ctx, "myRg", "myWorkspace", "4aa486e0-6f85-41af-99ea-7acdce7be6c8", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.EntitiesClientGetResponse{
	// 	                            EntityClassification: &armsecurityinsights.SecurityAlert{
	// 		Name: to.Ptr("e1d3d618-e11f-478b-98e3-bb381539a8e1"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/entities"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/entities/4aa486e0-6f85-41af-99ea-7acdce7be6c8"),
	// 		Kind: to.Ptr(armsecurityinsights.EntityKindSecurityAlert),
	// 		Properties: &armsecurityinsights.SecurityAlertProperties{
	// 			AdditionalData: map[string]any{
	// 				"Query": "Heartbeat \n| extend AccountCustomEntity = \"administrator\"",
	// 				"Query Period": "05:00:00",
	// 				"Search Query Results Overall Count": "203",
	// 				"Total Account Entities": "1",
	// 				"Trigger Operator": "GreaterThan",
	// 				"Trigger Threshold": "200",
	// 			},
	// 			FriendlyName: to.Ptr("Suspicious account detected"),
	// 			Description: to.Ptr(""),
	// 			AlertDisplayName: to.Ptr("Suspicious account detected"),
	// 			AlertLink: to.Ptr("https://portal.azure.com/#blade/Microsoft_Azure_Security/AlertBlade/alertId/2518119885989999999_4aa486e0-6f85-41af-99ea-7acdce7be6c8/subscriptionId/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/myRg/myWorkspace/referencedFrom/alertDeepLink/location/centralus"),
	// 			AlertType: to.Ptr("c8c99641-985d-4e4e-8e91-fb3466cd0e5b_46c7b6c0-ff43-44dd-8b4d-ceffff7aa7df"),
	// 			ConfidenceLevel: to.Ptr(armsecurityinsights.ConfidenceLevelUnknown),
	// 			EndTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-01T13:21:45.926Z"); return t}()),
	// 			Intent: to.Ptr(armsecurityinsights.KillChainIntentUnknown),
	// 			ProcessingEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-06T13:56:53.539Z"); return t}()),
	// 			ProductComponentName: to.Ptr("Scheduled Alerts"),
	// 			ProductName: to.Ptr("Azure Sentinel"),
	// 			ProviderAlertID: to.Ptr("c2bafff9-fb31-41d0-a177-ecbff7a02ffe"),
	// 			Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
	// 			StartTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-01T08:21:45.926Z"); return t}()),
	// 			Status: to.Ptr(armsecurityinsights.AlertStatusNew),
	// 			SystemAlertID: to.Ptr("4aa486e0-6f85-41af-99ea-7acdce7be6c8"),
	// 			Tactics: []*armsecurityinsights.AttackTactic{
	// 				to.Ptr(armsecurityinsights.AttackTacticPersistence),
	// 				to.Ptr(armsecurityinsights.AttackTacticLateralMovement)},
	// 				TimeGenerated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-01T13:56:53.539Z"); return t}()),
	// 				VendorName: to.Ptr("Microsoft"),
	// 			},
	// 		},
	// 		                        }
}
