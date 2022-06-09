```go
package armsecurityinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-04-01-preview/examples/incidents/CreateIncident.json
func ExampleIncidentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewIncidentsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<incident-id>",
		armsecurityinsights.Incident{
			Etag: to.Ptr("<etag>"),
			Properties: &armsecurityinsights.IncidentProperties{
				Description:           to.Ptr("<description>"),
				Classification:        to.Ptr(armsecurityinsights.IncidentClassificationFalsePositive),
				ClassificationComment: to.Ptr("<classification-comment>"),
				ClassificationReason:  to.Ptr(armsecurityinsights.IncidentClassificationReasonIncorrectAlertLogic),
				FirstActivityTimeUTC:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:00:30Z"); return t }()),
				LastActivityTimeUTC:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:05:30Z"); return t }()),
				Owner: &armsecurityinsights.IncidentOwnerInfo{
					ObjectID: to.Ptr("<object-id>"),
				},
				Severity: to.Ptr(armsecurityinsights.IncidentSeverityHigh),
				Status:   to.Ptr(armsecurityinsights.IncidentStatusClosed),
				Title:    to.Ptr("<title>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurityinsights%2Farmsecurityinsights%2Fv0.3.0/sdk/resourcemanager/securityinsights/armsecurityinsights/README.md) on how to add the SDK to your project and authenticate.
