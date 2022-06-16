package armsecurityinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/incidents/CreateIncident.json
func ExampleIncidentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewIncidentsClient("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"myRg",
		"myWorkspace",
		"73e01a99-5cd7-4139-a149-9f2736ff2ab5",
		armsecurityinsights.Incident{
			Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
			Properties: &armsecurityinsights.IncidentProperties{
				Description:           to.Ptr("This is a demo incident"),
				Classification:        to.Ptr(armsecurityinsights.IncidentClassificationFalsePositive),
				ClassificationComment: to.Ptr("Not a malicious activity"),
				ClassificationReason:  to.Ptr(armsecurityinsights.IncidentClassificationReasonIncorrectAlertLogic),
				FirstActivityTimeUTC:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:00:30Z"); return t }()),
				LastActivityTimeUTC:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:05:30Z"); return t }()),
				Owner: &armsecurityinsights.IncidentOwnerInfo{
					ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
				},
				Severity: to.Ptr(armsecurityinsights.IncidentSeverityHigh),
				Status:   to.Ptr(armsecurityinsights.IncidentStatusClosed),
				Title:    to.Ptr("My incident"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
