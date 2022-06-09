```go
package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
)

// x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/MarkDevicesShipped.json
func ExampleJobsClient_MarkDevicesShipped() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatabox.NewJobsClient("<subscription-id>", cred, nil)
	_, err = client.MarkDevicesShipped(ctx,
		"<job-name>",
		"<resource-group-name>",
		armdatabox.MarkDevicesShippedRequest{
			DeliverToDcPackageDetails: &armdatabox.PackageCarrierInfo{
				CarrierName: to.StringPtr("<carrier-name>"),
				TrackingID:  to.StringPtr("<tracking-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatabox%2Farmdatabox%2Fv0.2.0/sdk/resourcemanager/databox/armdatabox/README.md) on how to add the SDK to your project and authenticate.
