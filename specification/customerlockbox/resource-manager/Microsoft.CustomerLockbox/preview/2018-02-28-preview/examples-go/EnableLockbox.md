Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcustomerlockbox%2Farmcustomerlockbox%2Fv0.1.0/sdk/resourcemanager/customerlockbox/armcustomerlockbox/README.md) on how to add the SDK to your project and authenticate.

```go
package armcustomerlockbox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerlockbox/armcustomerlockbox"
)

// x-ms-original-file: specification/customerlockbox/resource-manager/Microsoft.CustomerLockbox/preview/2018-02-28-preview/examples/EnableLockbox.json
func ExamplePostClient_EnableLockbox() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcustomerlockbox.NewPostClient(cred, nil)
	_, err = client.EnableLockbox(ctx,
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
