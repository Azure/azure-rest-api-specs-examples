Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatabox%2Farmdatabox%2Fv0.1.0/sdk/resourcemanager/databox/armdatabox/README.md) on how to add the SDK to your project and authenticate.

```go
package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
)

// x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/RegionConfiguration.json
func ExampleServiceClient_RegionConfiguration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatabox.NewServiceClient("<subscription-id>", cred, nil)
	_, err = client.RegionConfiguration(ctx,
		"<location>",
		armdatabox.RegionConfigurationRequest{
			ScheduleAvailabilityRequest: &armdatabox.DataBoxScheduleAvailabilityRequest{
				ScheduleAvailabilityRequest: armdatabox.ScheduleAvailabilityRequest{
					SKUName:         armdatabox.SKUNameDataBox.ToPtr(),
					StorageLocation: to.StringPtr("<storage-location>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
