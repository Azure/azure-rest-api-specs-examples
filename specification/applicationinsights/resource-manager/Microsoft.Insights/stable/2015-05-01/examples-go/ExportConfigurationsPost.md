Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fapplicationinsights%2Farmapplicationinsights%2Fv0.2.0/sdk/resourcemanager/applicationinsights/armapplicationinsights/README.md) on how to add the SDK to your project and authenticate.

```go
package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
)

// x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/ExportConfigurationsPost.json
func ExampleExportConfigurationsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapplicationinsights.NewExportConfigurationsClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armapplicationinsights.ComponentExportRequest{
			DestinationAccountID:             to.StringPtr("<destination-account-id>"),
			DestinationAddress:               to.StringPtr("<destination-address>"),
			DestinationStorageLocationID:     to.StringPtr("<destination-storage-location-id>"),
			DestinationStorageSubscriptionID: to.StringPtr("<destination-storage-subscription-id>"),
			DestinationType:                  to.StringPtr("<destination-type>"),
			IsEnabled:                        to.StringPtr("<is-enabled>"),
			NotificationQueueEnabled:         to.StringPtr("<notification-queue-enabled>"),
			NotificationQueueURI:             to.StringPtr("<notification-queue-uri>"),
			RecordTypes:                      to.StringPtr("<record-types>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ExportConfigurationsClientCreateResult)
}
```
