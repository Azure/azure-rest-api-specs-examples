Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmaintenance%2Farmmaintenance%2Fv0.2.0/sdk/resourcemanager/maintenance/armmaintenance/README.md) on how to add the SDK to your project and authenticate.

```go
package armmaintenance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance"
)

// x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2021-09-01-preview/examples/MaintenanceConfigurations_UpdateForResource.json
func ExampleConfigurationsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmaintenance.NewConfigurationsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armmaintenance.Configuration{
			Location: to.StringPtr("<location>"),
			Properties: &armmaintenance.ConfigurationProperties{
				MaintenanceScope: armmaintenance.MaintenanceScope("OSImage").ToPtr(),
				MaintenanceWindow: &armmaintenance.Window{
					Duration:           to.StringPtr("<duration>"),
					ExpirationDateTime: to.StringPtr("<expiration-date-time>"),
					RecurEvery:         to.StringPtr("<recur-every>"),
					StartDateTime:      to.StringPtr("<start-date-time>"),
					TimeZone:           to.StringPtr("<time-zone>"),
				},
				Namespace:  to.StringPtr("<namespace>"),
				Visibility: armmaintenance.Visibility("Custom").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ConfigurationsClientUpdateResult)
}
```
