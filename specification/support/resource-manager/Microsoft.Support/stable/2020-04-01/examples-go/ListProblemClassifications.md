Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsupport%2Farmsupport%2Fv0.1.0/sdk/resourcemanager/support/armsupport/README.md) on how to add the SDK to your project and authenticate.

```go
package armsupport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport"
)

// x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/ListProblemClassifications.json
func ExampleProblemClassificationsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsupport.NewProblemClassificationsClient(cred, nil)
	_, err = client.List(ctx,
		"<service-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
