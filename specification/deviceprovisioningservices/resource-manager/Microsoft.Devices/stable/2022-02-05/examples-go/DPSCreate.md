Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdeviceprovisioningservices%2Farmdeviceprovisioningservices%2Fv1.0.0/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices/README.md) on how to add the SDK to your project and authenticate.

```go
package armdeviceprovisioningservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2022-02-05/examples/DPSCreate.json
func ExampleIotDpsResourceClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdeviceprovisioningservices.NewIotDpsResourceClient("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myFirstProvisioningService",
		armdeviceprovisioningservices.ProvisioningServiceDescription{
			Location: to.Ptr("East US"),
			Tags:     map[string]*string{},
			Properties: &armdeviceprovisioningservices.IotDpsPropertiesDescription{
				EnableDataResidency: to.Ptr(false),
			},
			SKU: &armdeviceprovisioningservices.IotDpsSKUInfo{
				Name:     to.Ptr(armdeviceprovisioningservices.IotDpsSKUS1),
				Capacity: to.Ptr[int64](1),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
