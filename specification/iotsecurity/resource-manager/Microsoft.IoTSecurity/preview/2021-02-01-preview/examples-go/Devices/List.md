Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fiotsecurity%2Farmiotsecurity%2Fv0.4.0/sdk/resourcemanager/iotsecurity/armiotsecurity/README.md) on how to add the SDK to your project and authenticate.

```go
package armiotsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotsecurity/armiotsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/iotsecurity/resource-manager/Microsoft.IoTSecurity/preview/2021-02-01-preview/examples/Devices/List.json
func ExampleDevicesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armiotsecurity.NewDevicesClient("<subscription-id>",
		"<iot-defender-location>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.NewListPager("<device-group-name>",
		&armiotsecurity.DevicesClientListOptions{SkipToken: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
```
