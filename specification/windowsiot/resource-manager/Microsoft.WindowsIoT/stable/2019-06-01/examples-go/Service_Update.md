Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fwindowsiot%2Farmwindowsiot%2Fv0.4.0/sdk/resourcemanager/windowsiot/armwindowsiot/README.md) on how to add the SDK to your project and authenticate.

```go
package armwindowsiot_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/windowsiot/armwindowsiot"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/windowsiot/resource-manager/Microsoft.WindowsIoT/stable/2019-06-01/examples/Service_Update.json
func ExampleServicesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armwindowsiot.NewServicesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<device-name>",
		armwindowsiot.DeviceService{
			Location: to.Ptr("<location>"),
			Properties: &armwindowsiot.DeviceServiceProperties{
				AdminDomainName:   to.Ptr("<admin-domain-name>"),
				BillingDomainName: to.Ptr("<billing-domain-name>"),
				Notes:             to.Ptr("<notes>"),
				Quantity:          to.Ptr[int64](1000000),
			},
		},
		&armwindowsiot.ServicesClientUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
