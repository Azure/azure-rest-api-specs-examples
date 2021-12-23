Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fwindowsesu%2Farmwindowsesu%2Fv0.1.0/sdk/resourcemanager/windowsesu/armwindowsesu/README.md) on how to add the SDK to your project and authenticate.

```go
package armwindowsesu_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/windowsesu/armwindowsesu"
)

// x-ms-original-file: specification/windowsesu/resource-manager/Microsoft.WindowsESU/preview/2019-09-16-preview/examples/UpdateMultipleActivationKey.json
func ExampleMultipleActivationKeysClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armwindowsesu.NewMultipleActivationKeysClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<multiple-activation-key-name>",
		armwindowsesu.MultipleActivationKeyUpdate{
			Tags: map[string]*string{
				"tag1": to.StringPtr("value1"),
				"tag2": to.StringPtr("value2"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("MultipleActivationKey.ID: %s\n", *res.ID)
}
```
