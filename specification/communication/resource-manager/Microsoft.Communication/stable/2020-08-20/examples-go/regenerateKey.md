Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcommunication%2Farmcommunication%2Fv0.2.0/sdk/resourcemanager/communication/armcommunication/README.md) on how to add the SDK to your project and authenticate.

```go
package armcommunication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/communication/armcommunication"
)

// x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/stable/2020-08-20/examples/regenerateKey.json
func ExampleServiceClient_RegenerateKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcommunication.NewServiceClient("<subscription-id>", cred, nil)
	res, err := client.RegenerateKey(ctx,
		"<resource-group-name>",
		"<communication-service-name>",
		armcommunication.RegenerateKeyParameters{
			KeyType: armcommunication.KeyTypePrimary.ToPtr(),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ServiceClientRegenerateKeyResult)
}
```
