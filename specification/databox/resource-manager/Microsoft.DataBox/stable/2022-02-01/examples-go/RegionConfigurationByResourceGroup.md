```go
package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/databox/resource-manager/Microsoft.DataBox/stable/2022-02-01/examples/RegionConfigurationByResourceGroup.json
func ExampleServiceClient_RegionConfigurationByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatabox.NewServiceClient("fa68082f-8ff7-4a25-95c7-ce9da541242f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.RegionConfigurationByResourceGroup(ctx,
		"SdkRg4981",
		"westus",
		armdatabox.RegionConfigurationRequest{
			ScheduleAvailabilityRequest: &armdatabox.ScheduleAvailabilityRequest{
				SKUName:         to.Ptr(armdatabox.SKUNameDataBox),
				StorageLocation: to.Ptr("westus"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatabox%2Farmdatabox%2Fv1.0.0/sdk/resourcemanager/databox/armdatabox/README.md) on how to add the SDK to your project and authenticate.
