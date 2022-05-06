Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fapplicationinsights%2Farmapplicationinsights%2Fv0.4.0/sdk/resourcemanager/applicationinsights/armapplicationinsights/README.md) on how to add the SDK to your project and authenticate.

```go
package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/ExportConfigurationsPost.json
func ExampleExportConfigurationsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewExportConfigurationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armapplicationinsights.ComponentExportRequest{
			DestinationAccountID:             to.Ptr("<destination-account-id>"),
			DestinationAddress:               to.Ptr("<destination-address>"),
			DestinationStorageLocationID:     to.Ptr("<destination-storage-location-id>"),
			DestinationStorageSubscriptionID: to.Ptr("<destination-storage-subscription-id>"),
			DestinationType:                  to.Ptr("<destination-type>"),
			IsEnabled:                        to.Ptr("<is-enabled>"),
			NotificationQueueEnabled:         to.Ptr("<notification-queue-enabled>"),
			NotificationQueueURI:             to.Ptr("<notification-queue-uri>"),
			RecordTypes:                      to.Ptr("<record-types>"),
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
