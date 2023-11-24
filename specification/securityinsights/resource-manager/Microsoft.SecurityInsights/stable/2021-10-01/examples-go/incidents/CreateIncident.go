package armsecurityinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/incidents/CreateIncident.json
func ExampleIncidentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIncidentsClient().CreateOrUpdate(ctx, "myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", armsecurityinsights.Incident{
		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		Properties: &armsecurityinsights.IncidentProperties{
			Description:           to.Ptr("This is a demo incident"),
			Classification:        to.Ptr(armsecurityinsights.IncidentClassificationFalsePositive),
			ClassificationComment: to.Ptr("Not a malicious activity"),
			ClassificationReason:  to.Ptr(armsecurityinsights.IncidentClassificationReasonIncorrectAlertLogic),
			FirstActivityTimeUTC:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:00:30.000Z"); return t }()),
			LastActivityTimeUTC:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:05:30.000Z"); return t }()),
			Owner: &armsecurityinsights.IncidentOwnerInfo{
				ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
			},
			Severity: to.Ptr(armsecurityinsights.IncidentSeverityHigh),
			Status:   to.Ptr(armsecurityinsights.IncidentStatusClosed),
			Title:    to.Ptr("My incident"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Incident = armsecurityinsights.Incident{
	// 	Name: to.Ptr("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 	Type: to.Ptr("Microsoft.SecurityInsights/incidents"),
	// 	ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/incidents/73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 	Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0001\""),
	// 	Properties: &armsecurityinsights.IncidentProperties{
	// 		Description: to.Ptr("This is a demo incident"),
	// 		AdditionalData: &armsecurityinsights.IncidentAdditionalData{
	// 			AlertProductNames: []*string{
	// 			},
	// 			AlertsCount: to.Ptr[int32](0),
	// 			BookmarksCount: to.Ptr[int32](0),
	// 			CommentsCount: to.Ptr[int32](3),
	// 			Tactics: []*armsecurityinsights.AttackTactic{
	// 			},
	// 		},
	// 		Classification: to.Ptr(armsecurityinsights.IncidentClassificationFalsePositive),
	// 		ClassificationComment: to.Ptr("Not a malicious activity"),
	// 		ClassificationReason: to.Ptr(armsecurityinsights.IncidentClassificationReasonIncorrectAlertLogic),
	// 		CreatedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:15:30.000Z"); return t}()),
	// 		FirstActivityTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:00:30.000Z"); return t}()),
	// 		IncidentNumber: to.Ptr[int32](3177),
	// 		IncidentURL: to.Ptr("https://portal.azure.com/#asset/Microsoft_Azure_Security_Insights/Incident/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/incidents/73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 		Labels: []*armsecurityinsights.IncidentLabel{
	// 		},
	// 		LastActivityTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:05:30.000Z"); return t}()),
	// 		LastModifiedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:15:30.000Z"); return t}()),
	// 		Owner: &armsecurityinsights.IncidentOwnerInfo{
	// 			AssignedTo: to.Ptr("john doe"),
	// 			Email: to.Ptr("john.doe@contoso.com"),
	// 			ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
	// 			UserPrincipalName: to.Ptr("john@contoso.com"),
	// 		},
	// 		RelatedAnalyticRuleIDs: []*string{
	// 		},
	// 		Severity: to.Ptr(armsecurityinsights.IncidentSeverityHigh),
	// 		Status: to.Ptr(armsecurityinsights.IncidentStatusClosed),
	// 		Title: to.Ptr("My incident"),
	// 	},
	// }
}
