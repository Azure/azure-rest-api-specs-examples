Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmaintenance%2Farmmaintenance%2Fv0.4.0/sdk/resourcemanager/maintenance/armmaintenance/README.md) on how to add the SDK to your project and authenticate.

```go
package armmaintenance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2021-09-01-preview/examples/MaintenanceConfigurations_UpdateForResource.json
func ExampleConfigurationsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmaintenance.NewConfigurationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armmaintenance.Configuration{
			Location: to.Ptr("<location>"),
			Properties: &armmaintenance.ConfigurationProperties{
				MaintenanceScope: to.Ptr(armmaintenance.MaintenanceScopeOSImage),
				MaintenanceWindow: &armmaintenance.Window{
					Duration:           to.Ptr("<duration>"),
					ExpirationDateTime: to.Ptr("<expiration-date-time>"),
					RecurEvery:         to.Ptr("<recur-every>"),
					StartDateTime:      to.Ptr("<start-date-time>"),
					TimeZone:           to.Ptr("<time-zone>"),
				},
				Namespace:  to.Ptr("<namespace>"),
				Visibility: to.Ptr(armmaintenance.VisibilityCustom),
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
