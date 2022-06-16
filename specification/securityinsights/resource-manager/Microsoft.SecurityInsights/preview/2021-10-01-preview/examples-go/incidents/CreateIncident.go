package armsecurityinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-10-01-preview/examples/incidents/CreateIncident.json
func ExampleIncidentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsights.NewIncidentsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<incident-id>",
		armsecurityinsights.Incident{
			Etag: to.StringPtr("<etag>"),
			Properties: &armsecurityinsights.IncidentProperties{
				Description:           to.StringPtr("<description>"),
				Classification:        armsecurityinsights.IncidentClassification("FalsePositive").ToPtr(),
				ClassificationComment: to.StringPtr("<classification-comment>"),
				ClassificationReason:  armsecurityinsights.IncidentClassificationReason("IncorrectAlertLogic").ToPtr(),
				FirstActivityTimeUTC:  to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:00:30Z"); return t }()),
				LastActivityTimeUTC:   to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:05:30Z"); return t }()),
				Owner: &armsecurityinsights.IncidentOwnerInfo{
					ObjectID: to.StringPtr("<object-id>"),
				},
				Severity: armsecurityinsights.IncidentSeverity("High").ToPtr(),
				Status:   armsecurityinsights.IncidentStatus("Closed").ToPtr(),
				Title:    to.StringPtr("<title>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.IncidentsClientCreateOrUpdateResult)
}
