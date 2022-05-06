Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcommunication%2Farmcommunication%2Fv0.4.0/sdk/resourcemanager/communication/armcommunication/README.md) on how to add the SDK to your project and authenticate.

```go
package armcommunication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/communication/armcommunication"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/communication/resource-manager/Microsoft.Communication/stable/2020-08-20/examples/get.json
func ExampleServiceClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcommunication.NewServiceClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<communication-service-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
