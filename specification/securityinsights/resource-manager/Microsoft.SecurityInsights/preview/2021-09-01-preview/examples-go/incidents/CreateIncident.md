Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurityinsight%2Farmsecurityinsight%2Fv0.2.1/sdk/resourcemanager/securityinsight/armsecurityinsight/README.md) on how to add the SDK to your project and authenticate.

```go
package armsecurityinsight_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsight/armsecurityinsight"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/incidents/CreateIncident.json
func ExampleIncidentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsight.NewIncidentsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<incident-id>",
		armsecurityinsight.Incident{
			Etag: to.StringPtr("<etag>"),
			Properties: &armsecurityinsight.IncidentProperties{
				Description:           to.StringPtr("<description>"),
				Classification:        armsecurityinsight.IncidentClassification("FalsePositive").ToPtr(),
				ClassificationComment: to.StringPtr("<classification-comment>"),
				ClassificationReason:  armsecurityinsight.IncidentClassificationReason("IncorrectAlertLogic").ToPtr(),
				FirstActivityTimeUTC:  to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:00:30Z"); return t }()),
				LastActivityTimeUTC:   to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:05:30Z"); return t }()),
				Owner: &armsecurityinsight.IncidentOwnerInfo{
					ObjectID: to.StringPtr("<object-id>"),
				},
				Severity: armsecurityinsight.IncidentSeverity("High").ToPtr(),
				Status:   armsecurityinsight.IncidentStatus("Closed").ToPtr(),
				Title:    to.StringPtr("<title>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.IncidentsClientCreateOrUpdateResult)
}
```
