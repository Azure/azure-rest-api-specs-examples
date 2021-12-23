Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fwindowsiot%2Farmwindowsiot%2Fv0.1.0/sdk/resourcemanager/windowsiot/armwindowsiot/README.md) on how to add the SDK to your project and authenticate.

```go
package armwindowsiot_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/windowsiot/armwindowsiot"
)

// x-ms-original-file: specification/windowsiot/resource-manager/Microsoft.WindowsIoT/stable/2019-06-01/examples/Service_Delete.json
func ExampleServicesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armwindowsiot.NewServicesClient("<subscription-id>", cred, nil)
	res, err := client.Delete(ctx,
		"<resource-group-name>",
		"<device-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DeviceService.ID: %s\n", *res.ID)
}
```
